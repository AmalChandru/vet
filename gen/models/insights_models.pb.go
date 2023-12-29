// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: insights_models.proto

package models

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

type InsightVulnerabilitySeverity_Type int32

const (
	InsightVulnerabilitySeverity_UNKNOWN_TYPE InsightVulnerabilitySeverity_Type = 0
	InsightVulnerabilitySeverity_CVSSV2       InsightVulnerabilitySeverity_Type = 1
	InsightVulnerabilitySeverity_CVSSV3       InsightVulnerabilitySeverity_Type = 2
)

// Enum value maps for InsightVulnerabilitySeverity_Type.
var (
	InsightVulnerabilitySeverity_Type_name = map[int32]string{
		0: "UNKNOWN_TYPE",
		1: "CVSSV2",
		2: "CVSSV3",
	}
	InsightVulnerabilitySeverity_Type_value = map[string]int32{
		"UNKNOWN_TYPE": 0,
		"CVSSV2":       1,
		"CVSSV3":       2,
	}
)

func (x InsightVulnerabilitySeverity_Type) Enum() *InsightVulnerabilitySeverity_Type {
	p := new(InsightVulnerabilitySeverity_Type)
	*p = x
	return p
}

func (x InsightVulnerabilitySeverity_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InsightVulnerabilitySeverity_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_insights_models_proto_enumTypes[0].Descriptor()
}

func (InsightVulnerabilitySeverity_Type) Type() protoreflect.EnumType {
	return &file_insights_models_proto_enumTypes[0]
}

func (x InsightVulnerabilitySeverity_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InsightVulnerabilitySeverity_Type.Descriptor instead.
func (InsightVulnerabilitySeverity_Type) EnumDescriptor() ([]byte, []int) {
	return file_insights_models_proto_rawDescGZIP(), []int{0, 0}
}

type InsightVulnerabilitySeverity_Risk int32

const (
	InsightVulnerabilitySeverity_UNKNOWN_RISK InsightVulnerabilitySeverity_Risk = 0
	InsightVulnerabilitySeverity_LOW          InsightVulnerabilitySeverity_Risk = 1
	InsightVulnerabilitySeverity_MEDIUM       InsightVulnerabilitySeverity_Risk = 2
	InsightVulnerabilitySeverity_HIGH         InsightVulnerabilitySeverity_Risk = 3
	InsightVulnerabilitySeverity_CRITICAL     InsightVulnerabilitySeverity_Risk = 4
)

// Enum value maps for InsightVulnerabilitySeverity_Risk.
var (
	InsightVulnerabilitySeverity_Risk_name = map[int32]string{
		0: "UNKNOWN_RISK",
		1: "LOW",
		2: "MEDIUM",
		3: "HIGH",
		4: "CRITICAL",
	}
	InsightVulnerabilitySeverity_Risk_value = map[string]int32{
		"UNKNOWN_RISK": 0,
		"LOW":          1,
		"MEDIUM":       2,
		"HIGH":         3,
		"CRITICAL":     4,
	}
)

func (x InsightVulnerabilitySeverity_Risk) Enum() *InsightVulnerabilitySeverity_Risk {
	p := new(InsightVulnerabilitySeverity_Risk)
	*p = x
	return p
}

func (x InsightVulnerabilitySeverity_Risk) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InsightVulnerabilitySeverity_Risk) Descriptor() protoreflect.EnumDescriptor {
	return file_insights_models_proto_enumTypes[1].Descriptor()
}

func (InsightVulnerabilitySeverity_Risk) Type() protoreflect.EnumType {
	return &file_insights_models_proto_enumTypes[1]
}

func (x InsightVulnerabilitySeverity_Risk) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InsightVulnerabilitySeverity_Risk.Descriptor instead.
func (InsightVulnerabilitySeverity_Risk) EnumDescriptor() ([]byte, []int) {
	return file_insights_models_proto_rawDescGZIP(), []int{0, 1}
}

type InsightProjectInfo_Type int32

const (
	InsightProjectInfo_UNKNOWN InsightProjectInfo_Type = 0
	InsightProjectInfo_GITHUB  InsightProjectInfo_Type = 1
)

// Enum value maps for InsightProjectInfo_Type.
var (
	InsightProjectInfo_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "GITHUB",
	}
	InsightProjectInfo_Type_value = map[string]int32{
		"UNKNOWN": 0,
		"GITHUB":  1,
	}
)

func (x InsightProjectInfo_Type) Enum() *InsightProjectInfo_Type {
	p := new(InsightProjectInfo_Type)
	*p = x
	return p
}

func (x InsightProjectInfo_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InsightProjectInfo_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_insights_models_proto_enumTypes[2].Descriptor()
}

func (InsightProjectInfo_Type) Type() protoreflect.EnumType {
	return &file_insights_models_proto_enumTypes[2]
}

