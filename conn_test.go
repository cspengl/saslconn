package saslconn

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	"github.com/cspengl/saslconn/internal/saslproto"
	"github.com/cspengl/saslconn/pkg/protorw"

	"github.com/emersion/go-sasl"
)

func testConfig() *Config {
	return &Config{
		Mechanisms: []*Mechanism{
			{
				Name:   sasl.Plain,
				Client: sasl.NewPlainClient("foo", "foo", "bar"),
				Server: sasl.NewPlainServer(
					func(identity, username, password string) error {
						if username != "foo" {
							return errors.New("wrong username")
						}
						if password != "bar" {
							return errors.New("wrong password")
						}
						return nil
					},
				),
			},
		},
	}
}

func TestSimpleConn(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	client = Client(client, testConfig())
	server = Server(server, testConfig())

	go func() {
		client.Write([]byte("Hello World"))
	}()

	buf := make([]byte, 100)
	n, err := server.Read(buf)
	if err != nil {
		t.Error(err)

	} else if string(buf[:n]) != "Hello World" {
		t.Fatal(string(buf))
	}
	state := client.(*Conn).ConnectionState()
	if !state.HandshakeComplete {
		t.Fatal("handshake not complete")
	}
}

func TestWithDialAndListen(t *testing.T) {

	listen, err := Listen("tcp", ":9001", testConfig())
	if err != nil {
		t.Error(err)
	}
	defer listen.Close()

	errChan := make(chan error, 1)
	go func() {
		conn, err := listen.Accept()
		defer func() {
			if err := conn.Close(); err != nil {
				panic(err)
			}
		}()
		if err != nil {
			t.Error(err)
		}
		buf := make([]byte, 100)
		n, err := conn.Read(buf)
		if err != nil {
			errChan <- err
			return

		} else if string(buf[:n]) != "Hello World" {
			errChan <- errors.New("unexpected string")
			return
		}
		errChan <- nil
	}()

	conn, err := Dial("tcp", ":9001", testConfig())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte("Hello World"))
	if err != nil {
		t.Error(err)
	}
	err = <-errChan
	if err != nil {
		t.Error(err)
	}
}

func TestReject(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	rejectConfig := testConfig()
	rejectConfig.Mechanisms[0].Server = sasl.NewPlainServer(
		func(identity, username, password string) error { return errors.New("failed") },
	)

	client = Client(client, testConfig())
	server = Server(server, rejectConfig)

	go func() {
		client.Write([]byte("Hello World"))
	}()

	buf := make([]byte, 100)
	_, err := server.Read(buf)
	if err == nil {
		t.Fatal("did not get expected error")

	}

}

type saslMessageSchedule struct {
	conn            net.Conn
	index           int
	schedule        []*saslproto.Message
	expectFinalRead bool
}

func (s *saslMessageSchedule) Run(role mechanismRoleType) error {
	if role == roleServer {
		if err := s.Next(); err != nil {
			return err
		}
	}

	for s.index < len(s.schedule) {
		_, err := s.conn.Read(make([]byte, 1024))
		if err != nil {
			return err
		}
		if err := s.Next(); err != nil {
			return err
		}
	}
	if s.expectFinalRead {
		_, err := s.conn.Read(make([]byte, 1024))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *saslMessageSchedule) Next() (err error) {
	err = protorw.WriteMessage(s.conn, s.schedule[s.index])
	s.index++
	return
}

func Test_handshake_client_errors(t *testing.T) {

	tests := []struct {
		name      string
		schedule  []*saslproto.Message
		finalRead bool
	}{
		{
			name: "server starts with challenge",
			schedule: []*saslproto.Message{
				{
					MessageType: saslproto.MessageType_MessageTypeChallengeResponse,
					Payload:     &saslproto.Message_ChallengeResponse{},
				},
			},
		},
		{
			name: "server does not advertise supported mechanism",
			schedule: []*saslproto.Message{
				{
					MessageType: saslproto.MessageType_MessageTypeServerMechanismAdvertisement,
					Payload: &saslproto.Message_ServerMechanismAdvertisement{
						ServerMechanismAdvertisement: &saslproto.ServerMechanismAdvertisement{
							Mechanisms: []string{},
						},
					},
				},
			},
		},
		{
			name: "server sends challenge",
			schedule: []*saslproto.Message{
				{
					MessageType: saslproto.MessageType_MessageTypeServerMechanismAdvertisement,
					Payload: &saslproto.Message_ServerMechanismAdvertisement{
						ServerMechanismAdvertisement: &saslproto.ServerMechanismAdvertisement{
							Mechanisms: testConfig().mechanismList(),
						},
					},
				},
				{
					MessageType: saslproto.MessageType_MessageTypeChallengeResponse,
					Payload: &saslproto.Message_ChallengeResponse{
						ChallengeResponse: &saslproto.ChallengeResponse{
							Payload: []byte{},
						},
					},
				},
			},
			finalRead: true,
		},
		{
			name: "server aborts",
			schedule: []*saslproto.Message{
				{
					MessageType: saslproto.MessageType_MessageTypeServerMechanismAdvertisement,
					Payload: &saslproto.Message_ServerMechanismAdvertisement{
						ServerMechanismAdvertisement: &saslproto.ServerMechanismAdvertisement{
							Mechanisms: testConfig().mechanismList(),
						},
					},
				},
				{
					MessageType: saslproto.MessageType_MessageTypeHandshakeAbortion,
					Payload: &saslproto.Message_HandshakeAbortion{
						HandshakeAbortion: &saslproto.HandshakeAbortion{
							Message: "server side error",
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := net.Pipe()
			defer client.Close()
			defer server.Close()

			client = Client(client, testConfig())

			schedule := &saslMessageSchedule{
				conn:            server,
				index:           0,
				schedule:        test.schedule,
				expectFinalRead: test.finalRead,
			}

			errChan := make(chan error, 1)
			go func() {
				errChan <- schedule.Run(roleServer)
			}()

			err := client.(*Conn).Handshake()
			if err == nil {
				t.Fatal("did not get expected error")
			}
			err = <-errChan
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestConnectionFunct(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	conn := Client(client, testConfig())

	if conn.LocalAddr() != client.LocalAddr() {
		t.Fatal("local address neq")
	}
	if conn.RemoteAddr() != conn.RemoteAddr() {
		t.Fatal("remote addr neq")
	}
	if err := conn.SetDeadline(time.Now().Add(1 * time.Second)); err != nil {
		t.Error(err)
	}
	if err := conn.SetReadDeadline(time.Now().Add(1 * time.Second)); err != nil {
		t.Error(err)
	}
	if err := conn.SetWriteDeadline(time.Now().Add(1 * time.Second)); err != nil {
		t.Error(err)
	}
}

func TestHandshakeCancelation(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	conn := Client(client, testConfig())

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	conn.HandshakeContext(ctx)
}
