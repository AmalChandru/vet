// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: violations.proto

package violations

import (
	checks "github.com/safedep/vet/gen/checks"
	filtersuite "github.com/safedep/vet/gen/filtersuite"
	models "github.com/safedep/vet/gen/models"
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

type Violation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckType checks.CheckType    `protobuf:"varint,1,opt,name=check_type,json=checkType,proto3,enum=CheckType" json:"check_type,omitempty"`
	Package   *models.Package     `protobuf:"bytes,2,opt,name=package,proto3" json:"package,omitempty"`
	Filter    *filtersuite.Filter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	Extra     map[string]string   `protobuf:"bytes,4,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Violation) Reset() {
	*x = Violation{}
	mi := &file_violations_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Violation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Violation) ProtoMessage() {}

func (x *Violation) ProtoReflect() protoreflect.Message {
	mi := &file_violations_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Violation.ProtoReflect.Descriptor instead.
func (*Violation) Descriptor() ([]byte, []int) {
	return file_violations_proto_rawDescGZIP(), []int{0}
}

func (x *Violation) GetCheckType() checks.CheckType {
	if x != nil {
		return x.CheckType
	}
	return checks.CheckType(0)
}

func (x *Violation) GetPackage() *models.Package {
	if x != nil {
		return x.Package
	}
	return nil
}

func (x *Violation) GetFilter() *filtersuite.Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *Violation) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_violations_proto protoreflect.FileDescriptor

var file_violations_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x09, 0x56, 0x69, 0x6f, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x22, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x27, 0x5a, 0x25,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x66, 0x65, 0x64,
	0x65, 0x70, 0x2f, 0x76, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x69, 0x6f, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_violations_proto_rawDescOnce sync.Once
	file_violations_proto_rawDescData = file_violations_proto_rawDesc
)

func file_violations_proto_rawDescGZIP() []byte {
	file_violations_proto_rawDescOnce.Do(func() {
		file_violations_proto_rawDescData = protoimpl.X.CompressGZIP(file_violations_proto_rawDescData)
	})
	return file_violations_proto_rawDescData
}

var file_violations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_violations_proto_goTypes = []any{
	(*Violation)(nil),          // 0: Violation
	nil,                        // 1: Violation.ExtraEntry
	(checks.CheckType)(0),      // 2: CheckType
	(*models.Package)(nil),     // 3: Package
	(*filtersuite.Filter)(nil), // 4: Filter
}
var file_violations_proto_depIdxs = []int32{
	2, // 0: Violation.check_type:type_name -> CheckType
	3, // 1: Violation.package:type_name -> Package
	4, // 2: Violation.filter:type_name -> Filter
	1, // 3: Violation.extra:type_name -> Violation.ExtraEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_violations_proto_init() }
func file_violations_proto_init() {
	if File_violations_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_violations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_violations_proto_goTypes,
		DependencyIndexes: file_violations_proto_depIdxs,
		MessageInfos:      file_violations_proto_msgTypes,
	}.Build()
	File_violations_proto = out.File
	file_violations_proto_rawDesc = nil
	file_violations_proto_goTypes = nil
	file_violations_proto_depIdxs = nil
}