func (x InsightProjectInfo_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InsightProjectInfo_Type.Descriptor instead.
func (InsightProjectInfo_Type) EnumDescriptor() ([]byte, []int) {
	return file_insights_models_proto_rawDescGZIP(), []int{4, 0}
}

type InsightVulnerabilitySeverity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  InsightVulnerabilitySeverity_Type `protobuf:"varint,1,opt,name=type,proto3,enum=InsightVulnerabilitySeverity_Type" json:"type,omitempty"`
	Score string                            `protobuf:"bytes,2,opt,name=score,proto3" json:"score,omitempty"` // Score based on type (usually the CVSS metric)
	Risk  InsightVulnerabilitySeverity_Risk `protobuf:"varint,3,opt,name=risk,proto3,enum=InsightVulnerabilitySeverity_Risk" json:"risk,omitempty"`
}

func (x *InsightVulnerabilitySeverity) Reset() {
	*x = InsightVulnerabilitySeverity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_insights_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsightVulnerabilitySeverity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsightVulnerabilitySeverity) ProtoMessage() {}

func (x *InsightVulnerabilitySeverity) ProtoReflect() protoreflect.Message {
	mi := &file_insights_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsightVulnerabilitySeverity.ProtoReflect.Descriptor instead.
func (*InsightVulnerabilitySeverity) Descriptor() ([]byte, []int) {
	return file_insights_models_proto_rawDescGZIP(), []int{0}
}

func (x *InsightVulnerabilitySeverity) GetType() InsightVulnerabilitySeverity_Type {
	if x != nil {
		return x.Type
	}
	return InsightVulnerabilitySeverity_UNKNOWN_TYPE
}

func (x *InsightVulnerabilitySeverity) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *InsightVulnerabilitySeverity) GetRisk() InsightVulnerabilitySeverity_Risk {
	if x != nil {
		return x.Risk
	}
	return InsightVulnerabilitySeverity_UNKNOWN_RISK
}

type InsightVulnerability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`   // OSV ID
	Cve        string                          `protobuf:"bytes,2,opt,name=cve,proto3" json:"cve,omitempty"` // CVE ID. DO NOT USE THIS outside vet. Its used for internal legacy reason
	Title      string                          `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Aliases    []string                        `protobuf:"bytes,4,rep,name=aliases,proto3" json:"aliases,omitempty"` // Other IDs for same vuln in different databases
	Severities []*InsightVulnerabilitySeverity `protobuf:"bytes,5,rep,name=severities,proto3" json:"severities,omitempty"`
}

func (x *InsightVulnerability) Reset() {
	*x = InsightVulnerability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_insights_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsightVulnerability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsightVulnerability) ProtoMessage() {}

func (x *InsightVulnerability) ProtoReflect() protoreflect.Message {
	mi := &file_insights_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsightVulnerability.ProtoReflect.Descriptor instead.
func (*InsightVulnerability) Descriptor() ([]byte, []int) {
	return file_insights_models_proto_rawDescGZIP(), []int{1}
}

func (x *InsightVulnerability) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InsightVulnerability) GetCve() string {
	if x != nil {
		return x.Cve
	}
	return ""
}

func (x *InsightVulnerability) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *InsightVulnerability) GetAliases() []string {
	if x != nil {
		return x.Aliases
	}
	return nil
}

func (x *InsightVulnerability) GetSeverities() []*InsightVulnerabilitySeverity {
	if x != nil {
		return x.Severities
	}
	return nil
}

type InsightLicenseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // SPDX license ID
}

func (x *InsightLicenseInfo) Reset() {
	*x = InsightLicenseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_insights_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsightLicenseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsightLicenseInfo) ProtoMessage() {}

