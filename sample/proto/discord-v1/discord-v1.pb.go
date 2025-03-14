// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: proto/discord-v1/discord-v1.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Only supported for Guilds
type ChatMessage struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ChannelId       string                 `protobuf:"bytes,2,opt,name=channelId,proto3" json:"channelId,omitempty"`
	GuildId         string                 `protobuf:"bytes,3,opt,name=guildId,proto3" json:"guildId,omitempty"`
	Content         string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	EditedTimestamp *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=edited_timestamp,json=editedTimestamp,proto3,oneof" json:"edited_timestamp,omitempty"`
	MentionRoles    []string               `protobuf:"bytes,7,rep,name=MentionRoles,proto3" json:"MentionRoles,omitempty"`
	Tts             bool                   `protobuf:"varint,8,opt,name=tts,proto3" json:"tts,omitempty"`
	MentionEveryone bool                   `protobuf:"varint,9,opt,name=mention_everyone,json=mentionEveryone,proto3" json:"mention_everyone,omitempty"`
	MessageFlags    int32                  `protobuf:"varint,10,opt,name=message_flags,json=messageFlags,proto3" json:"message_flags,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_proto_discord_v1_discord_v1_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChatMessage) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ChatMessage) GetGuildId() string {
	if x != nil {
		return x.GuildId
	}
	return ""
}

func (x *ChatMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ChatMessage) GetEditedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.EditedTimestamp
	}
	return nil
}

func (x *ChatMessage) GetMentionRoles() []string {
	if x != nil {
		return x.MentionRoles
	}
	return nil
}

func (x *ChatMessage) GetTts() bool {
	if x != nil {
		return x.Tts
	}
	return false
}

func (x *ChatMessage) GetMentionEveryone() bool {
	if x != nil {
		return x.MentionEveryone
	}
	return false
}

func (x *ChatMessage) GetMessageFlags() int32 {
	if x != nil {
		return x.MessageFlags
	}
	return 0
}

type InitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Interactions  []string               `protobuf:"bytes,1,rep,name=interactions,proto3" json:"interactions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_proto_discord_v1_discord_v1_proto_rawDescGZIP(), []int{1}
}

func (x *InitResponse) GetInteractions() []string {
	if x != nil {
		return x.Interactions
	}
	return nil
}

type ChatIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChatId        string                 `protobuf:"bytes,1,opt,name=chatId,proto3" json:"chatId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatIdResponse) Reset() {
	*x = ChatIdResponse{}
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatIdResponse) ProtoMessage() {}

func (x *ChatIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatIdResponse.ProtoReflect.Descriptor instead.
func (*ChatIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_discord_v1_discord_v1_proto_rawDescGZIP(), []int{2}
}

func (x *ChatIdResponse) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type ChatResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GuildId       string                 `protobuf:"bytes,1,opt,name=guildId,proto3" json:"guildId,omitempty"`
	ChannelId     string                 `protobuf:"bytes,2,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatResponse) Reset() {
	*x = ChatResponse{}
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResponse) ProtoMessage() {}

func (x *ChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResponse.ProtoReflect.Descriptor instead.
func (*ChatResponse) Descriptor() ([]byte, []int) {
	return file_proto_discord_v1_discord_v1_proto_rawDescGZIP(), []int{3}
}

func (x *ChatResponse) GetGuildId() string {
	if x != nil {
		return x.GuildId
	}
	return ""
}

func (x *ChatResponse) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ChatResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type OnCreateInteractionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	GuildId       string                 `protobuf:"bytes,2,opt,name=guildId,proto3" json:"guildId,omitempty"`
	ChannelId     string                 `protobuf:"bytes,3,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Message       string                 `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OnCreateInteractionRequest) Reset() {
	*x = OnCreateInteractionRequest{}
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnCreateInteractionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnCreateInteractionRequest) ProtoMessage() {}

