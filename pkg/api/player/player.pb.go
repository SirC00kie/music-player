// Code generated by protoc-gen-go. DO NOT EDIT.
// source: player.proto

package player

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type Song struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author               string   `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Duration             int64    `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Song) Reset()         { *m = Song{} }
func (m *Song) String() string { return proto.CompactTextString(m) }
func (*Song) ProtoMessage()    {}
func (*Song) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d803d1b635d5c6, []int{0}
}

func (m *Song) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Song.Unmarshal(m, b)
}
func (m *Song) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Song.Marshal(b, m, deterministic)
}
func (m *Song) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Song.Merge(m, src)
}
func (m *Song) XXX_Size() int {
	return xxx_messageInfo_Song.Size(m)
}
func (m *Song) XXX_DiscardUnknown() {
	xxx_messageInfo_Song.DiscardUnknown(m)
}

var xxx_messageInfo_Song proto.InternalMessageInfo

func (m *Song) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Song) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Song) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func init() {
	proto.RegisterType((*Song)(nil), "api.Song")
}

func init() { proto.RegisterFile("player.proto", fileDescriptor_41d803d1b635d5c6) }

var fileDescriptor_41d803d1b635d5c6 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8e, 0xb1, 0x4a, 0xc5, 0x30,
	0x14, 0x86, 0x8d, 0xbd, 0x96, 0xdb, 0x83, 0x2e, 0x41, 0x4a, 0xa9, 0x4b, 0xe9, 0xd4, 0x29, 0x11,
	0x7d, 0x03, 0xc1, 0x55, 0x42, 0xbb, 0xb9, 0xa5, 0x7a, 0x8c, 0x81, 0xd8, 0x84, 0x98, 0x14, 0xfb,
	0x9e, 0x3e, 0x90, 0x34, 0x55, 0xe7, 0xde, 0xf1, 0x3f, 0x7c, 0x1f, 0xe7, 0x83, 0x4b, 0x67, 0xe4,
	0x82, 0x9e, 0x39, 0x6f, 0x83, 0xa5, 0x99, 0x74, 0xba, 0xbe, 0x51, 0xd6, 0x2a, 0x83, 0x3c, 0x9d,
	0xc6, 0xf8, 0xc6, 0xf1, 0xc3, 0x85, 0x65, 0x23, 0x5a, 0x01, 0x87, 0xc1, 0x4e, 0x8a, 0x5e, 0xc3,
	0x45, 0xd0, 0xc1, 0x60, 0x45, 0x1a, 0xd2, 0x15, 0xfd, 0x36, 0x68, 0x09, 0xb9, 0x8c, 0xe1, 0xdd,
	0xfa, 0xea, 0x3c, 0x9d, 0x7f, 0x17, 0xad, 0xe1, 0xf8, 0x1a, 0xbd, 0x0c, 0xda, 0x4e, 0x55, 0xd6,
	0x90, 0x2e, 0xeb, 0xff, 0xf7, 0xdd, 0x37, 0x81, 0x2b, 0x91, 0x22, 0x06, 0xf4, 0xb3, 0x7e, 0x41,
	0xca, 0xe1, 0xf8, 0x84, 0x5f, 0x21, 0xfd, 0x29, 0xd9, 0x56, 0xc3, 0xfe, 0x6a, 0xd8, 0xe3, 0x5a,
	0x53, 0x17, 0x4c, 0x3a, 0xcd, 0x56, 0xa4, 0x3d, 0x5b, 0x05, 0xe1, 0x71, 0x3e, 0x4d, 0x30, 0x72,
	0xd9, 0x2f, 0xdc, 0x42, 0x21, 0x64, 0xfc, 0xc4, 0xdd, 0xc6, 0x43, 0xfe, 0x7c, 0xe0, 0x0a, 0xa7,
	0x31, 0x4f, 0xd0, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0x76, 0x8c, 0x76, 0x69, 0x01,
	0x00, 0x00,
}