func (x *InsightLicenseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_insights_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsightLicenseInfo.ProtoReflect.Descriptor instead.
func (*InsightLicenseInfo) Descriptor() ([]byte, []int) {
	return file_insights_models_proto_rawDescGZIP(), []int{2}
}

func (x *InsightLicenseInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type InsightScorecard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scores map[string]float32 `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	Score  float32            `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *InsightScorecard) Reset() {
	*x = InsightScorecard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_insights_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsightScorecard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsightScorecard) ProtoMessage() {}

func (x *InsightScorecard) ProtoReflect() protoreflect.Message {
	mi := &file_insights_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsightScorecard.ProtoReflect.Descriptor instead.
func (*InsightScorecard) Descriptor() ([]byte, []int) {
	return file_insights_models_proto_rawDescGZIP(), []int{3}
}

func (x *InsightScorecard) GetScores() map[string]float32 {
	if x != nil {
		return x.Scores
	}
	return nil
}

func (x *InsightScorecard) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type InsightProjectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type   InsightProjectInfo_Type `protobuf:"varint,2,opt,name=type,proto3,enum=InsightProjectInfo_Type" json:"type,omitempty"`
	Stars  int32                   `protobuf:"varint,3,opt,name=stars,proto3" json:"stars,omitempty"`
	Forks  int32                   `protobuf:"varint,4,opt,name=forks,proto3" json:"forks,omitempty"`
	Issues int32                   `protobuf:"varint,5,opt,name=issues,proto3" json:"issues,omitempty"`
}

func (x *InsightProjectInfo) Reset() {
	*x = InsightProjectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_insights_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsightProjectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsightProjectInfo) ProtoMessage() {}

func (x *InsightProjectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_insights_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsightProjectInfo.ProtoReflect.Descriptor instead.
func (*InsightProjectInfo) Descriptor() ([]byte, []int) {
	return file_insights_models_proto_rawDescGZIP(), []int{4}
}

func (x *InsightProjectInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InsightProjectInfo) GetType() InsightProjectInfo_Type {
	if x != nil {
		return x.Type
	}
	return InsightProjectInfo_UNKNOWN
}

func (x *InsightProjectInfo) GetStars() int32 {
	if x != nil {
		return x.Stars
	}
	return 0
}

func (x *InsightProjectInfo) GetForks() int32 {
	if x != nil {
		return x.Forks
	}
	return 0
}

func (x *InsightProjectInfo) GetIssues() int32 {
	if x != nil {
		return x.Issues
	}
	return 0
}

var File_insights_models_proto protoreflect.FileDescriptor

var file_insights_models_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x1c, 0x49, 0x6e, 0x73, 0x69,
	0x67, 0x68, 0x74, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74,
	0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x72, 0x69, 0x73, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x56, 0x75,
	0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x2e, 0x52, 0x69, 0x73, 0x6b, 0x52, 0x04, 0x72, 0x69, 0x73, 0x6b, 0x22, 0x30,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x56, 0x53, 0x53,
	0x56, 0x32, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x56, 0x53, 0x53, 0x56, 0x33, 0x10, 0x02,
	0x22, 0x45, 0x0a, 0x04, 0x52, 0x69, 0x73, 0x6b, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x52, 0x49, 0x53, 0x4b, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f,
	0x57, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x49,
	0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x04, 0x22, 0xa7, 0x01, 0x0a, 0x14, 0x49, 0x6e, 0x73, 0x69,
	0x67, 0x68, 0x74, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74,
	0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x22, 0x24, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9a, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x69,
	0x67, 0x68, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x12, 0x35, 0x0a, 0x06,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x49,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x2e,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xbb, 0x01, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x73, 0x22, 0x1f, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x49, 0x54, 0x48, 0x55, 0x42,
	0x10, 0x01, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x61, 0x66, 0x65, 0x64, 0x65, 0x70, 0x2f, 0x76, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_insights_models_proto_rawDescOnce sync.Once
	file_insights_models_proto_rawDescData = file_insights_models_proto_rawDesc
)

func file_insights_models_proto_rawDescGZIP() []byte {
	file_insights_models_proto_rawDescOnce.Do(func() {
		file_insights_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_insights_models_proto_rawDescData)
	})
	return file_insights_models_proto_rawDescData
}

var file_insights_models_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_insights_models_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_insights_models_proto_goTypes = []interface{}{
	(InsightVulnerabilitySeverity_Type)(0), // 0: InsightVulnerabilitySeverity.Type
	(InsightVulnerabilitySeverity_Risk)(0), // 1: InsightVulnerabilitySeverity.Risk
	(InsightProjectInfo_Type)(0),           // 2: InsightProjectInfo.Type
	(*InsightVulnerabilitySeverity)(nil),   // 3: InsightVulnerabilitySeverity
	(*InsightVulnerability)(nil),           // 4: InsightVulnerability
	(*InsightLicenseInfo)(nil),             // 5: InsightLicenseInfo
	(*InsightScorecard)(nil),               // 6: InsightScorecard
	(*InsightProjectInfo)(nil),             // 7: InsightProjectInfo
	nil,                                    // 8: InsightScorecard.ScoresEntry
}
var file_insights_models_proto_depIdxs = []int32{
	0, // 0: InsightVulnerabilitySeverity.type:type_name -> InsightVulnerabilitySeverity.Type
	1, // 1: InsightVulnerabilitySeverity.risk:type_name -> InsightVulnerabilitySeverity.Risk
	3, // 2: InsightVulnerability.severities:type_name -> InsightVulnerabilitySeverity
	8, // 3: InsightScorecard.scores:type_name -> InsightScorecard.ScoresEntry
	2, // 4: InsightProjectInfo.type:type_name -> InsightProjectInfo.Type
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_insights_models_proto_init() }
func file_insights_models_proto_init() {
	if File_insights_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_insights_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsightVulnerabilitySeverity); i {
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
		file_insights_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsightVulnerability); i {
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
		file_insights_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsightLicenseInfo); i {
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
		file_insights_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsightScorecard); i {
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
		file_insights_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsightProjectInfo); i {
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
			RawDescriptor: file_insights_models_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_insights_models_proto_goTypes,
		DependencyIndexes: file_insights_models_proto_depIdxs,
		EnumInfos:         file_insights_models_proto_enumTypes,
		MessageInfos:      file_insights_models_proto_msgTypes,
	}.Build()
	File_insights_models_proto = out.File
	file_insights_models_proto_rawDesc = nil
	file_insights_models_proto_goTypes = nil
	file_insights_models_proto_depIdxs = nil
}