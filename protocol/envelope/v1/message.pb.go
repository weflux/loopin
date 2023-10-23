// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.1
// source: envelope/v1/message.proto

package envelope

import (
	shared "github.com/weflux/loopin/protocol/shared"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are assignable to Body:
	//
	//	*Message_Request
	//	*Message_Reply
	//	*Message_Publication
	//	*Message_Error
	Body isMessage_Body `protobuf_oneof:"body"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envelope_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_envelope_v1_message_proto_msgTypes[0]
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
	return file_envelope_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (m *Message) GetBody() isMessage_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Message) GetRequest() *Request {
	if x, ok := x.GetBody().(*Message_Request); ok {
		return x.Request
	}
	return nil
}

func (x *Message) GetReply() *Reply {
	if x, ok := x.GetBody().(*Message_Reply); ok {
		return x.Reply
	}
	return nil
}

func (x *Message) GetPublication() *Publication {
	if x, ok := x.GetBody().(*Message_Publication); ok {
		return x.Publication
	}
	return nil
}

func (x *Message) GetError() *Error {
	if x, ok := x.GetBody().(*Message_Error); ok {
		return x.Error
	}
	return nil
}

type isMessage_Body interface {
	isMessage_Body()
}

type Message_Request struct {
	Request *Request `protobuf:"bytes,4,opt,name=request,proto3,oneof"`
}

type Message_Reply struct {
	Reply *Reply `protobuf:"bytes,5,opt,name=reply,proto3,oneof"`
}

type Message_Publication struct {
	Publication *Publication `protobuf:"bytes,6,opt,name=publication,proto3,oneof"`
}

type Message_Error struct {
	Error *Error `protobuf:"bytes,10,opt,name=error,proto3,oneof"`
}

func (*Message_Request) isMessage_Body() {}

func (*Message_Reply) isMessage_Body() {}

func (*Message_Publication) isMessage_Body() {}

func (*Message_Error) isMessage_Body() {}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command      string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	ContentType  string `protobuf:"bytes,2,opt,name=content_type,proto3" json:"content_type,omitempty"`
	PayloadText  string `protobuf:"bytes,3,opt,name=payload_text,proto3" json:"payload_text,omitempty"`
	PayloadBytes []byte `protobuf:"bytes,4,opt,name=payload_bytes,proto3" json:"payload_bytes,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envelope_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_envelope_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_envelope_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Request) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Request) GetPayloadText() string {
	if x != nil {
		return x.PayloadText
	}
	return ""
}

func (x *Request) GetPayloadBytes() []byte {
	if x != nil {
		return x.PayloadBytes
	}
	return nil
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error        *Error           `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Metadata     *shared.Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Command      string           `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	ContentType  string           `protobuf:"bytes,4,opt,name=content_type,proto3" json:"content_type,omitempty"`
	PayloadText  string           `protobuf:"bytes,5,opt,name=payload_text,proto3" json:"payload_text,omitempty"`
	PayloadBytes []byte           `protobuf:"bytes,6,opt,name=payload_bytes,proto3" json:"payload_bytes,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envelope_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_envelope_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_envelope_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *Reply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Reply) GetMetadata() *shared.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Reply) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Reply) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Reply) GetPayloadText() string {
	if x != nil {
		return x.PayloadText
	}
	return ""
}

func (x *Reply) GetPayloadBytes() []byte {
	if x != nil {
		return x.PayloadBytes
	}
	return nil
}

type SubRefresh struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubRefresh) Reset() {
	*x = SubRefresh{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envelope_v1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubRefresh) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubRefresh) ProtoMessage() {}

func (x *SubRefresh) ProtoReflect() protoreflect.Message {
	mi := &file_envelope_v1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubRefresh.ProtoReflect.Descriptor instead.
func (*SubRefresh) Descriptor() ([]byte, []int) {
	return file_envelope_v1_message_proto_rawDescGZIP(), []int{3}
}

type SubRefreshReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubRefreshReply) Reset() {
	*x = SubRefreshReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envelope_v1_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubRefreshReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubRefreshReply) ProtoMessage() {}

func (x *SubRefreshReply) ProtoReflect() protoreflect.Message {
	mi := &file_envelope_v1_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubRefreshReply.ProtoReflect.Descriptor instead.
func (*SubRefreshReply) Descriptor() ([]byte, []int) {
	return file_envelope_v1_message_proto_rawDescGZIP(), []int{4}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Extras  map[string]string `protobuf:"bytes,3,rep,name=extras,proto3" json:"extras,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envelope_v1_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_envelope_v1_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_envelope_v1_message_proto_rawDescGZIP(), []int{5}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetExtras() map[string]string {
	if x != nil {
		return x.Extras
	}
	return nil
}

