package saslconn

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cspengl/saslconn/internal/saslproto"
	"github.com/cspengl/saslconn/pkg/protorw"

	"github.com/emersion/go-sasl"
	"github.com/pkg/errors"
)

var (
	// ErrHandshakeAborted gets returned from handshake if the
	// client or server aborted the ongoing handshake
	ErrHandshakeAborted = errors.New("handshake got aborted")
)

const (
	successMessage = "SUCCESS"
)

// Conn describes a SASL authenticated connection. It implements the net.Conn interface
type Conn struct {
	net.Conn
	conn                net.Conn
	isClient            bool
	config              *Config
	handshakeFn         func(context.Context) error
	handshakeComplete   atomic.Bool
	handshakeMutex      sync.Mutex
	handshakeErr        error
	negotiatedMechanism string
}

// ConnectionState describes the state of a SASL connection
type ConnectionState struct {
	// HandshakeComplete gives the status of the handshake
	HandshakeComplete bool
	// NegotiatedMechanism contains the mechanism negotiated
	// during the handshake. If it has not yet been selected
	// it is ""
	NegotiatedMechanism string
}

// Client returns a new SASL connection using the given net.Conn
// as the underlying transport. It will act as the client.
func Client(conn net.Conn, config *Config) *Conn {
	c := &Conn{
		conn:     conn,
		isClient: true,
		config:   config,
	}
	c.handshakeFn = c.clientHandshake
	return c
}

// Server returns a new SASL connection using the given net.Conn
// as the underlying transport. It will act as the server.
func Server(conn net.Conn, config *Config) *Conn {
	c := &Conn{
		conn:     conn,
		isClient: false,
		config:   config,
	}
	c.handshakeFn = c.serverHandshake
	return c
}

// Dial connects to the given network address using net.Dial and
// initiates a handshake returning the resulting SASL connection
func Dial(network, addr string, config *Config) (*Conn, error) {
	return DialWithDialer(new(net.Dialer), network, addr, config)
}

// DialWithDialer connects to a given network address using net.Dial and
// initiates a handshake returning the resulting SASL connection
func DialWithDialer(dialer *net.Dialer, network, addr string, config *Config) (*Conn, error) {
	return dial(context.Background(), dialer, network, addr, config)
}

func dial(ctx context.Context, netDialer *net.Dialer, network, addr string, config *Config) (*Conn, error) {
	if netDialer.Timeout != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, netDialer.Timeout)
		defer cancel()
	}

	if !netDialer.Deadline.IsZero() {
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadline(ctx, netDialer.Deadline)
		defer cancel()
	}

	rawConn, err := netDialer.DialContext(ctx, network, addr)
	if err != nil {
		return nil, err
	}

	conn := Client(rawConn, config)
	if err := conn.HandshakeContext(ctx); err != nil {
		//rawConn.Close()
		return nil, err
	}
	return conn, nil
}

type listener struct {
	net.Listener
	config *Config
}

func (l *listener) Accept() (net.Conn, error) {
	c, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}
	return Server(c, l.config), nil
}

// NewListener creates a listener accepting connections from an inner
// listener. It's accept function will return a SASL conn. It does not
// initiate a handshake.
func NewListener(inner net.Listener, config *Config) net.Listener {
	return &listener{
		Listener: inner,
		config:   config,
	}
}

// Listen creates a SASL listener by calling NewListener after
// using net.Listen to create the inner listener
func Listen(network, laddr string, config *Config) (net.Listener, error) {
	l, err := net.Listen(network, laddr)
	if err != nil {
		return nil, err
	}
	return NewListener(l, config), nil
}

// ListenTLS creates a tls listener before creating a SASL listener.
// This enables using a already secured TLS connection before transmitting
// credentials in plain text
func ListenTLS(network, laddr string, tlsConfig *tls.Config, config *Config) (net.Listener, error) {
	l, err := tls.Listen(network, laddr, tlsConfig)
	if err != nil {
		return nil, err
	}
	return NewListener(l, config), nil
}

// Read reads data from the connection.
//
// As Read calls Handshake, in order to prevent indefinite blocking a deadline
// must be set for both Read and Write before Read is called when the handshake
// has not yet completed. See SetDeadline, SetReadDeadline, and SetWriteDeadline.
func (c *Conn) Read(b []byte) (n int, err error) {
	if err := c.Handshake(); err != nil {
		return 0, io.EOF
	}
	if len(b) == 0 {
		return 0, nil
	}
	return c.conn.Read(b)
}

// Write writes data to the connection.
//
// As Write calls Handshake, in order to prevent indefinite blocking a deadline
// must be set for both Read and Write before Write is called when the handshake
// has not yet completed. See SetDeadline, SetReadDeadline, and
// SetWriteDeadline.
func (c *Conn) Write(b []byte) (n int, err error) {
	if err := c.Handshake(); err != nil {
		return 0, err
	}
	return c.conn.Write(b)
}

// NetConn returns the underlying net.Conn
func (c *Conn) NetConn() net.Conn {
	return c.conn
}

// Close closes the inner connection
func (c *Conn) Close() error {
	return c.conn.Close()
}