func (x *OnCreateInteractionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnCreateInteractionRequest.ProtoReflect.Descriptor instead.
func (*OnCreateInteractionRequest) Descriptor() ([]byte, []int) {
	return file_proto_discord_v1_discord_v1_proto_rawDescGZIP(), []int{4}
}

func (x *OnCreateInteractionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OnCreateInteractionRequest) GetGuildId() string {
	if x != nil {
		return x.GuildId
	}
	return ""
}

func (x *OnCreateInteractionRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *OnCreateInteractionRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResponseInteractionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GuildId       string                 `protobuf:"bytes,1,opt,name=guildId,proto3" json:"guildId,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResponseInteractionRequest) Reset() {
	*x = ResponseInteractionRequest{}
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseInteractionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseInteractionRequest) ProtoMessage() {}

func (x *ResponseInteractionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseInteractionRequest.ProtoReflect.Descriptor instead.
func (*ResponseInteractionRequest) Descriptor() ([]byte, []int) {
	return file_proto_discord_v1_discord_v1_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseInteractionRequest) GetGuildId() string {
	if x != nil {
		return x.GuildId
	}
	return ""
}

func (x *ResponseInteractionRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EditInteractionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InteractionId string                 `protobuf:"bytes,1,opt,name=interactionId,proto3" json:"interactionId,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EditInteractionRequest) Reset() {
	*x = EditInteractionRequest{}
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditInteractionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditInteractionRequest) ProtoMessage() {}

func (x *EditInteractionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditInteractionRequest.ProtoReflect.Descriptor instead.
func (*EditInteractionRequest) Descriptor() ([]byte, []int) {
	return file_proto_discord_v1_discord_v1_proto_rawDescGZIP(), []int{6}
}

func (x *EditInteractionRequest) GetInteractionId() string {
	if x != nil {
		return x.InteractionId
	}
	return ""
}

func (x *EditInteractionRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type OnEventRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Event         string                 `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OnEventRequest) Reset() {
	*x = OnEventRequest{}
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnEventRequest) ProtoMessage() {}

func (x *OnEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_v1_discord_v1_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnEventRequest.ProtoReflect.Descriptor instead.
func (*OnEventRequest) Descriptor() ([]byte, []int) {
	return file_proto_discord_v1_discord_v1_proto_rawDescGZIP(), []int{7}
}

func (x *OnEventRequest) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

var File_proto_discord_v1_discord_v1_proto protoreflect.FileDescriptor

var file_proto_discord_v1_discord_v1_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2d,
	0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2d, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x03, 0x0a, 0x0b, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x4a, 0x0a, 0x10, 0x65, 0x64, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0f, 0x65, 0x64,
	0x69, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x88, 0x01, 0x01,
	0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x4d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x74, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x65, 0x76, 0x65, 0x72, 0x79, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x72, 0x79, 0x6f, 0x6e,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x32, 0x0a, 0x0c, 0x49,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x28, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x0c, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c,
	0x64, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x7e, 0x0a, 0x1a, 0x4f,
	0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c,
	0x64, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x1a, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c,
	0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x58, 0x0a,
	0x16, 0x45, 0x64, 0x69, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x0e, 0x4f, 0x6e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32,
	0x9f, 0x04, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x35, 0x0a, 0x06, 0x4f,
	0x6e, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0f, 0x4f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x50, 0x0a, 0x13, 0x4f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x07, 0x4f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a,
	0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x13, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0b,
	0x45, 0x64, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x0f, 0x45, 0x64, 0x69, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_discord_v1_discord_v1_proto_rawDescOnce sync.Once
	file_proto_discord_v1_discord_v1_proto_rawDescData []byte
)

func file_proto_discord_v1_discord_v1_proto_rawDescGZIP() []byte {
	file_proto_discord_v1_discord_v1_proto_rawDescOnce.Do(func() {
		file_proto_discord_v1_discord_v1_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_discord_v1_discord_v1_proto_rawDesc), len(file_proto_discord_v1_discord_v1_proto_rawDesc)))
	})
	return file_proto_discord_v1_discord_v1_proto_rawDescData
}

var file_proto_discord_v1_discord_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_discord_v1_discord_v1_proto_goTypes = []any{
	(*ChatMessage)(nil),                // 0: proto.ChatMessage
	(*InitResponse)(nil),               // 1: proto.InitResponse
	(*ChatIdResponse)(nil),             // 2: proto.ChatIdResponse
	(*ChatResponse)(nil),               // 3: proto.ChatResponse
	(*OnCreateInteractionRequest)(nil), // 4: proto.OnCreateInteractionRequest
	(*ResponseInteractionRequest)(nil), // 5: proto.ResponseInteractionRequest
	(*EditInteractionRequest)(nil),     // 6: proto.EditInteractionRequest
	(*OnEventRequest)(nil),             // 7: proto.OnEventRequest
	(*timestamppb.Timestamp)(nil),      // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),              // 9: google.protobuf.Empty
}
var file_proto_discord_v1_discord_v1_proto_depIdxs = []int32{
	8,  // 0: proto.ChatMessage.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 1: proto.ChatMessage.edited_timestamp:type_name -> google.protobuf.Timestamp
	9,  // 2: proto.Discord.OnInit:input_type -> google.protobuf.Empty
	0,  // 3: proto.Discord.OnCreateMessage:input_type -> proto.ChatMessage
	4,  // 4: proto.Discord.OnCreateInteraction:input_type -> proto.OnCreateInteractionRequest
	7,  // 5: proto.Discord.OnEvent:input_type -> proto.OnEventRequest
	0,  // 6: proto.Discord.ResponseMessage:input_type -> proto.ChatMessage
	5,  // 7: proto.Discord.ResponseInteraction:input_type -> proto.ResponseInteractionRequest
	0,  // 8: proto.Discord.EditMessage:input_type -> proto.ChatMessage
	6,  // 9: proto.Discord.EditInteraction:input_type -> proto.EditInteractionRequest
	1,  // 10: proto.Discord.OnInit:output_type -> proto.InitResponse
	9,  // 11: proto.Discord.OnCreateMessage:output_type -> google.protobuf.Empty
	9,  // 12: proto.Discord.OnCreateInteraction:output_type -> google.protobuf.Empty
	9,  // 13: proto.Discord.OnEvent:output_type -> google.protobuf.Empty
	2,  // 14: proto.Discord.ResponseMessage:output_type -> proto.ChatIdResponse
	2,  // 15: proto.Discord.ResponseInteraction:output_type -> proto.ChatIdResponse
	9,  // 16: proto.Discord.EditMessage:output_type -> google.protobuf.Empty
	9,  // 17: proto.Discord.EditInteraction:output_type -> google.protobuf.Empty
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_discord_v1_discord_v1_proto_init() }
func file_proto_discord_v1_discord_v1_proto_init() {
	if File_proto_discord_v1_discord_v1_proto != nil {
		return
	}
	file_proto_discord_v1_discord_v1_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_discord_v1_discord_v1_proto_rawDesc), len(file_proto_discord_v1_discord_v1_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_discord_v1_discord_v1_proto_goTypes,
		DependencyIndexes: file_proto_discord_v1_discord_v1_proto_depIdxs,
		MessageInfos:      file_proto_discord_v1_discord_v1_proto_msgTypes,
	}.Build()
	File_proto_discord_v1_discord_v1_proto = out.File
	file_proto_discord_v1_discord_v1_proto_goTypes = nil
	file_proto_discord_v1_discord_v1_proto_depIdxs = nil
}
