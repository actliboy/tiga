// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: utils/errorcode/errrep.proto

package errorcode

import (
	_ "github.com/liov/tiga/protobuf/utils/proto/patch"
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

type ErrRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ErrCode `protobuf:"varint,1,opt,name=code,proto3,enum=errorcode.ErrCode" json:"code"`
	Message string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrRep) Reset() {
	*x = ErrRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_utils_errorcode_errrep_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrRep) ProtoMessage() {}

func (x *ErrRep) ProtoReflect() protoreflect.Message {
	mi := &file_utils_errorcode_errrep_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrRep.ProtoReflect.Descriptor instead.
func (*ErrRep) Descriptor() ([]byte, []int) {
	return file_utils_errorcode_errrep_proto_rawDescGZIP(), []int{0}
}

func (x *ErrRep) GetCode() ErrCode {
	if x != nil {
		return x.Code
	}
	return SUCCESS
}

func (x *ErrRep) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_utils_errorcode_errrep_proto protoreflect.FileDescriptor

var file_utils_errorcode_errrep_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x2f, 0x65, 0x72, 0x72, 0x72, 0x65, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x1a, 0x75, 0x74, 0x69, 0x6c, 0x73,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x52, 0x65, 0x70, 0x12,
	0x3a, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x42, 0x12, 0xd2, 0xb5, 0x03, 0x0e, 0xa2, 0x01, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x51, 0x0a, 0x1c, 0x78, 0x79, 0x7a, 0x2e, 0x68, 0x6f, 0x70,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x69, 0x6f, 0x76, 0x2f, 0x68, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_utils_errorcode_errrep_proto_rawDescOnce sync.Once
	file_utils_errorcode_errrep_proto_rawDescData = file_utils_errorcode_errrep_proto_rawDesc
)

func file_utils_errorcode_errrep_proto_rawDescGZIP() []byte {
	file_utils_errorcode_errrep_proto_rawDescOnce.Do(func() {
		file_utils_errorcode_errrep_proto_rawDescData = protoimpl.X.CompressGZIP(file_utils_errorcode_errrep_proto_rawDescData)
	})
	return file_utils_errorcode_errrep_proto_rawDescData
}

var file_utils_errorcode_errrep_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_utils_errorcode_errrep_proto_goTypes = []interface{}{
	(*ErrRep)(nil), // 0: errorcode.ErrRep
	(ErrCode)(0),   // 1: errorcode.ErrCode
}
var file_utils_errorcode_errrep_proto_depIdxs = []int32{
	1, // 0: errorcode.ErrRep.code:type_name -> errorcode.ErrCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_utils_errorcode_errrep_proto_init() }
func file_utils_errorcode_errrep_proto_init() {
	if File_utils_errorcode_errrep_proto != nil {
		return
	}
	file_utils_errorcode_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_utils_errorcode_errrep_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrRep); i {
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
			RawDescriptor: file_utils_errorcode_errrep_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_utils_errorcode_errrep_proto_goTypes,
		DependencyIndexes: file_utils_errorcode_errrep_proto_depIdxs,
		MessageInfos:      file_utils_errorcode_errrep_proto_msgTypes,
	}.Build()
	File_utils_errorcode_errrep_proto = out.File
	file_utils_errorcode_errrep_proto_rawDesc = nil
	file_utils_errorcode_errrep_proto_goTypes = nil
	file_utils_errorcode_errrep_proto_depIdxs = nil
}