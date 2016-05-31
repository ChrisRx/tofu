// Code generated by protoc-gen-go.
// source: proto/common.proto
// DO NOT EDIT!

package tofu

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BytesValue struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *BytesValue) Reset()                    { *m = BytesValue{} }
func (m *BytesValue) String() string            { return proto.CompactTextString(m) }
func (*BytesValue) ProtoMessage()               {}
func (*BytesValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type EmptyValue struct {
}

func (m *EmptyValue) Reset()                    { *m = EmptyValue{} }
func (m *EmptyValue) String() string            { return proto.CompactTextString(m) }
func (*EmptyValue) ProtoMessage()               {}
func (*EmptyValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*BytesValue)(nil), "tofu.BytesValue")
	proto.RegisterType((*EmptyValue)(nil), "tofu.EmptyValue")
}

var fileDescriptor1 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x03, 0x73, 0x84, 0x58, 0x4a, 0xf2, 0xd3,
	0x4a, 0xa5, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x80, 0x52, 0xfa, 0xe9, 0xf9, 0xe9,
	0xf9, 0xfa, 0x60, 0xc9, 0xa4, 0xd2, 0x34, 0x30, 0x0f, 0xa2, 0x0d, 0xc4, 0x82, 0x68, 0x52, 0x52,
	0xe0, 0xe2, 0x72, 0xaa, 0x2c, 0x49, 0x2d, 0x0e, 0x4b, 0xcc, 0x29, 0x4d, 0x15, 0x12, 0xe2, 0x62,
	0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x95, 0x78, 0xb8,
	0xb8, 0x5c, 0x73, 0x0b, 0x4a, 0x2a, 0xc1, 0x2a, 0x92, 0xd8, 0xc0, 0xda, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x3a, 0x84, 0xd9, 0x17, 0x81, 0x00, 0x00, 0x00,
}