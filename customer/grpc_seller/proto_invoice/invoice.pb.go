// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_seller/proto_invoice/invoice.proto

package proto_invoice

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

type InvoiceData struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DriverId             int32    `protobuf:"varint,3,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Code                 string   `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	ShippingCost         int32    `protobuf:"varint,5,opt,name=shipping_cost,json=shippingCost,proto3" json:"shipping_cost,omitempty"`
	FromLatitude         float64  `protobuf:"fixed64,6,opt,name=from_latitude,json=fromLatitude,proto3" json:"from_latitude,omitempty"`
	FromLongitude        float64  `protobuf:"fixed64,7,opt,name=from_longitude,json=fromLongitude,proto3" json:"from_longitude,omitempty"`
	ToLatitude           float64  `protobuf:"fixed64,8,opt,name=to_latitude,json=toLatitude,proto3" json:"to_latitude,omitempty"`
	ToLongitude          float64  `protobuf:"fixed64,9,opt,name=to_longitude,json=toLongitude,proto3" json:"to_longitude,omitempty"`
	Address              string   `protobuf:"bytes,10,opt,name=address,proto3" json:"address,omitempty"`
	Status               int32    `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvoiceData) Reset()         { *m = InvoiceData{} }
func (m *InvoiceData) String() string { return proto.CompactTextString(m) }
func (*InvoiceData) ProtoMessage()    {}
func (*InvoiceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_24a7bef299e19e38, []int{0}
}

func (m *InvoiceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceData.Unmarshal(m, b)
}
func (m *InvoiceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceData.Marshal(b, m, deterministic)
}
func (m *InvoiceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceData.Merge(m, src)
}
func (m *InvoiceData) XXX_Size() int {
	return xxx_messageInfo_InvoiceData.Size(m)
}
func (m *InvoiceData) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceData.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceData proto.InternalMessageInfo

func (m *InvoiceData) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InvoiceData) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *InvoiceData) GetDriverId() int32 {
	if m != nil {
		return m.DriverId
	}
	return 0
}

func (m *InvoiceData) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *InvoiceData) GetShippingCost() int32 {
	if m != nil {
		return m.ShippingCost
	}
	return 0
}

func (m *InvoiceData) GetFromLatitude() float64 {
	if m != nil {
		return m.FromLatitude
	}
	return 0
}

func (m *InvoiceData) GetFromLongitude() float64 {
	if m != nil {
		return m.FromLongitude
	}
	return 0
}

func (m *InvoiceData) GetToLatitude() float64 {
	if m != nil {
		return m.ToLatitude
	}
	return 0
}

func (m *InvoiceData) GetToLongitude() float64 {
	if m != nil {
		return m.ToLongitude
	}
	return 0
}

func (m *InvoiceData) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *InvoiceData) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type InvoiceFilter struct {
	DriverId             int32    `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId            int32    `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	FromLatitude         float64  `protobuf:"fixed64,4,opt,name=from_latitude,json=fromLatitude,proto3" json:"from_latitude,omitempty"`
	FromLongitude        float64  `protobuf:"fixed64,5,opt,name=from_longitude,json=fromLongitude,proto3" json:"from_longitude,omitempty"`
	ToLatitude           float64  `protobuf:"fixed64,6,opt,name=to_latitude,json=toLatitude,proto3" json:"to_latitude,omitempty"`
	ToLongitude          float64  `protobuf:"fixed64,7,opt,name=to_longitude,json=toLongitude,proto3" json:"to_longitude,omitempty"`
	Address              string   `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvoiceFilter) Reset()         { *m = InvoiceFilter{} }
func (m *InvoiceFilter) String() string { return proto.CompactTextString(m) }
func (*InvoiceFilter) ProtoMessage()    {}
func (*InvoiceFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_24a7bef299e19e38, []int{1}
}

func (m *InvoiceFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceFilter.Unmarshal(m, b)
}
func (m *InvoiceFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceFilter.Marshal(b, m, deterministic)
}
func (m *InvoiceFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceFilter.Merge(m, src)
}
func (m *InvoiceFilter) XXX_Size() int {
	return xxx_messageInfo_InvoiceFilter.Size(m)
}
func (m *InvoiceFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceFilter.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceFilter proto.InternalMessageInfo

func (m *InvoiceFilter) GetDriverId() int32 {
	if m != nil {
		return m.DriverId
	}
	return 0
}

func (m *InvoiceFilter) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *InvoiceFilter) GetProductId() int32 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *InvoiceFilter) GetFromLatitude() float64 {
	if m != nil {
		return m.FromLatitude
	}
	return 0
}

func (m *InvoiceFilter) GetFromLongitude() float64 {
	if m != nil {
		return m.FromLongitude
	}
	return 0
}

func (m *InvoiceFilter) GetToLatitude() float64 {
	if m != nil {
		return m.ToLatitude
	}
	return 0
}

func (m *InvoiceFilter) GetToLongitude() float64 {
	if m != nil {
		return m.ToLongitude
	}
	return 0
}

func (m *InvoiceFilter) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*InvoiceData)(nil), "InvoiceData")
	proto.RegisterType((*InvoiceFilter)(nil), "InvoiceFilter")
}

