// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Time.proto

package Transport

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

type Time struct {
	Time                 string   `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Time) Reset()         { *m = Time{} }
func (m *Time) String() string { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()    {}
func (*Time) Descriptor() ([]byte, []int) {
	return fileDescriptor_044a53e760627d5b, []int{0}
}

func (m *Time) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Time.Unmarshal(m, b)
}
func (m *Time) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Time.Marshal(b, m, deterministic)
}
func (m *Time) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Time.Merge(m, src)
}
func (m *Time) XXX_Size() int {
	return xxx_messageInfo_Time.Size(m)
}
func (m *Time) XXX_DiscardUnknown() {
	xxx_messageInfo_Time.DiscardUnknown(m)
}

var xxx_messageInfo_Time proto.InternalMessageInfo

func (m *Time) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func init() {
	proto.RegisterType((*Time)(nil), "Transport.Time")
}

func init() { proto.RegisterFile("Time.proto", fileDescriptor_044a53e760627d5b) }

var fileDescriptor_044a53e760627d5b = []byte{
	// 75 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0a, 0xc9, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0c, 0x29, 0x4a, 0xcc, 0x2b, 0x2e, 0xc8, 0x2f,
	0x2a, 0x51, 0x92, 0xe2, 0x62, 0x01, 0x49, 0x08, 0x09, 0x71, 0xb1, 0x94, 0x64, 0xe6, 0xa6, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x49, 0x6c, 0x60, 0xd5, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x30, 0x7b, 0x71, 0xdb, 0x3b, 0x00, 0x00, 0x00,
}