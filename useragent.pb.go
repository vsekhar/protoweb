// Code generated by protoc-gen-go. DO NOT EDIT.
// source: useragent.proto

package web

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CommonUserAgents int32

const (
	CommonUserAgents_COMMON_USER_AGENTS_UNSPECIFIED CommonUserAgents = 0
	CommonUserAgents_MOZILLA                        CommonUserAgents = 1
	CommonUserAgents_APPLE_WEBKIT                   CommonUserAgents = 2
	CommonUserAgents_SAFARI                         CommonUserAgents = 3
	CommonUserAgents_CHROME                         CommonUserAgents = 4
	CommonUserAgents_CRIOS                          CommonUserAgents = 5
)

var CommonUserAgents_name = map[int32]string{
	0: "COMMON_USER_AGENTS_UNSPECIFIED",
	1: "MOZILLA",
	2: "APPLE_WEBKIT",
	3: "SAFARI",
	4: "CHROME",
	5: "CRIOS",
}

var CommonUserAgents_value = map[string]int32{
	"COMMON_USER_AGENTS_UNSPECIFIED": 0,
	"MOZILLA":                        1,
	"APPLE_WEBKIT":                   2,
	"SAFARI":                         3,
	"CHROME":                         4,
	"CRIOS":                          5,
}

func (x CommonUserAgents) String() string {
	return proto.EnumName(CommonUserAgents_name, int32(x))
}

func (CommonUserAgents) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9555fdd29a380cd, []int{0}
}

type UserAgentDescriptor struct {
	HttpName             string   `protobuf:"bytes,1,opt,name=http_name,json=httpName,proto3" json:"http_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAgentDescriptor) Reset()         { *m = UserAgentDescriptor{} }
func (m *UserAgentDescriptor) String() string { return proto.CompactTextString(m) }
func (*UserAgentDescriptor) ProtoMessage()    {}
func (*UserAgentDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9555fdd29a380cd, []int{0}
}

func (m *UserAgentDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAgentDescriptor.Unmarshal(m, b)
}
func (m *UserAgentDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAgentDescriptor.Marshal(b, m, deterministic)
}
func (m *UserAgentDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAgentDescriptor.Merge(m, src)
}
func (m *UserAgentDescriptor) XXX_Size() int {
	return xxx_messageInfo_UserAgentDescriptor.Size(m)
}
func (m *UserAgentDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAgentDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_UserAgentDescriptor proto.InternalMessageInfo

func (m *UserAgentDescriptor) GetHttpName() string {
	if m != nil {
		return m.HttpName
	}
	return ""
}

type UserAgent struct {
	Addenda              []string `protobuf:"bytes,3,rep,name=addenda,proto3" json:"addenda,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAgent) Reset()         { *m = UserAgent{} }
func (m *UserAgent) String() string { return proto.CompactTextString(m) }
func (*UserAgent) ProtoMessage()    {}
func (*UserAgent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9555fdd29a380cd, []int{1}
}

func (m *UserAgent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAgent.Unmarshal(m, b)
}
func (m *UserAgent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAgent.Marshal(b, m, deterministic)
}
func (m *UserAgent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAgent.Merge(m, src)
}
func (m *UserAgent) XXX_Size() int {
	return xxx_messageInfo_UserAgent.Size(m)
}
func (m *UserAgent) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAgent.DiscardUnknown(m)
}

var xxx_messageInfo_UserAgent proto.InternalMessageInfo

func (m *UserAgent) GetAddenda() []string {
	if m != nil {
		return m.Addenda
	}
	return nil
}

