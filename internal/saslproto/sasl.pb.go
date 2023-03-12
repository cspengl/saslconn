//*
// SASL protocol messages.
//
// This file contains the SASL protocol definitions according to section 4 of
// RFC-4422.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: sasl.proto

package saslproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// MessageType is used for specifying the type of
// message transmitted.
type MessageType int32

const (
	// Used for error cases or as fallback
	MessageType_MessageTypeUnknown MessageType = 0
	// Used for ServerMechanismAdvertisement messages
	MessageType_MessageTypeServerMechanismAdvertisement MessageType = 1
	// Used for ClientInitiation messages
	MessageType_MessageTypeClientInitiation MessageType = 2
	// Used for ChallengeResponse messages
	MessageType_MessageTypeChallengeResponse MessageType = 3
	// Used for HandshakeAbortion messages
	MessageType_MessageTypeHandshakeAbortion MessageType = 4
	// Used for ServerDone messages
	MessageType_MessageTypeServerDone MessageType = 5
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "MessageTypeUnknown",
		1: "MessageTypeServerMechanismAdvertisement",
		2: "MessageTypeClientInitiation",
		3: "MessageTypeChallengeResponse",
		4: "MessageTypeHandshakeAbortion",
		5: "MessageTypeServerDone",
	}
	MessageType_value = map[string]int32{
		"MessageTypeUnknown":                      0,
		"MessageTypeServerMechanismAdvertisement": 1,
		"MessageTypeClientInitiation":             2,
		"MessageTypeChallengeResponse":            3,
		"MessageTypeHandshakeAbortion":            4,
		"MessageTypeServerDone":                   5,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_sasl_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_sasl_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_sasl_proto_rawDescGZIP(), []int{0}
}

type ServerDoneResult int32

const (
	ServerDoneResult_ResultUnknown ServerDoneResult = 0 // Used for error cases or as fallback
	ServerDoneResult_ResultSuccess ServerDoneResult = 1 // Indicating a successful authentication
	ServerDoneResult_ResultReject  ServerDoneResult = 2 // Indicating that the authentication was not successful
)

// Enum value maps for ServerDoneResult.
var (
	ServerDoneResult_name = map[int32]string{
		0: "ResultUnknown",
		1: "ResultSuccess",
		2: "ResultReject",
	}
	ServerDoneResult_value = map[string]int32{
		"ResultUnknown": 0,
		"ResultSuccess": 1,
		"ResultReject":  2,
	}
)

func (x ServerDoneResult) Enum() *ServerDoneResult {
	p := new(ServerDoneResult)
	*p = x
	return p
}

func (x ServerDoneResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerDoneResult) Descriptor() protoreflect.EnumDescriptor {
	return file_sasl_proto_enumTypes[1].Descriptor()
}

func (ServerDoneResult) Type() protoreflect.EnumType {
	return &file_sasl_proto_enumTypes[1]
}

func (x ServerDoneResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerDoneResult.Descriptor instead.
func (ServerDoneResult) EnumDescriptor() ([]byte, []int) {
	return file_sasl_proto_rawDescGZIP(), []int{1}
}

// based on section 3 in RFC4422
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of the message indicating which of the other fields
	// is set
	MessageType MessageType `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3,enum=saslproto.MessageType" json:"message_type,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*Message_ServerMechanismAdvertisement
	//	*Message_ClientInitiation
	//	*Message_ChallengeResponse
	//	*Message_HandshakeAbortion
	//	*Message_ServerDone
	Payload isMessage_Payload `protobuf_oneof:"payload"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sasl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_sasl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_sasl_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMessageType() MessageType {
	if x != nil {
		return x.MessageType
	}
	return MessageType_MessageTypeUnknown
}

func (m *Message) GetPayload() isMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Message) GetServerMechanismAdvertisement() *ServerMechanismAdvertisement {
	if x, ok := x.GetPayload().(*Message_ServerMechanismAdvertisement); ok {
		return x.ServerMechanismAdvertisement
	}
	return nil
}