type Publication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	ContentType  string `protobuf:"bytes,2,opt,name=content_type,proto3" json:"content_type,omitempty"`
	PayloadText  string `protobuf:"bytes,3,opt,name=payload_text,proto3" json:"payload_text,omitempty"`
	PayloadBytes []byte `protobuf:"bytes,4,opt,name=payload_bytes,proto3" json:"payload_bytes,omitempty"`
}

func (x *Publication) Reset() {
	*x = Publication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envelope_v1_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Publication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Publication) ProtoMessage() {}

func (x *Publication) ProtoReflect() protoreflect.Message {
	mi := &file_envelope_v1_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Publication.ProtoReflect.Descriptor instead.
func (*Publication) Descriptor() ([]byte, []int) {
	return file_envelope_v1_message_proto_rawDescGZIP(), []int{6}
}

func (x *Publication) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Publication) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Publication) GetPayloadText() string {
	if x != nil {
		return x.PayloadText
	}
	return ""
}

func (x *Publication) GetPayloadBytes() []byte {
	if x != nil {
		return x.PayloadBytes
	}
	return nil
}

var File_envelope_v1_message_proto protoreflect.FileDescriptor

var file_envelope_v1_message_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6c, 0x6f, 0x6f,
	0x70, 0x69, 0x6e, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x1a, 0x13, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf6, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3f, 0x0a,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x6e, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x34,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x6e, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x6e, 0x2e, 0x65, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x00, 0x52, 0x05, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x6f, 0x6f, 0x70,
	0x69, 0x6e, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x6e, 0x2e, 0x65,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x91, 0x01, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0xf2,
	0x01, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x6e,
	0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x69,
	0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0xac, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x06,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c,
	0x6f, 0x6f, 0x70, 0x69, 0x6e, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x65, 0x78, 0x74, 0x72, 0x61, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x42, 0x1f, 0x5a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envelope_v1_message_proto_rawDescOnce sync.Once
	file_envelope_v1_message_proto_rawDescData = file_envelope_v1_message_proto_rawDesc
)

func file_envelope_v1_message_proto_rawDescGZIP() []byte {
	file_envelope_v1_message_proto_rawDescOnce.Do(func() {
		file_envelope_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_envelope_v1_message_proto_rawDescData)
	})
	return file_envelope_v1_message_proto_rawDescData
}

var file_envelope_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_envelope_v1_message_proto_goTypes = []interface{}{
	(*Message)(nil),         // 0: loopin.envelope.Message
	(*Request)(nil),         // 1: loopin.envelope.Request
	(*Reply)(nil),           // 2: loopin.envelope.Reply
	(*SubRefresh)(nil),      // 3: loopin.envelope.SubRefresh
	(*SubRefreshReply)(nil), // 4: loopin.envelope.SubRefreshReply
	(*Error)(nil),           // 5: loopin.envelope.Error
	(*Publication)(nil),     // 6: loopin.envelope.Publication
	nil,                     // 7: loopin.envelope.Message.HeadersEntry
	nil,                     // 8: loopin.envelope.Error.ExtrasEntry
	(*shared.Metadata)(nil), // 9: loopin.shared.Metadata
}
var file_envelope_v1_message_proto_depIdxs = []int32{
	7, // 0: loopin.envelope.Message.headers:type_name -> loopin.envelope.Message.HeadersEntry
	1, // 1: loopin.envelope.Message.request:type_name -> loopin.envelope.Request
	2, // 2: loopin.envelope.Message.reply:type_name -> loopin.envelope.Reply
	6, // 3: loopin.envelope.Message.publication:type_name -> loopin.envelope.Publication
	5, // 4: loopin.envelope.Message.error:type_name -> loopin.envelope.Error
	5, // 5: loopin.envelope.Reply.error:type_name -> loopin.envelope.Error
	9, // 6: loopin.envelope.Reply.metadata:type_name -> loopin.shared.Metadata
	8, // 7: loopin.envelope.Error.extras:type_name -> loopin.envelope.Error.ExtrasEntry
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_envelope_v1_message_proto_init() }
func file_envelope_v1_message_proto_init() {
	if File_envelope_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envelope_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_envelope_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_envelope_v1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_envelope_v1_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubRefresh); i {
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
		file_envelope_v1_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubRefreshReply); i {
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
		file_envelope_v1_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_envelope_v1_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Publication); i {
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
	file_envelope_v1_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_Request)(nil),
		(*Message_Reply)(nil),
		(*Message_Publication)(nil),
		(*Message_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envelope_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envelope_v1_message_proto_goTypes,
		DependencyIndexes: file_envelope_v1_message_proto_depIdxs,
		MessageInfos:      file_envelope_v1_message_proto_msgTypes,
	}.Build()
	File_envelope_v1_message_proto = out.File
	file_envelope_v1_message_proto_rawDesc = nil
	file_envelope_v1_message_proto_goTypes = nil
	file_envelope_v1_message_proto_depIdxs = nil
}
