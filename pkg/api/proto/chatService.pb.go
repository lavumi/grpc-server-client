// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: api/proto/chatService.proto

package proto

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

type RoomRequest_Action int32

const (
	RoomRequest_CREATE RoomRequest_Action = 0
	RoomRequest_JOIN   RoomRequest_Action = 1
	RoomRequest_LEAVE  RoomRequest_Action = 2
)

// Enum value maps for RoomRequest_Action.
var (
	RoomRequest_Action_name = map[int32]string{
		0: "CREATE",
		1: "JOIN",
		2: "LEAVE",
	}
	RoomRequest_Action_value = map[string]int32{
		"CREATE": 0,
		"JOIN":   1,
		"LEAVE":  2,
	}
)

func (x RoomRequest_Action) Enum() *RoomRequest_Action {
	p := new(RoomRequest_Action)
	*p = x
	return p
}

func (x RoomRequest_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomRequest_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_chatService_proto_enumTypes[0].Descriptor()
}

func (RoomRequest_Action) Type() protoreflect.EnumType {
	return &file_api_proto_chatService_proto_enumTypes[0]
}

func (x RoomRequest_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomRequest_Action.Descriptor instead.
func (RoomRequest_Action) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_chatService_proto_rawDescGZIP(), []int{3, 0}
}

type RoomResponse_Status int32

const (
	RoomResponse_SUCCESS RoomResponse_Status = 0
	RoomResponse_FAILURE RoomResponse_Status = 1
)

// Enum value maps for RoomResponse_Status.
var (
	RoomResponse_Status_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILURE",
	}
	RoomResponse_Status_value = map[string]int32{
		"SUCCESS": 0,
		"FAILURE": 1,
	}
)

func (x RoomResponse_Status) Enum() *RoomResponse_Status {
	p := new(RoomResponse_Status)
	*p = x
	return p
}

func (x RoomResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_chatService_proto_enumTypes[1].Descriptor()
}

func (RoomResponse_Status) Type() protoreflect.EnumType {
	return &file_api_proto_chatService_proto_enumTypes[1]
}

func (x RoomResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomResponse_Status.Descriptor instead.
func (RoomResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_chatService_proto_rawDescGZIP(), []int{4, 0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_chatService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_chatService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_proto_chatService_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId    string `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	User      *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp string `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_chatService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_chatService_proto_msgTypes[1]
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
	return file_api_proto_chatService_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *Message) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Message) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId   string     `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Name     string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Members  []*User    `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	Messages []*Message `protobuf:"bytes,4,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_chatService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_chatService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_api_proto_chatService_proto_rawDescGZIP(), []int{2}
}

func (x *Room) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Room) GetMembers() []*User {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Room) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

// 채팅룸 관련 요청 메시지
type RoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   *User              `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Action RoomRequest_Action `protobuf:"varint,2,opt,name=action,proto3,enum=room.RoomRequest_Action" json:"action,omitempty"`
	Room   *Room              `protobuf:"bytes,3,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *RoomRequest) Reset() {
	*x = RoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_chatService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomRequest) ProtoMessage() {}

func (x *RoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_chatService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomRequest.ProtoReflect.Descriptor instead.
func (*RoomRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_chatService_proto_rawDescGZIP(), []int{3}
}

func (x *RoomRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *RoomRequest) GetAction() RoomRequest_Action {
	if x != nil {
		return x.Action
	}
	return RoomRequest_CREATE
}

func (x *RoomRequest) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

// 채팅룸 관련 응답 메시지
type RoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  RoomResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=room.RoomResponse_Status" json:"status,omitempty"`
	Message string              `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Room    *Room               `protobuf:"bytes,3,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *RoomResponse) Reset() {
	*x = RoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_chatService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomResponse) ProtoMessage() {}

func (x *RoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_chatService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomResponse.ProtoReflect.Descriptor instead.
func (*RoomResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_chatService_proto_rawDescGZIP(), []int{4}
}

func (x *RoomResponse) GetStatus() RoomResponse_Status {
	if x != nil {
		return x.Status
	}
	return RoomResponse_SUCCESS
}

func (x *RoomResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RoomResponse) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type RoomList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomList []*Room `protobuf:"bytes,1,rep,name=room_list,json=roomList,proto3" json:"room_list,omitempty"`
}

func (x *RoomList) Reset() {
	*x = RoomList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_chatService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomList) ProtoMessage() {}

func (x *RoomList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_chatService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomList.ProtoReflect.Descriptor instead.
func (*RoomList) Descriptor() ([]byte, []int) {
	return file_api_proto_chatService_proto_rawDescGZIP(), []int{5}
}

func (x *RoomList) GetRoomList() []*Room {
	if x != nil {
		return x.RoomList
	}
	return nil
}

type Close struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Close) Reset() {
	*x = Close{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_chatService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Close) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Close) ProtoMessage() {}

