// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/proto_product/product.proto

package proto_product

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

type ProductData struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price                int32    `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Qty                  int32    `protobuf:"varint,5,opt,name=qty,proto3" json:"qty,omitempty"`
	Desc                 string   `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductData) Reset()         { *m = ProductData{} }
func (m *ProductData) String() string { return proto.CompactTextString(m) }
func (*ProductData) ProtoMessage()    {}
func (*ProductData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0794dfa2d91fc27, []int{0}
}

func (m *ProductData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductData.Unmarshal(m, b)
}
func (m *ProductData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductData.Marshal(b, m, deterministic)
}
func (m *ProductData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductData.Merge(m, src)
}
func (m *ProductData) XXX_Size() int {
	return xxx_messageInfo_ProductData.Size(m)
}
func (m *ProductData) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductData.DiscardUnknown(m)
}

var xxx_messageInfo_ProductData proto.InternalMessageInfo

func (m *ProductData) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductData) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ProductData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductData) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ProductData) GetQty() int32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *ProductData) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type ProductFilter struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductFilter) Reset()         { *m = ProductFilter{} }
func (m *ProductFilter) String() string { return proto.CompactTextString(m) }
func (*ProductFilter) ProtoMessage()    {}
func (*ProductFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0794dfa2d91fc27, []int{1}
}

func (m *ProductFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductFilter.Unmarshal(m, b)
}
func (m *ProductFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductFilter.Marshal(b, m, deterministic)
}
func (m *ProductFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductFilter.Merge(m, src)
}
func (m *ProductFilter) XXX_Size() int {
	return xxx_messageInfo_ProductFilter.Size(m)
}
func (m *ProductFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ProductFilter proto.InternalMessageInfo

func (m *ProductFilter) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*ProductData)(nil), "ProductData")
	proto.RegisterType((*ProductFilter)(nil), "ProductFilter")
}

func init() {
	proto.RegisterFile("grpc/proto_product/product.proto", fileDescriptor_a0794dfa2d91fc27)
}

var fileDescriptor_a0794dfa2d91fc27 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0x2f, 0x2a, 0x48,
	0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x8f, 0x2f, 0x28, 0xca, 0x4f, 0x29, 0x4d, 0x2e, 0xd1, 0x87,
	0xd2, 0x7a, 0x60, 0x51, 0xa5, 0x36, 0x46, 0x2e, 0xee, 0x00, 0x88, 0x88, 0x4b, 0x62, 0x49, 0xa2,
	0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x53, 0x66, 0x8a,
	0x90, 0x38, 0x17, 0x7b, 0x69, 0x71, 0x6a, 0x51, 0x7c, 0x66, 0x8a, 0x04, 0x13, 0x58, 0x90, 0x0d,
	0xc4, 0xf5, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x56, 0x60, 0xd4,
	0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x44, 0xb8, 0x58, 0x0b, 0x8a, 0x32, 0x93, 0x53, 0x25, 0x58, 0xc0,
	0x4a, 0x21, 0x1c, 0x21, 0x01, 0x2e, 0xe6, 0xc2, 0x92, 0x4a, 0x09, 0x56, 0xb0, 0x18, 0x88, 0x09,
	0xd2, 0x9b, 0x92, 0x5a, 0x9c, 0x2c, 0xc1, 0x06, 0xd1, 0x0b, 0x62, 0x2b, 0xc9, 0x73, 0xf1, 0x42,
	0xdd, 0xe1, 0x96, 0x99, 0x53, 0x92, 0x5a, 0x84, 0xee, 0x12, 0x23, 0x27, 0x2e, 0x3e, 0xa8, 0x82,
	0xe0, 0xd4, 0xa2, 0x32, 0x90, 0xc1, 0x06, 0x5c, 0x7c, 0xee, 0xa9, 0x25, 0x50, 0x41, 0xa7, 0x4a,
	0xcf, 0x14, 0x21, 0x3e, 0x3d, 0x14, 0x33, 0xa4, 0x78, 0xf4, 0x90, 0xfc, 0xa6, 0xc4, 0xe0, 0x24,
	0x13, 0x25, 0xa5, 0x5f, 0x9c, 0x9a, 0x93, 0x93, 0x5a, 0xa4, 0x8f, 0x08, 0x19, 0x3d, 0x68, 0x88,
	0x24, 0xb1, 0x81, 0xb9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x34, 0xa2, 0x0b, 0x36,
	0x01, 0x00, 0x00,
}
