// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: content/action.enum.proto

package content

import (
	_ "github.com/liov/tiga/protobuf/utils/proto/gogo"
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

// 操作类型
type ActionType int32

const (
	ActionPlaceholder ActionType = 0
	ActionBrowse      ActionType = 1
	ActionLike        ActionType = 2
	ActionUnlike      ActionType = 3
	ActionComment     ActionType = 4
	ActionCollect     ActionType = 5
	ActionReport      ActionType = 6
	ActionGiveAction  ActionType = 7
	ActionApprove     ActionType = 8
	ActionDelete      ActionType = 9
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0: "ActionPlaceholder",
		1: "ActionBrowse",
		2: "ActionLike",
		3: "ActionUnlike",
		4: "ActionComment",
		5: "ActionCollect",
		6: "ActionReport",
		7: "ActionGiveAction",
		8: "ActionApprove",
		9: "ActionDelete",
	}
	ActionType_value = map[string]int32{
		"ActionPlaceholder": 0,
		"ActionBrowse":      1,
		"ActionLike":        2,
		"ActionUnlike":      3,
		"ActionComment":     4,
		"ActionCollect":     5,
		"ActionReport":      6,
		"ActionGiveAction":  7,
		"ActionApprove":     8,
		"ActionDelete":      9,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) OrigString() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_content_action_enum_proto_enumTypes[0].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_content_action_enum_proto_enumTypes[0]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_content_action_enum_proto_rawDescGZIP(), []int{0}
}

// 删除原因
type DelReason int32

const (
	DelReasonPlaceholder                   DelReason = 0
	DelReasonViolationOfLawsAndRegulations DelReason = 1
	DelReasonEroticViolence                DelReason = 3
	DelReasonOther                         DelReason = 255
)

// Enum value maps for DelReason.
var (
	DelReason_name = map[int32]string{
		0:   "DelReasonPlaceholder",
		1:   "DelReasonViolationOfLawsAndRegulations",
		3:   "DelReasonEroticViolence",
		255: "DelReasonOther",
	}
	DelReason_value = map[string]int32{
		"DelReasonPlaceholder":                   0,
		"DelReasonViolationOfLawsAndRegulations": 1,
		"DelReasonEroticViolence":                3,
		"DelReasonOther":                         255,
	}
)

func (x DelReason) Enum() *DelReason {
	p := new(DelReason)
	*p = x
	return p
}

func (x DelReason) OrigString() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DelReason) Descriptor() protoreflect.EnumDescriptor {
	return file_content_action_enum_proto_enumTypes[1].Descriptor()
}

func (DelReason) Type() protoreflect.EnumType {
	return &file_content_action_enum_proto_enumTypes[1]
}

func (x DelReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DelReason.Descriptor instead.
func (DelReason) EnumDescriptor() ([]byte, []int) {
	return file_content_action_enum_proto_rawDescGZIP(), []int{1}
}

// 评论类型
type CommentType int32

const (
	CommentPlaceholder CommentType = 0
	CommentMoment      CommentType = 1
	CommentDiary       CommentType = 2
	CommentDiaryBook   CommentType = 3
	CommentArticle     CommentType = 4
)

// Enum value maps for CommentType.
var (
	CommentType_name = map[int32]string{
		0: "CommentPlaceholder",
		1: "CommentMoment",
		2: "CommentDiary",
		3: "CommentDiaryBook",
		4: "CommentArticle",
	}
	CommentType_value = map[string]int32{
		"CommentPlaceholder": 0,
		"CommentMoment":      1,
		"CommentDiary":       2,
		"CommentDiaryBook":   3,
		"CommentArticle":     4,
	}
)

func (x CommentType) Enum() *CommentType {
	p := new(CommentType)
	*p = x
	return p
}

func (x CommentType) OrigString() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommentType) Descriptor() protoreflect.EnumDescriptor {
	return file_content_action_enum_proto_enumTypes[2].Descriptor()
}

func (CommentType) Type() protoreflect.EnumType {
	return &file_content_action_enum_proto_enumTypes[2]
}