func (x *Close) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_chatService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Close.ProtoReflect.Descriptor instead.
func (*Close) Descriptor() ([]byte, []int) {
	return file_api_proto_chatService_proto_rawDescGZIP(), []int{6}
}

var File_api_proto_chatService_proto protoreflect.FileDescriptor

var file_api_proto_chatService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72,
	0x6f, 0x6f, 0x6d, 0x22, 0x41, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x79, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x83, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x6f, 0x6f,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x29, 0x0a, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x4a, 0x4f, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x45, 0x41,
	0x56, 0x45, 0x10, 0x02, 0x22, 0x9f, 0x01, 0x0a, 0x0c, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x22, 0x22, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x10, 0x01, 0x22, 0x33, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x32, 0xbf, 0x02, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x11, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x11, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x6f, 0x6f,
	0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x11,
	0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0d, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x30, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0d, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12,
	0x11, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_chatService_proto_rawDescOnce sync.Once
	file_api_proto_chatService_proto_rawDescData = file_api_proto_chatService_proto_rawDesc
)

func file_api_proto_chatService_proto_rawDescGZIP() []byte {
	file_api_proto_chatService_proto_rawDescOnce.Do(func() {
		file_api_proto_chatService_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_chatService_proto_rawDescData)
	})
	return file_api_proto_chatService_proto_rawDescData
}

var file_api_proto_chatService_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_proto_chatService_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_chatService_proto_goTypes = []interface{}{
	(RoomRequest_Action)(0),  // 0: room.RoomRequest.Action
	(RoomResponse_Status)(0), // 1: room.RoomResponse.Status
	(*User)(nil),             // 2: room.User
	(*Message)(nil),          // 3: room.Message
	(*Room)(nil),             // 4: room.Room
	(*RoomRequest)(nil),      // 5: room.RoomRequest
	(*RoomResponse)(nil),     // 6: room.RoomResponse
	(*RoomList)(nil),         // 7: room.RoomList
	(*Close)(nil),            // 8: room.Close
}
var file_api_proto_chatService_proto_depIdxs = []int32{
	2,  // 0: room.Message.user:type_name -> room.User
	2,  // 1: room.Room.members:type_name -> room.User
	3,  // 2: room.Room.messages:type_name -> room.Message
	2,  // 3: room.RoomRequest.user:type_name -> room.User
	0,  // 4: room.RoomRequest.action:type_name -> room.RoomRequest.Action
	4,  // 5: room.RoomRequest.room:type_name -> room.Room
	1,  // 6: room.RoomResponse.status:type_name -> room.RoomResponse.Status
	4,  // 7: room.RoomResponse.room:type_name -> room.Room
	4,  // 8: room.RoomList.room_list:type_name -> room.Room
	5,  // 9: room.RoomService.GetRoomList:input_type -> room.RoomRequest
	5,  // 10: room.RoomService.CreateRoom:input_type -> room.RoomRequest
	5,  // 11: room.RoomService.JoinRoom:input_type -> room.RoomRequest
	5,  // 12: room.RoomService.CreateStream:input_type -> room.RoomRequest
	3,  // 13: room.RoomService.BroadcastMessage:input_type -> room.Message
	5,  // 14: room.RoomService.LeaveRoom:input_type -> room.RoomRequest
	7,  // 15: room.RoomService.GetRoomList:output_type -> room.RoomList
	6,  // 16: room.RoomService.CreateRoom:output_type -> room.RoomResponse
	6,  // 17: room.RoomService.JoinRoom:output_type -> room.RoomResponse
	3,  // 18: room.RoomService.CreateStream:output_type -> room.Message
	8,  // 19: room.RoomService.BroadcastMessage:output_type -> room.Close
	6,  // 20: room.RoomService.LeaveRoom:output_type -> room.RoomResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_proto_chatService_proto_init() }
func file_api_proto_chatService_proto_init() {
	if File_api_proto_chatService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_chatService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_api_proto_chatService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_proto_chatService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_api_proto_chatService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomRequest); i {
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
		file_api_proto_chatService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomResponse); i {
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
		file_api_proto_chatService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomList); i {
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
		file_api_proto_chatService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Close); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_chatService_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_chatService_proto_goTypes,
		DependencyIndexes: file_api_proto_chatService_proto_depIdxs,
		EnumInfos:         file_api_proto_chatService_proto_enumTypes,
		MessageInfos:      file_api_proto_chatService_proto_msgTypes,
	}.Build()
	File_api_proto_chatService_proto = out.File
	file_api_proto_chatService_proto_rawDesc = nil
	file_api_proto_chatService_proto_goTypes = nil
	file_api_proto_chatService_proto_depIdxs = nil
}