func (x *Message) GetClientInitiation() *ClientInitiation {
	if x, ok := x.GetPayload().(*Message_ClientInitiation); ok {
		return x.ClientInitiation
	}
	return nil
}

func (x *Message) GetChallengeResponse() *ChallengeResponse {
	if x, ok := x.GetPayload().(*Message_ChallengeResponse); ok {
		return x.ChallengeResponse
	}
	return nil
}

func (x *Message) GetHandshakeAbortion() *HandshakeAbortion {
	if x, ok := x.GetPayload().(*Message_HandshakeAbortion); ok {
		return x.HandshakeAbortion
	}
	return nil
}

func (x *Message) GetServerDone() *ServerDone {
	if x, ok := x.GetPayload().(*Message_ServerDone); ok {
		return x.ServerDone
	}
	return nil
}

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_ServerMechanismAdvertisement struct {
	// Payload for message type MessageTypeServerMechanismAdvertisement
	ServerMechanismAdvertisement *ServerMechanismAdvertisement `protobuf:"bytes,2,opt,name=server_mechanism_advertisement,json=serverMechanismAdvertisement,proto3,oneof"`
}

type Message_ClientInitiation struct {
	// Payload for message type MessageTypeClientInitiation
	ClientInitiation *ClientInitiation `protobuf:"bytes,3,opt,name=client_initiation,json=clientInitiation,proto3,oneof"`
}

type Message_ChallengeResponse struct {
	// Payload for message type MessageTypeChallengeResponse
	ChallengeResponse *ChallengeResponse `protobuf:"bytes,4,opt,name=challenge_response,json=challengeResponse,proto3,oneof"`
}

type Message_HandshakeAbortion struct {
	// Payload for message type MessageTypeHandshakeAbortion
	HandshakeAbortion *HandshakeAbortion `protobuf:"bytes,5,opt,name=handshake_abortion,json=handshakeAbortion,proto3,oneof"`
}

type Message_ServerDone struct {
	// Payload for message type MessageTypeServerDone
	ServerDone *ServerDone `protobuf:"bytes,6,opt,name=server_done,json=serverDone,proto3,oneof"`
}

func (*Message_ServerMechanismAdvertisement) isMessage_Payload() {}

func (*Message_ClientInitiation) isMessage_Payload() {}

func (*Message_ChallengeResponse) isMessage_Payload() {}

func (*Message_HandshakeAbortion) isMessage_Payload() {}

func (*Message_ServerDone) isMessage_Payload() {}

// ServerMechanismAdvertisement is the first message of the handshake
// indicating which SASL mechanisms the server supports (see section 3.2)
type ServerMechanismAdvertisement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of supported SASL mechanisms
	Mechanisms []string `protobuf:"bytes,1,rep,name=mechanisms,proto3" json:"mechanisms,omitempty"`
}

func (x *ServerMechanismAdvertisement) Reset() {
	*x = ServerMechanismAdvertisement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sasl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMechanismAdvertisement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMechanismAdvertisement) ProtoMessage() {}

func (x *ServerMechanismAdvertisement) ProtoReflect() protoreflect.Message {
	mi := &file_sasl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMechanismAdvertisement.ProtoReflect.Descriptor instead.
func (*ServerMechanismAdvertisement) Descriptor() ([]byte, []int) {
	return file_sasl_proto_rawDescGZIP(), []int{1}
}

func (x *ServerMechanismAdvertisement) GetMechanisms() []string {
	if x != nil {
		return x.Mechanisms
	}
	return nil
}

// ClientInitiation initiates the SASL challenge-response exchange
// between client and server by indicating the by the client selected
// SASL mechanism and an optional initial response (see section 3.3)
type ClientInitiation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Selected SASL mechanism
	Mechanism string `protobuf:"bytes,1,opt,name=mechanism,proto3" json:"mechanism,omitempty"`
	// Used to distinguish between nil and empty initial responses
	// as described in section 4.3.b
	InitialReponseIsNil bool `protobuf:"varint,2,opt,name=initial_reponse_is_nil,json=initialReponseIsNil,proto3" json:"initial_reponse_is_nil,omitempty"`
	// Optional initial response sent by the client (depending on the mechanism)
	InitialResponse []byte `protobuf:"bytes,3,opt,name=initial_response,json=initialResponse,proto3" json:"initial_response,omitempty"`
}

