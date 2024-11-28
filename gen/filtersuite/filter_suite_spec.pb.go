// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: filter_suite_spec.proto

package filtersuite

import (
	checks "github.com/safedep/vet/gen/checks"
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

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value       string           `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	CheckType   checks.CheckType `protobuf:"varint,3,opt,name=check_type,json=checkType,proto3,enum=CheckType" json:"check_type,omitempty"`
	Summary     string           `protobuf:"bytes,4,opt,name=summary,proto3" json:"summary,omitempty"`
	Description string           `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	References  []string         `protobuf:"bytes,6,rep,name=references,proto3" json:"references,omitempty"`
	Tags        []string         `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	mi := &file_filter_suite_spec_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_filter_suite_spec_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_filter_suite_spec_proto_rawDescGZIP(), []int{0}
}

func (x *Filter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Filter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Filter) GetCheckType() checks.CheckType {
	if x != nil {
		return x.CheckType
	}
	return checks.CheckType(0)
}

func (x *Filter) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Filter) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Filter) GetReferences() []string {
	if x != nil {
		return x.References
	}
	return nil
}

func (x *Filter) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type FilterSuite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Filters     []*Filter `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	Tags        []string  `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *FilterSuite) Reset() {
	*x = FilterSuite{}
	mi := &file_filter_suite_spec_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilterSuite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterSuite) ProtoMessage() {}

func (x *FilterSuite) ProtoReflect() protoreflect.Message {
	mi := &file_filter_suite_spec_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterSuite.ProtoReflect.Descriptor instead.
func (*FilterSuite) Descriptor() ([]byte, []int) {
	return file_filter_suite_spec_proto_rawDescGZIP(), []int{1}
}

func (x *FilterSuite) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FilterSuite) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FilterSuite) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *FilterSuite) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_filter_suite_spec_proto protoreflect.FileDescriptor

var file_filter_suite_spec_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x0a,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0a, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x7a, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x61, 0x66, 0x65, 0x64, 0x65, 0x70, 0x2f, 0x76, 0x65, 0x74, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x75, 0x69, 0x74, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filter_suite_spec_proto_rawDescOnce sync.Once
	file_filter_suite_spec_proto_rawDescData = file_filter_suite_spec_proto_rawDesc
)

func file_filter_suite_spec_proto_rawDescGZIP() []byte {
	file_filter_suite_spec_proto_rawDescOnce.Do(func() {
		file_filter_suite_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_filter_suite_spec_proto_rawDescData)
	})
	return file_filter_suite_spec_proto_rawDescData
}

var file_filter_suite_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_filter_suite_spec_proto_goTypes = []any{
	(*Filter)(nil),        // 0: Filter
	(*FilterSuite)(nil),   // 1: FilterSuite
	(checks.CheckType)(0), // 2: CheckType
}
var file_filter_suite_spec_proto_depIdxs = []int32{
	2, // 0: Filter.check_type:type_name -> CheckType
	0, // 1: FilterSuite.filters:type_name -> Filter
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_filter_suite_spec_proto_init() }
func file_filter_suite_spec_proto_init() {
	if File_filter_suite_spec_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_filter_suite_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_filter_suite_spec_proto_goTypes,
		DependencyIndexes: file_filter_suite_spec_proto_depIdxs,
		MessageInfos:      file_filter_suite_spec_proto_msgTypes,
	}.Build()
	File_filter_suite_spec_proto = out.File
	file_filter_suite_spec_proto_rawDesc = nil
	file_filter_suite_spec_proto_goTypes = nil
	file_filter_suite_spec_proto_depIdxs = nil
}
