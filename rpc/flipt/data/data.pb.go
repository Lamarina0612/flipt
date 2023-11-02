// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: data/data.proto

package data

import (
	flipt "go.flipt.io/flipt/rpc/flipt"
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

type SnapshotNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *SnapshotNamespaceRequest) Reset() {
	*x = SnapshotNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotNamespaceRequest) ProtoMessage() {}

func (x *SnapshotNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotNamespaceRequest.ProtoReflect.Descriptor instead.
func (*SnapshotNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{0}
}

func (x *SnapshotNamespaceRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type SnapshotNamespaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string                        `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Flags    []*flipt.Flag                 `protobuf:"bytes,2,rep,name=flags,proto3" json:"flags,omitempty"`
	Segments []*flipt.Segment              `protobuf:"bytes,3,rep,name=segments,proto3" json:"segments,omitempty"`
	Rules    map[string]*flipt.RuleList    `protobuf:"bytes,4,rep,name=rules,proto3" json:"rules,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Rollouts map[string]*flipt.RolloutList `protobuf:"bytes,5,rep,name=rollouts,proto3" json:"rollouts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SnapshotNamespaceResponse) Reset() {
	*x = SnapshotNamespaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotNamespaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotNamespaceResponse) ProtoMessage() {}

func (x *SnapshotNamespaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotNamespaceResponse.ProtoReflect.Descriptor instead.
func (*SnapshotNamespaceResponse) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{1}
}

func (x *SnapshotNamespaceResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SnapshotNamespaceResponse) GetFlags() []*flipt.Flag {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *SnapshotNamespaceResponse) GetSegments() []*flipt.Segment {
	if x != nil {
		return x.Segments
	}
	return nil
}

func (x *SnapshotNamespaceResponse) GetRules() map[string]*flipt.RuleList {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *SnapshotNamespaceResponse) GetRollouts() map[string]*flipt.RolloutList {
	if x != nil {
		return x.Rollouts
	}
	return nil
}

var File_data_data_proto protoreflect.FileDescriptor

var file_data_data_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x0b, 0x66,
	0x6c, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x18, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xb1, 0x03, 0x0a, 0x19, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2e,
	0x46, 0x6c, 0x61, 0x67, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x66, 0x6c, 0x69, 0x70, 0x74, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x46, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x75,
	0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12,
	0x4f, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x73,
	0x1a, 0x49, 0x0a, 0x0a, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4f, 0x0a, 0x0d, 0x52,
	0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x66, 0x6c, 0x69, 0x70, 0x74, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x6f, 0x0a, 0x0b,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x24, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x22, 0x5a,
	0x20, 0x67, 0x6f, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2e, 0x69, 0x6f, 0x2f, 0x66, 0x6c, 0x69,
	0x70, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_data_proto_rawDescOnce sync.Once
	file_data_data_proto_rawDescData = file_data_data_proto_rawDesc
)

func file_data_data_proto_rawDescGZIP() []byte {
	file_data_data_proto_rawDescOnce.Do(func() {
		file_data_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_data_proto_rawDescData)
	})
	return file_data_data_proto_rawDescData
}

var file_data_data_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_data_data_proto_goTypes = []interface{}{
	(*SnapshotNamespaceRequest)(nil),  // 0: flipt.data.SnapshotNamespaceRequest
	(*SnapshotNamespaceResponse)(nil), // 1: flipt.data.SnapshotNamespaceResponse
	nil,                               // 2: flipt.data.SnapshotNamespaceResponse.RulesEntry
	nil,                               // 3: flipt.data.SnapshotNamespaceResponse.RolloutsEntry
	(*flipt.Flag)(nil),                // 4: flipt.Flag
	(*flipt.Segment)(nil),             // 5: flipt.Segment
	(*flipt.RuleList)(nil),            // 6: flipt.RuleList
	(*flipt.RolloutList)(nil),         // 7: flipt.RolloutList
}
var file_data_data_proto_depIdxs = []int32{
	4, // 0: flipt.data.SnapshotNamespaceResponse.flags:type_name -> flipt.Flag
	5, // 1: flipt.data.SnapshotNamespaceResponse.segments:type_name -> flipt.Segment
	2, // 2: flipt.data.SnapshotNamespaceResponse.rules:type_name -> flipt.data.SnapshotNamespaceResponse.RulesEntry
	3, // 3: flipt.data.SnapshotNamespaceResponse.rollouts:type_name -> flipt.data.SnapshotNamespaceResponse.RolloutsEntry
	6, // 4: flipt.data.SnapshotNamespaceResponse.RulesEntry.value:type_name -> flipt.RuleList
	7, // 5: flipt.data.SnapshotNamespaceResponse.RolloutsEntry.value:type_name -> flipt.RolloutList
	0, // 6: flipt.data.DataService.SnapshotNamespace:input_type -> flipt.data.SnapshotNamespaceRequest
	1, // 7: flipt.data.DataService.SnapshotNamespace:output_type -> flipt.data.SnapshotNamespaceResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_data_data_proto_init() }
func file_data_data_proto_init() {
	if File_data_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotNamespaceRequest); i {
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
		file_data_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotNamespaceResponse); i {
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
			RawDescriptor: file_data_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_data_data_proto_goTypes,
		DependencyIndexes: file_data_data_proto_depIdxs,
		MessageInfos:      file_data_data_proto_msgTypes,
	}.Build()
	File_data_data_proto = out.File
	file_data_data_proto_rawDesc = nil
	file_data_data_proto_goTypes = nil
	file_data_data_proto_depIdxs = nil
}