func (x *ClientInitiation) Reset() {
	*x = ClientInitiation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sasl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInitiation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInitiation) ProtoMessage() {}

func (x *ClientInitiation) ProtoReflect() protoreflect.Message {
	mi := &file_sasl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInitiation.ProtoReflect.Descriptor instead.
func (*ClientInitiation) Descriptor() ([]byte, []int) {
	return file_sasl_proto_rawDescGZIP(), []int{2}
}

func (x *ClientInitiation) GetMechanism() string {
	if x != nil {
		return x.Mechanism
	}
	return ""
}

func (x *ClientInitiation) GetInitialReponseIsNil() bool {
	if x != nil {
		return x.InitialReponseIsNil
	}
	return false
}

func (x *ClientInitiation) GetInitialResponse() []byte {
	if x != nil {
		return x.InitialResponse
	}
	return nil
}

// ChallengeResponse is the message used during the challenge-response
// phase. It either contains a challenge sent by the server or a response
// by the client (see section 3.4).
type ChallengeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Payload in bytes
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ChallengeResponse) Reset() {
	*x = ChallengeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sasl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeResponse) ProtoMessage() {}

func (x *ChallengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sasl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeResponse.ProtoReflect.Descriptor instead.
func (*ChallengeResponse) Descriptor() ([]byte, []int) {
	return file_sasl_proto_rawDescGZIP(), []int{3}
}

func (x *ChallengeResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// HandshakeAbortion is used to abort an ongoing handshake (authentication exchange)
// as described in section 3.5
type HandshakeAbortion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional message which may be used to provide a abortion cause/ error
	// code for the client
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HandshakeAbortion) Reset() {
	*x = HandshakeAbortion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sasl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeAbortion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeAbortion) ProtoMessage() {}

func (x *HandshakeAbortion) ProtoReflect() protoreflect.Message {
	mi := &file_sasl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeAbortion.ProtoReflect.Descriptor instead.
func (*HandshakeAbortion) Descriptor() ([]byte, []int) {
	return file_sasl_proto_rawDescGZIP(), []int{4}
}

func (x *HandshakeAbortion) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// ServerDone is used for describing the authentication outcome
// and sent by the server after the client initiated a mechanism
// or the last challenge of a mechanism has been answered by the client
// (see section 3.6)
type ServerDone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result of the authentication mechanism which is either
	// success or reject
	Result ServerDoneResult `protobuf:"varint,1,opt,name=result,proto3,enum=saslproto.ServerDoneResult" json:"result,omitempty"`
	// Optional message which may be used to provide additional information
	// for the client
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ServerDone) Reset() {
	*x = ServerDone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sasl_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerDone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerDone) ProtoMessage() {}

func (x *ServerDone) ProtoReflect() protoreflect.Message {
	mi := &file_sasl_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerDone.ProtoReflect.Descriptor instead.
func (*ServerDone) Descriptor() ([]byte, []int) {
	return file_sasl_proto_rawDescGZIP(), []int{5}
}

func (x *ServerDone) GetResult() ServerDoneResult {
	if x != nil {
		return x.Result
	}
	return ServerDoneResult_ResultUnknown
}

func (x *ServerDone) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_sasl_proto protoreflect.FileDescriptor

var file_sasl_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x61, 0x73, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x61,
	0x73, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x03, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73, 0x61, 0x73, 0x6c,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x6f,
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69,
	0x73, 0x6d, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x61, 0x73, 0x6c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69,
	0x73, 0x6d, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x48,
	0x00, 0x52, 0x1c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69,
	0x73, 0x6d, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x4a, 0x0a, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x61, 0x73,
	0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x12, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x61, 0x73, 0x6c, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x11, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x12, 0x68, 0x61,
	0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x5f, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x61, 0x73, 0x6c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x41, 0x62, 0x6f, 0x72,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x11, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b,
	0x65, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x73, 0x61, 0x73, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x44, 0x6f, 0x6e, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44,
	0x6f, 0x6e, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x3e,
	0x0a, 0x1c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73,
	0x6d, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x73, 0x22, 0x90,
	0x01, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73,
	0x6d, 0x12, 0x33, 0x0a, 0x16, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x73, 0x5f, 0x6e, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x13, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x49, 0x73, 0x4e, 0x69, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2d, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x2d, 0x0a, 0x11, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x41, 0x62, 0x6f,
	0x72, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x5b, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x33, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x73, 0x61, 0x73, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0xd2, 0x01, 0x0a,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x10, 0x00, 0x12, 0x2b, 0x0a, 0x27, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69,
	0x73, 0x6d, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x10,
	0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x41, 0x62, 0x6f, 0x72,
	0x74, 0x69, 0x6f, 0x6e, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x6f, 0x6e, 0x65, 0x10,
	0x05, 0x2a, 0x4a, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x6f, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x10, 0x02, 0x42, 0x0d, 0x5a,
	0x0b, 0x2e, 0x3b, 0x73, 0x61, 0x73, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sasl_proto_rawDescOnce sync.Once
	file_sasl_proto_rawDescData = file_sasl_proto_rawDesc
)

func file_sasl_proto_rawDescGZIP() []byte {
	file_sasl_proto_rawDescOnce.Do(func() {
		file_sasl_proto_rawDescData = protoimpl.X.CompressGZIP(file_sasl_proto_rawDescData)
	})
	return file_sasl_proto_rawDescData
}

var file_sasl_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_sasl_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sasl_proto_goTypes = []interface{}{
	(MessageType)(0),                     // 0: saslproto.MessageType
	(ServerDoneResult)(0),                // 1: saslproto.ServerDoneResult
	(*Message)(nil),                      // 2: saslproto.Message
	(*ServerMechanismAdvertisement)(nil), // 3: saslproto.ServerMechanismAdvertisement
	(*ClientInitiation)(nil),             // 4: saslproto.ClientInitiation
	(*ChallengeResponse)(nil),            // 5: saslproto.ChallengeResponse
	(*HandshakeAbortion)(nil),            // 6: saslproto.HandshakeAbortion
	(*ServerDone)(nil),                   // 7: saslproto.ServerDone
}
var file_sasl_proto_depIdxs = []int32{
	0, // 0: saslproto.Message.message_type:type_name -> saslproto.MessageType
	3, // 1: saslproto.Message.server_mechanism_advertisement:type_name -> saslproto.ServerMechanismAdvertisement
	4, // 2: saslproto.Message.client_initiation:type_name -> saslproto.ClientInitiation
	5, // 3: saslproto.Message.challenge_response:type_name -> saslproto.ChallengeResponse
	6, // 4: saslproto.Message.handshake_abortion:type_name -> saslproto.HandshakeAbortion
	7, // 5: saslproto.Message.server_done:type_name -> saslproto.ServerDone
	1, // 6: saslproto.ServerDone.result:type_name -> saslproto.ServerDoneResult
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_sasl_proto_init() }
func file_sasl_proto_init() {
	if File_sasl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sasl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sasl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMechanismAdvertisement); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sasl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInitiation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sasl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sasl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeAbortion); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sasl_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerDone); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_sasl_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_ServerMechanismAdvertisement)(nil),
		(*Message_ClientInitiation)(nil),
		(*Message_ChallengeResponse)(nil),
		(*Message_HandshakeAbortion)(nil),
		(*Message_ServerDone)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sasl_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sasl_proto_goTypes,
		DependencyIndexes: file_sasl_proto_depIdxs,
		EnumInfos:         file_sasl_proto_enumTypes,
		MessageInfos:      file_sasl_proto_msgTypes,
	}.Build()
	File_sasl_proto = out.File
	file_sasl_proto_rawDesc = nil
	file_sasl_proto_goTypes = nil
	file_sasl_proto_depIdxs = nil
}