func init() {
	proto.RegisterFile("grpc_seller/proto_invoice/invoice.proto", fileDescriptor_24a7bef299e19e38)
}

var fileDescriptor_24a7bef299e19e38 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x6f, 0x7b, 0xa1, 0xa5, 0xa7, 0xa5, 0x8b, 0x59, 0xe8, 0x44, 0x63, 0x44, 0x08, 0x91,
	0x55, 0x9b, 0xe8, 0x13, 0x28, 0xc6, 0xa4, 0x09, 0x2b, 0xdc, 0xb9, 0x69, 0x6a, 0x67, 0xc4, 0x49,
	0x0a, 0xd3, 0xcc, 0x9c, 0xf2, 0x18, 0x3e, 0x9a, 0xcf, 0x64, 0x3a, 0x2d, 0x20, 0x12, 0x95, 0x55,
	0x7b, 0xbe, 0xf9, 0xe7, 0x4f, 0xe6, 0xcb, 0x81, 0xeb, 0x85, 0x2a, 0xf3, 0x54, 0xf3, 0xa2, 0xe0,
	0x2a, 0x2e, 0x95, 0x44, 0x99, 0x8a, 0xd5, 0x5a, 0x8a, 0x9c, 0xc7, 0xed, 0x37, 0x32, 0x74, 0xf8,
	0x61, 0x83, 0x9f, 0x34, 0xe4, 0x21, 0xc3, 0x8c, 0x84, 0x60, 0x0b, 0x46, 0xad, 0x81, 0x35, 0xe9,
	0xce, 0x6d, 0xc1, 0xc8, 0x29, 0xb8, 0x95, 0xe6, 0x2a, 0x15, 0x8c, 0xda, 0x06, 0x3a, 0xf5, 0x98,
	0x30, 0x72, 0x0e, 0x1e, 0x53, 0x62, 0xdd, 0x1c, 0xfd, 0x37, 0x47, 0xbd, 0x06, 0x24, 0x8c, 0x10,
	0xe8, 0xe4, 0x92, 0x71, 0xda, 0x19, 0x58, 0x13, 0x6f, 0x6e, 0xfe, 0xc9, 0x08, 0xfa, 0xfa, 0x4d,
	0x94, 0xa5, 0x58, 0x2d, 0xd2, 0x5c, 0x6a, 0xa4, 0x5d, 0x73, 0x29, 0xd8, 0xc0, 0xa9, 0xd4, 0x58,
	0x87, 0x5e, 0x95, 0x5c, 0xa6, 0x45, 0x86, 0x02, 0x2b, 0xc6, 0xa9, 0x33, 0xb0, 0x26, 0xd6, 0x3c,
	0xa8, 0xe1, 0xac, 0x65, 0x64, 0x0c, 0x61, 0x13, 0x92, 0xab, 0x45, 0x93, 0x72, 0x4d, 0xca, 0x5c,
	0x9d, 0x6d, 0x20, 0xb9, 0x04, 0x1f, 0xe5, 0xae, 0xa9, 0x67, 0x32, 0x80, 0x72, 0xdb, 0x73, 0x05,
	0x41, 0x1d, 0xd8, 0xb6, 0x78, 0x26, 0xe1, 0xa3, 0xdc, 0x75, 0x50, 0x70, 0x33, 0xc6, 0x14, 0xd7,
	0x9a, 0x82, 0x79, 0xcb, 0x66, 0x24, 0x27, 0xe0, 0x68, 0xcc, 0xb0, 0xd2, 0xd4, 0x6f, 0xbc, 0x34,
	0xd3, 0xf0, 0xdd, 0x86, 0x7e, 0x2b, 0xf4, 0x51, 0x14, 0xc8, 0xd5, 0xbe, 0x29, 0xeb, 0x9b, 0xa9,
	0x1f, 0xfd, 0x5e, 0x00, 0x94, 0x4a, 0xb2, 0x2a, 0xc7, 0x9d, 0x60, 0xaf, 0x25, 0x09, 0x3b, 0x14,
	0xd5, 0x39, 0x4a, 0x54, 0xf7, 0x08, 0x51, 0xce, 0x9f, 0xa2, 0xdc, 0x5f, 0x45, 0xf5, 0xf6, 0x44,
	0xdd, 0xdc, 0x41, 0xd8, 0xfa, 0x78, 0xe2, 0x6a, 0x2d, 0x72, 0x4e, 0x62, 0xe8, 0x4f, 0x15, 0xcf,
	0x90, 0xb7, 0x9c, 0x84, 0xd1, 0x9e, 0xb1, 0xb3, 0x20, 0xfa, 0xb2, 0x92, 0xc3, 0x7f, 0xf7, 0xe3,
	0xe7, 0x51, 0x9c, 0x57, 0x1a, 0xe5, 0x92, 0xab, 0xf8, 0x60, 0xb3, 0xa3, 0x76, 0xa3, 0x5f, 0x1c,
	0x33, 0xde, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x05, 0xf0, 0x04, 0xfd, 0x02, 0x00, 0x00,
}