type UserAgent_UserAgentEntry struct {
	// Types that are valid to be assigned to UserAgent:
	//	*UserAgent_UserAgentEntry_Common
	//	*UserAgent_UserAgentEntry_Other
	UserAgent            isUserAgent_UserAgentEntry_UserAgent `protobuf_oneof:"UserAgent"`
	VersionNumbers       []uint64                             `protobuf:"varint,3,rep,packed,name=version_numbers,json=versionNumbers,proto3" json:"version_numbers,omitempty"`
	VersionString        string                               `protobuf:"bytes,4,opt,name=version_string,json=versionString,proto3" json:"version_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *UserAgent_UserAgentEntry) Reset()         { *m = UserAgent_UserAgentEntry{} }
func (m *UserAgent_UserAgentEntry) String() string { return proto.CompactTextString(m) }
func (*UserAgent_UserAgentEntry) ProtoMessage()    {}
func (*UserAgent_UserAgentEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9555fdd29a380cd, []int{1, 0}
}

func (m *UserAgent_UserAgentEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAgent_UserAgentEntry.Unmarshal(m, b)
}
func (m *UserAgent_UserAgentEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAgent_UserAgentEntry.Marshal(b, m, deterministic)
}
func (m *UserAgent_UserAgentEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAgent_UserAgentEntry.Merge(m, src)
}
func (m *UserAgent_UserAgentEntry) XXX_Size() int {
	return xxx_messageInfo_UserAgent_UserAgentEntry.Size(m)
}
func (m *UserAgent_UserAgentEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAgent_UserAgentEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UserAgent_UserAgentEntry proto.InternalMessageInfo

type isUserAgent_UserAgentEntry_UserAgent interface {
	isUserAgent_UserAgentEntry_UserAgent()
}

type UserAgent_UserAgentEntry_Common struct {
	Common CommonUserAgents `protobuf:"varint,1,opt,name=common,proto3,enum=web.CommonUserAgents,oneof"`
}

type UserAgent_UserAgentEntry_Other struct {
	Other string `protobuf:"bytes,2,opt,name=other,proto3,oneof"`
}

func (*UserAgent_UserAgentEntry_Common) isUserAgent_UserAgentEntry_UserAgent() {}

func (*UserAgent_UserAgentEntry_Other) isUserAgent_UserAgentEntry_UserAgent() {}

func (m *UserAgent_UserAgentEntry) GetUserAgent() isUserAgent_UserAgentEntry_UserAgent {
	if m != nil {
		return m.UserAgent
	}
	return nil
}

func (m *UserAgent_UserAgentEntry) GetCommon() CommonUserAgents {
	if x, ok := m.GetUserAgent().(*UserAgent_UserAgentEntry_Common); ok {
		return x.Common
	}
	return CommonUserAgents_COMMON_USER_AGENTS_UNSPECIFIED
}

func (m *UserAgent_UserAgentEntry) GetOther() string {
	if x, ok := m.GetUserAgent().(*UserAgent_UserAgentEntry_Other); ok {
		return x.Other
	}
	return ""
}

func (m *UserAgent_UserAgentEntry) GetVersionNumbers() []uint64 {
	if m != nil {
		return m.VersionNumbers
	}
	return nil
}

func (m *UserAgent_UserAgentEntry) GetVersionString() string {
	if m != nil {
		return m.VersionString
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UserAgent_UserAgentEntry) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UserAgent_UserAgentEntry_Common)(nil),
		(*UserAgent_UserAgentEntry_Other)(nil),
	}
}

var E_UserAgentDescriptor = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*UserAgentDescriptor)(nil),
	Field:         7981885,
	Name:          "web.user_agent_descriptor",
	Tag:           "bytes,7981885,opt,name=user_agent_descriptor",
	Filename:      "useragent.proto",
}

func init() {
	proto.RegisterEnum("web.CommonUserAgents", CommonUserAgents_name, CommonUserAgents_value)
	proto.RegisterType((*UserAgentDescriptor)(nil), "web.UserAgentDescriptor")
	proto.RegisterType((*UserAgent)(nil), "web.UserAgent")
	proto.RegisterType((*UserAgent_UserAgentEntry)(nil), "web.UserAgent.UserAgentEntry")
	proto.RegisterExtension(E_UserAgentDescriptor)
}

func init() { proto.RegisterFile("useragent.proto", fileDescriptor_b9555fdd29a380cd) }

var fileDescriptor_b9555fdd29a380cd = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x9b, 0xed, 0xbf, 0xcd, 0xe9, 0xb6, 0x1b, 0x66, 0x59, 0x8d, 0x59, 0x58, 0xea, 0xa2,
	0x58, 0x14, 0x52, 0xa8, 0x77, 0xde, 0x65, 0xb3, 0x59, 0x1b, 0xb6, 0x49, 0xca, 0x64, 0xeb, 0x82,
	0x37, 0x21, 0xd9, 0xce, 0xb6, 0x81, 0x64, 0x26, 0x4c, 0x26, 0x2e, 0xfa, 0x20, 0xbe, 0x88, 0x20,
	0x78, 0xe1, 0x43, 0xf8, 0x02, 0xde, 0x0a, 0x3e, 0x85, 0x24, 0x6d, 0xa3, 0xa8, 0x97, 0xe7, 0xf7,
	0x7d, 0xe7, 0xcc, 0x39, 0xcc, 0x07, 0x87, 0x45, 0x4e, 0x78, 0xb8, 0x22, 0x54, 0xe8, 0x19, 0x67,
	0x82, 0xa1, 0xe6, 0x3d, 0x89, 0xb4, 0xe1, 0x8a, 0xb1, 0x55, 0x42, 0xc6, 0x15, 0x8a, 0x8a, 0xbb,
	0xf1, 0x92, 0xe4, 0xb7, 0x3c, 0xce, 0x04, 0xe3, 0x1b, 0xdb, 0xd9, 0x04, 0x8e, 0x16, 0x39, 0xe1,
	0x46, 0xd9, 0x79, 0x51, 0x8b, 0xe8, 0x04, 0xe4, 0xb5, 0x10, 0x59, 0x40, 0xc3, 0x94, 0xa8, 0xd2,
	0x50, 0x1a, 0xc9, 0x78, 0xbf, 0x04, 0x6e, 0x98, 0x92, 0xb3, 0xef, 0x12, 0xc8, 0x75, 0x13, 0x52,
	0xa1, 0x1b, 0x2e, 0x97, 0x84, 0x2e, 0x43, 0xb5, 0x39, 0x6c, 0x8e, 0x64, 0xbc, 0x2b, 0xb5, 0xcf,
	0x12, 0x0c, 0x6a, 0x9f, 0x45, 0x05, 0x7f, 0x8f, 0xc6, 0xd0, 0xb9, 0x65, 0x69, 0xca, 0x68, 0x35,
	0x74, 0x30, 0x39, 0xd6, 0xef, 0x49, 0xa4, 0x9b, 0x15, 0xaa, 0xad, 0xf9, 0xb4, 0x81, 0xb7, 0x36,
	0xf4, 0x00, 0xda, 0x4c, 0xac, 0x09, 0x57, 0xf7, 0xca, 0x25, 0xa6, 0x0d, 0xbc, 0x29, 0xd1, 0x33,
	0x38, 0x7c, 0x47, 0x78, 0x1e, 0x33, 0x1a, 0xd0, 0x22, 0x8d, 0x08, 0xcf, 0xab, 0xd7, 0x5b, 0x78,
	0xb0, 0xc5, 0xee, 0x86, 0xa2, 0xa7, 0xb0, 0x23, 0x41, 0x2e, 0x78, 0x4c, 0x57, 0x6a, 0xab, 0x3a,
	0xa7, 0xbf, 0xa5, 0x7e, 0x05, 0xcf, 0x7b, 0x7f, 0x9c, 0xf4, 0xfc, 0x9b, 0x04, 0xca, 0xdf, 0x3b,
	0xa1, 0x17, 0x70, 0x6a, 0x7a, 0x8e, 0xe3, 0xb9, 0xc1, 0xc2, 0xb7, 0x70, 0x60, 0xbc, 0xb6, 0xdc,
	0x6b, 0x3f, 0x58, 0xb8, 0xfe, 0xdc, 0x32, 0xed, 0x4b, 0xdb, 0xba, 0x50, 0x1a, 0x5a, 0xf7, 0xe7,
	0xa7, 0x2f, 0xa7, 0x7b, 0xd0, 0x40, 0x27, 0xd0, 0x75, 0xbc, 0xb7, 0xf6, 0x6c, 0x66, 0x28, 0x92,
	0x36, 0x28, 0xa9, 0x0c, 0x5d, 0x87, 0x7d, 0x88, 0x93, 0x24, 0x44, 0x4f, 0xe0, 0xc0, 0x98, 0xcf,
	0x67, 0x56, 0x70, 0x63, 0x9d, 0x5f, 0xd9, 0xd7, 0xca, 0x9e, 0x86, 0x4a, 0x47, 0x1f, 0x7a, 0x46,
	0x96, 0x25, 0xe4, 0x86, 0x44, 0x57, 0xb1, 0x40, 0x8f, 0xa0, 0xe3, 0x1b, 0x97, 0x06, 0xb6, 0x95,
	0xa6, 0xd6, 0x2f, 0xf5, 0x7d, 0xe8, 0xf8, 0xe1, 0x5d, 0xc8, 0xe3, 0x52, 0x32, 0xa7, 0xd8, 0x73,
	0x2c, 0xa5, 0x55, 0x4b, 0xe6, 0x9a, 0xb3, 0x94, 0xa0, 0x87, 0xd0, 0x36, 0xb1, 0xed, 0xf9, 0x4a,
	0x5b, 0x3b, 0x28, 0x95, 0x2e, 0xb4, 0x4d, 0x1e, 0x7b, 0xfe, 0x2b, 0x0e, 0xc7, 0x65, 0x44, 0x82,
	0x2a, 0x23, 0xc1, 0xef, 0x1c, 0xa0, 0xc7, 0xfa, 0x26, 0x24, 0xfa, 0x2e, 0x24, 0xba, 0x45, 0x8b,
	0xf4, 0x4d, 0x98, 0x14, 0xc4, 0xcb, 0x44, 0xcc, 0x68, 0xae, 0x7e, 0xfd, 0xf8, 0xa3, 0x39, 0x94,
	0x46, 0xbd, 0x89, 0x5a, 0x7d, 0xd7, 0x7f, 0x02, 0x83, 0x8f, 0x8a, 0x7f, 0x61, 0xd4, 0xa9, 0x46,
	0xbe, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x72, 0x83, 0xe7, 0x3b, 0x9d, 0x02, 0x00, 0x00,
}
