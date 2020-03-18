// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/ai/vision/v1/classification.proto

package vision

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

type ClassAnnotation struct {
	// Properties extracted by a specified model.
	//
	// For example, if you ask to evaluate the image quality,
	// the service could return such properties as `good` and `bad`.
	Properties           []*Property `protobuf:"bytes,1,rep,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClassAnnotation) Reset()         { *m = ClassAnnotation{} }
func (m *ClassAnnotation) String() string { return proto.CompactTextString(m) }
func (*ClassAnnotation) ProtoMessage()    {}
func (*ClassAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b6a3b28cb9191ec, []int{0}
}

func (m *ClassAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassAnnotation.Unmarshal(m, b)
}
func (m *ClassAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassAnnotation.Marshal(b, m, deterministic)
}
func (m *ClassAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassAnnotation.Merge(m, src)
}
func (m *ClassAnnotation) XXX_Size() int {
	return xxx_messageInfo_ClassAnnotation.Size(m)
}
func (m *ClassAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_ClassAnnotation proto.InternalMessageInfo

func (m *ClassAnnotation) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

type Property struct {
	// Property name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Probability of the property, from 0 to 1.
	Probability          float64  `protobuf:"fixed64,2,opt,name=probability,proto3" json:"probability,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Property) Reset()         { *m = Property{} }
func (m *Property) String() string { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()    {}
func (*Property) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b6a3b28cb9191ec, []int{1}
}

func (m *Property) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Property.Unmarshal(m, b)
}
func (m *Property) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Property.Marshal(b, m, deterministic)
}
func (m *Property) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Property.Merge(m, src)
}
func (m *Property) XXX_Size() int {
	return xxx_messageInfo_Property.Size(m)
}
func (m *Property) XXX_DiscardUnknown() {
	xxx_messageInfo_Property.DiscardUnknown(m)
}

var xxx_messageInfo_Property proto.InternalMessageInfo

func (m *Property) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Property) GetProbability() float64 {
	if m != nil {
		return m.Probability
	}
	return 0
}

func init() {
	proto.RegisterType((*ClassAnnotation)(nil), "yandex.cloud.ai.vision.v1.ClassAnnotation")
	proto.RegisterType((*Property)(nil), "yandex.cloud.ai.vision.v1.Property")
}

func init() {
	proto.RegisterFile("yandex/cloud/ai/vision/v1/classification.proto", fileDescriptor_4b6a3b28cb9191ec)
}

var fileDescriptor_4b6a3b28cb9191ec = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0x89, 0x8a, 0xe8, 0x5c, 0x21, 0xa4, 0x5a, 0x0b, 0x61, 0x39, 0x9b, 0x6d, 0x6e, 0xc2,
	0x69, 0x69, 0xa3, 0x9e, 0x3f, 0x40, 0xb6, 0xb0, 0xb0, 0x4b, 0x72, 0x71, 0x1d, 0xd8, 0xcb, 0x84,
	0x24, 0xb7, 0xb8, 0xff, 0x5e, 0x4c, 0x10, 0xee, 0x8a, 0xeb, 0x1e, 0x8f, 0xef, 0xcd, 0x63, 0x1e,
	0xe0, 0xac, 0xfd, 0xd6, 0xfd, 0x28, 0x3b, 0xf2, 0x7e, 0xab, 0x34, 0xa9, 0x89, 0x12, 0xb1, 0x57,
	0xd3, 0x5a, 0xd9, 0x51, 0xa7, 0x44, 0x5f, 0x64, 0x75, 0x26, 0xf6, 0x18, 0x22, 0x67, 0x96, 0xb7,
	0x95, 0xc7, 0xc2, 0xa3, 0x26, 0xac, 0x3c, 0x4e, 0xeb, 0xe5, 0x07, 0xdc, 0x6c, 0xfe, 0x22, 0x2f,
	0xde, 0x73, 0x2e, 0x19, 0xb9, 0x01, 0x08, 0x91, 0x83, 0x8b, 0x99, 0x5c, 0x6a, 0x44, 0x7b, 0xde,
	0x2d, 0x1e, 0xee, 0xf1, 0xe4, 0x09, 0x7c, 0xaf, 0xf0, 0xdc, 0x1f, 0xc4, 0x96, 0xcf, 0x70, 0xf5,
	0xef, 0x4b, 0x09, 0x17, 0x5e, 0xef, 0x5c, 0x23, 0x5a, 0xd1, 0x5d, 0xf7, 0x45, 0xcb, 0x16, 0x16,
	0x21, 0xb2, 0xd1, 0x86, 0x46, 0xca, 0x73, 0x73, 0xd6, 0x8a, 0x4e, 0xf4, 0x87, 0xd6, 0xab, 0x83,
	0xbb, 0xe3, 0xce, 0x40, 0x47, 0xbd, 0x9f, 0x6f, 0x03, 0xe5, 0xef, 0xbd, 0x41, 0xcb, 0x3b, 0x55,
	0xc9, 0x55, 0x1d, 0x64, 0xe0, 0xd5, 0xe0, 0x7c, 0x79, 0x5d, 0x9d, 0x5c, 0xea, 0xa9, 0x2a, 0x73,
	0x59, 0xb8, 0xc7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0xdf, 0xa0, 0x72, 0x54, 0x01, 0x00,
	0x00,
}