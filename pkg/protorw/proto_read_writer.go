// package protorw solves the size delimination problem for protobuf messages.
// As the protobuf documentation (https://protobuf.dev/programming-guides/techniques/#streaming) states,
// "it is up to you to keep track of where one message ends and the next begins".
// The easiest solution for doing so, is to prefix every message with it's length/size.
// ProtoReadWriter implements this behavior
package protorw

import (
	"encoding/binary"
	"io"

	"google.golang.org/protobuf/proto"
)

// Message is taken from the protobuf libary
type Message proto.Message

// ProtoReadWriter describes the interface to be fulfilled
// for a reader-writer with the purpose to be used
// for protobuf messages.
//
// protoReadWriter implements a size delimiting mechanism on
// top of an reader-writer. The size is limtited by
// the limit of int64 which means that it needs 8 bytes to
// be encoded for transmission. It uses big endian to encode
// the size into bytes
//
// For reading from that connection that means it will assume
// a 8-byte size prefixed message.
//
// For writing this means that every message will get size prefixed
// by 8 bytes before writing the message itself.
type ProtoReadWriter interface {
	io.ReadWriter
	// WriteMessage writes a message into the underlying
	WriteMessage(Message) error
	ReadMessage(Message) error
}

// New returns an implementation of the ProtoReadWriter interface
func New(rw io.ReadWriter) ProtoReadWriter {
	return &protoReadWriter{underlying: rw}
}

type protoReadWriter struct {
	underlying io.ReadWriter
}

func (c *protoReadWriter) WriteMessage(m Message) error {
	size, serializedMessage, err := marshalMessage(m)
	if err != nil {
		return err
	}
	_, err = write(c.underlying, size, serializedMessage)
	return err
}

func (c *protoReadWriter) ReadMessage(m Message) error {
	read, _, err := read(c.underlying)
	if err != nil {
		return err
	}
	return proto.Unmarshal(read, m)
}

func (c *protoReadWriter) Read(b []byte) (n int, err error) {
	read, size, err := read(c.underlying)
	if err != nil {
		return 0, nil
	}
	copy(b, read)
	return int(size), err
}

func (c *protoReadWriter) Write(b []byte) (n int, err error) {
	return write(c.underlying, uint64(len(b)), b)
}

// WriteMessage marshals a given protobuf message and writes it into
// the given writer. The payload will be size prefixed by eight bytes
// to enable size delimited reading of the message.
func WriteMessage(w io.Writer, m Message) error {
	size, serializedMessage, err := marshalMessage(m)
	if err != nil {
		return err
	}
	_, err = write(w, size, serializedMessage)
	return err
}

// ReadMessage reads a protobuf message from a given connection/reader.
// It assumes that the first eight byte are used for specifying
// the message's length
func ReadMessage(r io.Reader, m Message) error {
	read, _, err := read(r)
	if err != nil {
		return err
	}
	return proto.Unmarshal(read, m)
}

func write(w io.Writer, size uint64, payload []byte) (n int, err error) {
	sizeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(sizeBytes, size)
	n, err = w.Write(append(sizeBytes, payload...))
	if err != nil {
		return 0, err
	}
	return n, nil
}

func read(r io.Reader) (b []byte, size uint64, err error) {
	sizeBytes := make([]byte, 8)
	_, err = r.Read(sizeBytes)
	if err != nil {
		return nil, 0, err
	}
	size = binary.BigEndian.Uint64(sizeBytes)
	messageBytes := make([]byte, size)
	n, err := r.Read(messageBytes)
	return messageBytes, uint64(n), err

}

func marshalMessage(m Message) (size uint64, b []byte, err error) {
	size = uint64(proto.Size(m))
	b, err = proto.Marshal(m)
	return
}
