// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frame.proto

package web

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Frame is an extensible message for raw wire communication.
// A protoweb connection is simply a stream of Frames between
// two end points. The end of a connection is denoted when the
// underlying stream or socket is closed.
type Frame struct {
	Request              *Request  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response             *Response `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Frame) Reset()         { *m = Frame{} }
func (m *Frame) String() string { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()    {}
func (*Frame) Descriptor() ([]byte, []int) {
	return fileDescriptor_5379e2b825e15002, []int{0}
}

func (m *Frame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Frame.Unmarshal(m, b)
}
func (m *Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Frame.Marshal(b, m, deterministic)
}
func (m *Frame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Frame.Merge(m, src)
}
func (m *Frame) XXX_Size() int {
	return xxx_messageInfo_Frame.Size(m)
}
func (m *Frame) XXX_DiscardUnknown() {
	xxx_messageInfo_Frame.DiscardUnknown(m)
}

var xxx_messageInfo_Frame proto.InternalMessageInfo

func (m *Frame) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Frame) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*Frame)(nil), "web.Frame")
}

func init() { proto.RegisterFile("frame.proto", fileDescriptor_5379e2b825e15002) }

var fileDescriptor_5379e2b825e15002 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2b, 0x4a, 0xcc,
	0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x4f, 0x4d, 0x92, 0xe2, 0x2d, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x81, 0x88, 0x49, 0xf1, 0x15, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15,
	0x43, 0xd5, 0x28, 0x45, 0x71, 0xb1, 0xba, 0x81, 0xb4, 0x08, 0xa9, 0x71, 0xb1, 0x43, 0x55, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xf1, 0xe8, 0x95, 0xa7, 0x26, 0xe9, 0x05, 0x41, 0xc4, 0x82,
	0x60, 0x92, 0x42, 0x9a, 0x5c, 0x1c, 0x30, 0x23, 0x24, 0x98, 0xc0, 0x0a, 0x79, 0xa1, 0x0a, 0x21,
	0x82, 0x41, 0x70, 0xe9, 0x24, 0x36, 0xb0, 0x15, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee,
	0x4f, 0x85, 0x2b, 0x95, 0x00, 0x00, 0x00,
}
