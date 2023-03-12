package protorw

import (
	"bytes"
	"io"
	"testing"

	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/protobuf/proto"
)

func testMessage() Message {
	return &pb.HelloRequest{Name: "Bob"}
}

func assertMessage(t *testing.T, r *pb.HelloRequest) {
	if r.Name != "Bob" {
		t.Fatal("unexpected name")
	}
}

func TestWriteMessage(t *testing.T) {
	if err := WriteMessage(io.Discard, testMessage()); err != nil {
		t.Error(err)
	}
}

func TestReadMessage(t *testing.T) {

	size := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5}
	payload, _ := proto.Marshal(testMessage())

	reader := bytes.NewBuffer(append(size, payload...))
	var receivedMessage pb.HelloRequest
	if err := ReadMessage(reader, &receivedMessage); err != nil {
		t.Error(err)
	}
	assertMessage(t, &receivedMessage)
}

func TestProtoReadWriter(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	rw := New(buf)

	if err := rw.WriteMessage(testMessage()); err != nil {
		t.Error(err)
	}

	var receivedMessage pb.HelloRequest
	if err := rw.ReadMessage(&receivedMessage); err != nil {
		t.Error(err)
	}
	assertMessage(t, &receivedMessage)

	payload, _ := proto.Marshal(testMessage())
	if _, err := rw.Write(payload); err != nil {
		t.Error(err)
	}
	//5 is the expected size of message
	received := make([]byte, 5)
	if _, err := rw.Read(received); err != nil {
		t.Error(err)
	}
	if err := proto.Unmarshal(received, &receivedMessage); err != nil {
		t.Error(err)
	}
	assertMessage(t, &receivedMessage)
}