func (x CommentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommentType.Descriptor instead.
func (CommentType) EnumDescriptor() ([]byte, []int) {
	return file_content_action_enum_proto_rawDescGZIP(), []int{2}
}

var File_content_action_enum_proto protoreflect.FileDescriptor

var file_content_action_enum_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0xd8, 0x02, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x21, 0x0a, 0x11, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x10, 0x00, 0x1a, 0x0a, 0x92, 0x9d, 0x20, 0x06, 0xe5, 0x8d, 0xa0,
	0xe4, 0xbd, 0x8d, 0x12, 0x1c, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x72, 0x6f,
	0x77, 0x73, 0x65, 0x10, 0x01, 0x1a, 0x0a, 0x92, 0x9d, 0x20, 0x06, 0xe6, 0xb5, 0x8f, 0xe8, 0xa7,
	0x88, 0x12, 0x1a, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6b, 0x65, 0x10,
	0x02, 0x1a, 0x0a, 0x92, 0x9d, 0x20, 0x06, 0xe7, 0x82, 0xb9, 0xe8, 0xb5, 0x9e, 0x12, 0x1f, 0x0a,
	0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x10, 0x03, 0x1a,
	0x0d, 0x92, 0x9d, 0x20, 0x09, 0xe4, 0xb8, 0x8d, 0xe5, 0x96, 0x9c, 0xe6, 0xac, 0xa2, 0x12, 0x1d,
	0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x10,
	0x04, 0x1a, 0x0a, 0x92, 0x9d, 0x20, 0x06, 0xe8, 0xaf, 0x84, 0xe8, 0xae, 0xba, 0x12, 0x1d, 0x0a,
	0x0d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x10, 0x05,
	0x1a, 0x0a, 0x92, 0x9d, 0x20, 0x06, 0xe6, 0x94, 0xb6, 0xe8, 0x97, 0x8f, 0x12, 0x1c, 0x0a, 0x0c,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x10, 0x06, 0x1a, 0x0a,
	0x92, 0x9d, 0x20, 0x06, 0xe4, 0xb8, 0xbe, 0xe6, 0x8a, 0xa5, 0x12, 0x20, 0x0a, 0x10, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x47, 0x69, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x07,
	0x1a, 0x0a, 0x92, 0x9d, 0x20, 0x06, 0xe5, 0x9b, 0x9e, 0xe9, 0xa6, 0x88, 0x12, 0x1d, 0x0a, 0x0d,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x10, 0x08, 0x1a,
	0x0a, 0x92, 0x9d, 0x20, 0x06, 0xe8, 0xb5, 0x9e, 0xe5, 0x90, 0x8c, 0x12, 0x1c, 0x0a, 0x0c, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x09, 0x1a, 0x0a, 0x92,
	0x9d, 0x20, 0x06, 0xe5, 0x88, 0xa0, 0xe9, 0x99, 0xa4, 0x1a, 0x11, 0xd2, 0xb5, 0x03, 0x0d, 0xf2,
	0x01, 0x0a, 0x4f, 0x72, 0x69, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2a, 0xde, 0x01, 0x0a,
	0x09, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x10, 0x00, 0x1a, 0x0a, 0x92, 0x9d, 0x20, 0x06, 0xe5, 0x8d, 0xa0, 0xe4, 0xbd, 0x8d,
	0x12, 0x42, 0x0a, 0x26, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x56, 0x69, 0x6f,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x4c, 0x61, 0x77, 0x73, 0x41, 0x6e, 0x64, 0x52,
	0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x10, 0x01, 0x1a, 0x16, 0x92, 0x9d,
	0x20, 0x12, 0xe8, 0xbf, 0x9d, 0xe8, 0xbf, 0x94, 0xe6, 0xb3, 0x95, 0xe5, 0xbe, 0x8b, 0xe6, 0xb3,
	0x95, 0xe8, 0xa7, 0x84, 0x12, 0x2d, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x45, 0x72, 0x6f, 0x74, 0x69, 0x63, 0x56, 0x69, 0x6f, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x10,
	0x03, 0x1a, 0x10, 0x92, 0x9d, 0x20, 0x0c, 0xe8, 0x89, 0xb2, 0xe6, 0x83, 0x85, 0xe6, 0x9a, 0xb4,
	0xe5, 0x8a, 0x9b, 0x12, 0x25, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x4f, 0x74, 0x68, 0x65, 0x72, 0x10, 0xff, 0x01, 0x1a, 0x10, 0x92, 0x9d, 0x20, 0x0c, 0xe5, 0x85,
	0xb6, 0xe4, 0xbb, 0x96, 0xe5, 0x8e, 0x9f, 0xe5, 0x9b, 0xa0, 0x1a, 0x11, 0xd2, 0xb5, 0x03, 0x0d,
	0xf2, 0x01, 0x0a, 0x4f, 0x72, 0x69, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2a, 0xc6, 0x01,
	0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a,
	0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x10, 0x00, 0x1a, 0x0a, 0x92, 0x9d, 0x20, 0x06, 0xe5, 0x8d, 0xa0, 0xe4, 0xbd,
	0x8d, 0x12, 0x1d, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x10, 0x01, 0x1a, 0x0a, 0x92, 0x9d, 0x20, 0x06, 0xe7, 0x9e, 0xac, 0xe9, 0x97, 0xb4,
	0x12, 0x1c, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x61, 0x72, 0x79,
	0x10, 0x02, 0x1a, 0x0a, 0x92, 0x9d, 0x20, 0x06, 0xe6, 0x97, 0xa5, 0xe8, 0xae, 0xb0, 0x12, 0x23,
	0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x61, 0x72, 0x79, 0x42, 0x6f,
	0x6f, 0x6b, 0x10, 0x03, 0x1a, 0x0d, 0x92, 0x9d, 0x20, 0x09, 0xe6, 0x97, 0xa5, 0xe8, 0xae, 0xb0,
	0xe6, 0x9c, 0xac, 0x12, 0x1e, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x10, 0x04, 0x1a, 0x0a, 0x92, 0x9d, 0x20, 0x06, 0xe6, 0x96, 0x87,
	0xe7, 0xab, 0xa0, 0x1a, 0x11, 0xd2, 0xb5, 0x03, 0x0d, 0xf2, 0x01, 0x0a, 0x4f, 0x72, 0x69, 0x67,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x53, 0x0a, 0x1a, 0x78, 0x79, 0x7a, 0x2e, 0x68, 0x6f,
	0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x69, 0x6f, 0x76, 0x2f, 0x68, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0xd2,
	0xb5, 0x03, 0x02, 0x50, 0x01, 0xc8, 0x3e, 0x01, 0xd0, 0x3e, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_content_action_enum_proto_rawDescOnce sync.Once
	file_content_action_enum_proto_rawDescData = file_content_action_enum_proto_rawDesc
)

func file_content_action_enum_proto_rawDescGZIP() []byte {
	file_content_action_enum_proto_rawDescOnce.Do(func() {
		file_content_action_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_content_action_enum_proto_rawDescData)
	})
	return file_content_action_enum_proto_rawDescData
}

var file_content_action_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_content_action_enum_proto_goTypes = []interface{}{
	(ActionType)(0),  // 0: content.ActionType
	(DelReason)(0),   // 1: content.DelReason
	(CommentType)(0), // 2: content.CommentType
}
var file_content_action_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_content_action_enum_proto_init() }
func file_content_action_enum_proto_init() {
	if File_content_action_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_content_action_enum_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_content_action_enum_proto_goTypes,
		DependencyIndexes: file_content_action_enum_proto_depIdxs,
		EnumInfos:         file_content_action_enum_proto_enumTypes,
	}.Build()
	File_content_action_enum_proto = out.File
	file_content_action_enum_proto_rawDesc = nil
	file_content_action_enum_proto_goTypes = nil
	file_content_action_enum_proto_depIdxs = nil
}
