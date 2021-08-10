// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: pkg/events/login.proto

package events

import (
	messages "github.com/JungleMC/sdk/pkg/messages"
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

type PlayerLoginEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientType messages.ClientType `protobuf:"varint,1,opt,name=clientType,proto3,enum=messages.ClientType" json:"clientType,omitempty"`
	NetworkId  []byte              `protobuf:"bytes,2,opt,name=networkId,proto3" json:"networkId,omitempty"`
	ProfileId  []byte              `protobuf:"bytes,3,opt,name=profileId,proto3" json:"profileId,omitempty"`
	Username   string              `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	SkinUrl    string              `protobuf:"bytes,5,opt,name=skinUrl,proto3" json:"skinUrl,omitempty"`
	CapeUrl    string              `protobuf:"bytes,6,opt,name=capeUrl,proto3" json:"capeUrl,omitempty"`
}

func (x *PlayerLoginEvent) Reset() {
	*x = PlayerLoginEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_events_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLoginEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLoginEvent) ProtoMessage() {}

func (x *PlayerLoginEvent) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_events_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLoginEvent.ProtoReflect.Descriptor instead.
func (*PlayerLoginEvent) Descriptor() ([]byte, []int) {
	return file_pkg_events_login_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerLoginEvent) GetClientType() messages.ClientType {
	if x != nil {
		return x.ClientType
	}
	return messages.ClientType(0)
}

func (x *PlayerLoginEvent) GetNetworkId() []byte {
	if x != nil {
		return x.NetworkId
	}
	return nil
}

func (x *PlayerLoginEvent) GetProfileId() []byte {
	if x != nil {
		return x.ProfileId
	}
	return nil
}

func (x *PlayerLoginEvent) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PlayerLoginEvent) GetSkinUrl() string {
	if x != nil {
		return x.SkinUrl
	}
	return ""
}

func (x *PlayerLoginEvent) GetCapeUrl() string {
	if x != nil {
		return x.CapeUrl
	}
	return ""
}

var File_pkg_events_login_proto protoreflect.FileDescriptor

var file_pkg_events_login_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x1a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a,
	0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x6b, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70,
	0x65, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x65,
	0x55, 0x72, 0x6c, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4a, 0x75, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x43, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_events_login_proto_rawDescOnce sync.Once
	file_pkg_events_login_proto_rawDescData = file_pkg_events_login_proto_rawDesc
)

func file_pkg_events_login_proto_rawDescGZIP() []byte {
	file_pkg_events_login_proto_rawDescOnce.Do(func() {
		file_pkg_events_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_events_login_proto_rawDescData)
	})
	return file_pkg_events_login_proto_rawDescData
}

var file_pkg_events_login_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_events_login_proto_goTypes = []interface{}{
	(*PlayerLoginEvent)(nil), // 0: messages.PlayerLoginEvent
	(messages.ClientType)(0), // 1: messages.ClientType
}
var file_pkg_events_login_proto_depIdxs = []int32{
	1, // 0: messages.PlayerLoginEvent.clientType:type_name -> messages.ClientType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_events_login_proto_init() }
func file_pkg_events_login_proto_init() {
	if File_pkg_events_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_events_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLoginEvent); i {
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
			RawDescriptor: file_pkg_events_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_events_login_proto_goTypes,
		DependencyIndexes: file_pkg_events_login_proto_depIdxs,
		MessageInfos:      file_pkg_events_login_proto_msgTypes,
	}.Build()
	File_pkg_events_login_proto = out.File
	file_pkg_events_login_proto_rawDesc = nil
	file_pkg_events_login_proto_goTypes = nil
	file_pkg_events_login_proto_depIdxs = nil
}