// LocalAddr returns the local address of the inner connection
func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

// RemoteAddr returns the remote address of the inner connection
func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

// SetDeadline sets the read and write deadline for the inner connection
func (c *Conn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

// SetReadDeadline sets the read deadline for the inner connection
func (c *Conn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

// SetWriteDeadline sets the write deadline for the inner connection
func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

// ConnectionState returns the connection state of the connection
func (c *Conn) ConnectionState() ConnectionState {
	c.handshakeMutex.Lock()
	defer c.handshakeMutex.Unlock()
	return ConnectionState{
		HandshakeComplete:   c.handshakeComplete.Load(),
		NegotiatedMechanism: c.negotiatedMechanism,
	}
}

// Handshake calls HandshakeContext internally by using context.Background()
//
// For controlling canceling or timeouts use HandshakeContext instead
func (c *Conn) Handshake() error {
	return c.HandshakeContext(context.Background())
}

// HandshakeContext initiates the handshake based on the connection
// role as client or server if it has not been run.
//
// Since it is called as part of Read and Write it is mostly not
// neccessary to call it manually.
func (c *Conn) HandshakeContext(ctx context.Context) error {
	return c.handshakeContext(ctx)
}

func (c *Conn) handshakeContext(ctx context.Context) (err error) {
	if c == nil {
		panic("ahhhh")
	}
	if c.handshakeComplete.Load() {
		return nil
	}

	if c.config == nil {
		c.config = defaultConfig()
	}

	c.handshakeErr = c.config.validate(c.role())
	if c.handshakeErr != nil {
		return c.handshakeErr
	}

	handshakeCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	if ctx.Done() != nil {
		done := make(chan struct{})
		interruptRes := make(chan error, 1)
		defer func() {
			close(done)
			if ctxErr := <-interruptRes; ctxErr != nil {
				err = ctxErr
			}
		}()
		go func() {
			select {
			case <-handshakeCtx.Done():
				_ = c.conn.Close()
				interruptRes <- handshakeCtx.Err()
			case <-done:
				interruptRes <- nil
			}
		}()
	}

	c.handshakeMutex.Lock()
	defer c.handshakeMutex.Unlock()

	if c.handshakeErr != nil {
		return err
	}
	if c.handshakeComplete.Load() {
		return nil
	}

	c.handshakeErr = c.handshakeFn(handshakeCtx)

	if c.handshakeErr == nil {
		c.handshakeComplete.Store(true)
	}

	return c.handshakeErr
}

func (c *Conn) clientHandshake(ctx context.Context) error {

	// 1. read advertised mechanisms
	advertisedMechanisms, err := c.readAdvertisedMechanisms()
	if err != nil {
		return err
	}

	// 2. select first supported mechanism
	for _, advertisedMechanism := range advertisedMechanisms {
		if c.config.mechanismIsSupported(advertisedMechanism) {
			c.negotiatedMechanism = advertisedMechanism
		}
	}
	if c.negotiatedMechanism == "" {
		return fmt.Errorf("server did not advertise supported mechanism type")
	}

	mechanism := c.config.getMechanism(c.negotiatedMechanism)

	if err := mechanism.validate(c.role()); err != nil {
		c.writeHandshakeAborted(err.Error())
		return err
	}

	// 3. Initiate challenge response phase
	_, initialResponse, err := mechanism.Client.Start()
	if err != nil {
		c.writeHandshakeAborted(err.Error())
		return err
	}

	if err := c.writeClientInitiation(c.negotiatedMechanism, initialResponse); err != nil {
		return err
	}

	// 4. challenge response phase
	for {
		var serverMessage saslproto.Message
		if err := protorw.ReadMessage(c.conn, &serverMessage); err != nil {
			return err
		}

		switch serverMessage.MessageType {
		case saslproto.MessageType_MessageTypeServerDone:
			if result := serverMessage.GetServerDone(); result.Result == saslproto.ServerDoneResult_ResultSuccess {
				return nil
			} else {
				return errors.New(result.Message)
			}
		case saslproto.MessageType_MessageTypeChallengeResponse:
			response, err := mechanism.Client.Next(serverMessage.GetChallengeResponse().Payload)
			if err != nil {
				c.writeHandshakeAborted(err.Error())
				return err
			}
			if err := c.writeChallengeResponse(response); err != nil {
				return err
			}
		case saslproto.MessageType_MessageTypeHandshakeAbortion:
			return errors.Wrap(
				ErrHandshakeAborted,
				serverMessage.GetHandshakeAbortion().Message,
			)
		default:
			return sasl.ErrUnexpectedServerChallenge
		}
	}
}

func (c *Conn) serverHandshake(ctx context.Context) error {

	//1. advertise mechanism to client
	if err := c.advertiseMechanisms(); err != nil {
		return err
	}

	//2. read client initiation
	clientInitiation, err := c.readClientInitiation()
	if err != nil {
		return err
	}

	//3. check if mechanism is supported
	if !c.config.mechanismIsSupported(clientInitiation.Mechanism) {
		return fmt.Errorf("%s, as initiated by client is not supported", clientInitiation.Mechanism)
	} else {
		c.negotiatedMechanism = clientInitiation.Mechanism
	}

	mechanism := *c.config.getMechanism(c.negotiatedMechanism)

	if err := mechanism.validate(c.role()); err != nil {
		c.writeHandshakeAborted(err.Error())
		return err
	}

	//4 if mechanism supported enter challenge response loop
	response := clientInitiation.InitialResponse
	for {
		challenge, done, err := mechanism.Server.Next(response)
		if err != nil {
			if done {
				c.writeServerDoneMessage(saslproto.ServerDoneResult_ResultReject, err.Error())
			} else {
				c.writeHandshakeAborted(err.Error())
			}
			return err
		}

		if done {
			c.writeServerDoneMessage(saslproto.ServerDoneResult_ResultSuccess, successMessage)
			return nil
		}

		if err := c.writeChallengeResponse(challenge); err != nil {
			return err
		}

		response, err = c.readChallengeResponse()
		if err != nil {
			return err
		}
	}
}

func (c *Conn) role() mechanismRoleType {
	if c.isClient {
		return roleClient
	}
	return roleServer
}

func defaultConfig() *Config {
	return &Config{
		Mechanisms: []*Mechanism{},
	}
}

func (c *Conn) advertiseMechanisms() error {
	return protorw.WriteMessage(
		c.conn,
		&saslproto.Message{
			MessageType: saslproto.MessageType_MessageTypeServerMechanismAdvertisement,
			Payload: &saslproto.Message_ServerMechanismAdvertisement{
				ServerMechanismAdvertisement: &saslproto.ServerMechanismAdvertisement{
					Mechanisms: c.config.mechanismList(),
				},
			},
		},
	)
}

func (c *Conn) readMessage(expectedType saslproto.MessageType) (*saslproto.Message, error) {
	var saslMessage saslproto.Message
	err := protorw.ReadMessage(c.conn, &saslMessage)
	if err != nil {
		return nil, err
	}
	if saslMessage.MessageType == saslproto.MessageType_MessageTypeHandshakeAbortion {
		return nil, errors.Wrap(
			ErrHandshakeAborted,
			saslMessage.GetHandshakeAbortion().Message,
		)
	}
	if saslMessage.MessageType != expectedType {
		return nil, fmt.Errorf(
			"unexpected sasl message type: got %s but expected %s",
			saslMessage.MessageType.String(), expectedType.String(),
		)
	}
	return &saslMessage, nil
}

func (c *Conn) readAdvertisedMechanisms() (mechanisms []string, err error) {
	advertisement, err := c.readMessage(saslproto.MessageType_MessageTypeServerMechanismAdvertisement)
	if err != nil {
		return []string{}, err
	}
	return advertisement.GetServerMechanismAdvertisement().Mechanisms, nil
}

func (c *Conn) writeClientInitiation(selectedMechanism string, initialResponse []byte) error {
	return protorw.WriteMessage(
		c.conn,
		&saslproto.Message{
			MessageType: saslproto.MessageType_MessageTypeClientInitiation,
			Payload: &saslproto.Message_ClientInitiation{
				ClientInitiation: &saslproto.ClientInitiation{
					Mechanism:           selectedMechanism,
					InitialReponseIsNil: initialResponse == nil,
					InitialResponse:     initialResponse,
				},
			},
		},
	)
}

func (c *Conn) readClientInitiation() (*saslproto.ClientInitiation, error) {
	clientInitiation, err := c.readMessage(saslproto.MessageType_MessageTypeClientInitiation)
	if err != nil {
		return nil, err
	}
	return clientInitiation.GetClientInitiation(), nil
}

func (c *Conn) writeServerDoneMessage(result saslproto.ServerDoneResult, message string) error {
	return protorw.WriteMessage(
		c.conn,
		&saslproto.Message{
			MessageType: saslproto.MessageType_MessageTypeServerDone,
			Payload: &saslproto.Message_ServerDone{
				ServerDone: &saslproto.ServerDone{
					Result:  result,
					Message: message,
				},
			},
		},
	)
}

func (c *Conn) writeChallengeResponse(payload []byte) error {
	return protorw.WriteMessage(
		c.conn,
		&saslproto.Message{
			MessageType: saslproto.MessageType_MessageTypeChallengeResponse,
			Payload: &saslproto.Message_ChallengeResponse{
				ChallengeResponse: &saslproto.ChallengeResponse{
					Payload: payload,
				},
			},
		},
	)
}

func (c *Conn) readChallengeResponse() (payload []byte, err error) {
	challengeResponse, err := c.readMessage(saslproto.MessageType_MessageTypeChallengeResponse)
	if err != nil {
		return []byte{}, nil
	}
	return challengeResponse.GetChallengeResponse().Payload, nil
}

func (c *Conn) writeHandshakeAborted(msg string) error {
	return protorw.WriteMessage(
		c.conn,
		&saslproto.Message{
			MessageType: saslproto.MessageType_MessageTypeHandshakeAbortion,
			Payload: &saslproto.Message_HandshakeAbortion{
				HandshakeAbortion: &saslproto.HandshakeAbortion{
					Message: msg,
				},
			},
		},
	)
}
