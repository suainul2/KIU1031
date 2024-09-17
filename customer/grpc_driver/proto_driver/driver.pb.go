// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_driver/proto_driver/driver.proto

package proto_driver

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

type DriverData struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Role                 uint32   `protobuf:"varint,4,opt,name=role,proto3" json:"role,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Latitude             float64  `protobuf:"fixed64,6,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,7,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DriverData) Reset()         { *m = DriverData{} }
func (m *DriverData) String() string { return proto.CompactTextString(m) }
func (*DriverData) ProtoMessage()    {}
func (*DriverData) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfedaba5302dbbed, []int{0}
}

func (m *DriverData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverData.Unmarshal(m, b)
}
func (m *DriverData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverData.Marshal(b, m, deterministic)
}
func (m *DriverData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverData.Merge(m, src)
}
func (m *DriverData) XXX_Size() int {
	return xxx_messageInfo_DriverData.Size(m)
}
func (m *DriverData) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverData.DiscardUnknown(m)
}

var xxx_messageInfo_DriverData proto.InternalMessageInfo

func (m *DriverData) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DriverData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DriverData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *DriverData) GetRole() uint32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *DriverData) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *DriverData) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *DriverData) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type DriverFilter struct {
	Latitude             float64  `protobuf:"fixed64,6,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,7,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DriverFilter) Reset()         { *m = DriverFilter{} }
func (m *DriverFilter) String() string { return proto.CompactTextString(m) }
func (*DriverFilter) ProtoMessage()    {}
func (*DriverFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfedaba5302dbbed, []int{1}
}

func (m *DriverFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverFilter.Unmarshal(m, b)
}
func (m *DriverFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverFilter.Marshal(b, m, deterministic)
}
func (m *DriverFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverFilter.Merge(m, src)
}
func (m *DriverFilter) XXX_Size() int {
	return xxx_messageInfo_DriverFilter.Size(m)
}
func (m *DriverFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverFilter.DiscardUnknown(m)
}

var xxx_messageInfo_DriverFilter proto.InternalMessageInfo

func (m *DriverFilter) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *DriverFilter) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func init() {
	proto.RegisterType((*DriverData)(nil), "DriverData")
	proto.RegisterType((*DriverFilter)(nil), "DriverFilter")
}

func init() {
	proto.RegisterFile("grpc_driver/proto_driver/driver.proto", fileDescriptor_dfedaba5302dbbed)
}

var fileDescriptor_dfedaba5302dbbed = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0xd8, 0x54, 0x3b, 0x1a, 0x0f, 0xab, 0x87, 0xa5, 0x78, 0x08, 0x41, 0x21, 0xa7,
	0x04, 0xec, 0x1b, 0x94, 0xa0, 0x9e, 0xe3, 0xcd, 0x8b, 0xac, 0xc9, 0x50, 0x17, 0x36, 0xbb, 0x61,
	0xbb, 0x2d, 0xf8, 0x4a, 0x3e, 0xa5, 0x74, 0x06, 0x8d, 0xe0, 0xad, 0xa7, 0x9d, 0xef, 0x9b, 0x9f,
	0xd9, 0x61, 0xe0, 0x7e, 0x13, 0xc6, 0xee, 0xad, 0x0f, 0x66, 0x8f, 0xa1, 0x1e, 0x83, 0x8f, 0xfe,
	0x07, 0xf8, 0xa9, 0xc8, 0x15, 0x5f, 0x02, 0xa0, 0x21, 0xd1, 0xe8, 0xa8, 0xe5, 0x15, 0x24, 0xa6,
	0x57, 0x22, 0x17, 0x65, 0xda, 0x26, 0xa6, 0x97, 0x12, 0x66, 0x4e, 0x0f, 0xa8, 0x92, 0x5c, 0x94,
	0x8b, 0x96, 0x6a, 0x79, 0x03, 0x29, 0x0e, 0xda, 0x58, 0x75, 0x4a, 0x92, 0xe1, 0x90, 0x0c, 0xde,
	0xa2, 0x9a, 0xe5, 0xa2, 0xcc, 0x5a, 0xaa, 0x0f, 0xc9, 0xf1, 0xc3, 0x3b, 0x54, 0x29, 0x27, 0x09,
	0xe4, 0x12, 0xce, 0xad, 0x8e, 0x26, 0xee, 0x7a, 0x54, 0xf3, 0x5c, 0x94, 0xa2, 0xfd, 0x65, 0x79,
	0x0b, 0x0b, 0xeb, 0xdd, 0x86, 0x9b, 0x67, 0xd4, 0x9c, 0x44, 0xf1, 0x0c, 0x97, 0xbc, 0xeb, 0xa3,
	0xb1, 0x11, 0xc3, 0xf1, 0x93, 0x1e, 0x1a, 0xc8, 0x78, 0xd2, 0x0b, 0x86, 0xbd, 0xe9, 0x50, 0xae,
	0xe0, 0xfa, 0x09, 0x23, 0xbb, 0xf5, 0x67, 0x63, 0xb6, 0x51, 0xbb, 0x0e, 0x65, 0x56, 0xfd, 0xfd,
	0x70, 0x79, 0x51, 0x4d, 0xb7, 0x2a, 0x4e, 0xd6, 0x77, 0xaf, 0x45, 0xdd, 0xed, 0xb6, 0xd1, 0x0f,
	0x18, 0xea, 0x7f, 0xf7, 0xae, 0x18, 0xde, 0xe7, 0x44, 0xab, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xb5, 0xd0, 0x90, 0xd1, 0x92, 0x01, 0x00, 0x00,
}
