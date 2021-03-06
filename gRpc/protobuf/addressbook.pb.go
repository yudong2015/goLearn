// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gRpc/addressbook.proto

package gRpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9f34c83cfe6008a5, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdate           *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f34c83cfe6008a5, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdate() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdate
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=gRpc.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f34c83cfe6008a5, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f34c83cfe6008a5, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("gRpc.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "gRpc.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "gRpc.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "gRpc.AddressBook")
}

func init() { proto.RegisterFile("gRpc/addressbook.proto", fileDescriptor_9f34c83cfe6008a5) }

var fileDescriptor_9f34c83cfe6008a5 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0x80, 0xdf, 0xa4, 0x69, 0x78, 0x3b, 0x91, 0x52, 0x06, 0x29, 0xa1, 0x17, 0x43, 0xf1, 0x10,
	0x14, 0xb6, 0xd2, 0x1e, 0x3d, 0x59, 0x28, 0x28, 0x5a, 0x5b, 0x97, 0x8a, 0x47, 0x49, 0xcc, 0x58,
	0x43, 0x93, 0xec, 0x92, 0xdd, 0x1e, 0xfa, 0xb3, 0xfc, 0x87, 0xb2, 0x9b, 0x54, 0x0a, 0xde, 0xe6,
	0xe3, 0x61, 0xe6, 0x99, 0x81, 0xe1, 0x96, 0xcb, 0x8f, 0x49, 0x92, 0x65, 0x35, 0x29, 0x95, 0x0a,
	0xb1, 0x63, 0xb2, 0x16, 0x5a, 0xa0, 0x67, 0xea, 0xa3, 0x8b, 0xad, 0x10, 0xdb, 0x82, 0x26, 0xb6,
	0x96, 0xee, 0x3f, 0x27, 0x3a, 0x2f, 0x49, 0xe9, 0xa4, 0x94, 0x0d, 0x36, 0xfe, 0x76, 0xc1, 0x5f,
	0x53, 0xad, 0x44, 0x85, 0x08, 0x5e, 0x95, 0x94, 0x14, 0x3a, 0x91, 0x13, 0xf7, 0xb8, 0x8d, 0xb1,
	0x0f, 0x6e, 0x9e, 0x85, 0x6e, 0xe4, 0xc4, 0x5d, 0xee, 0xe6, 0x19, 0x9e, 0x43, 0x97, 0xca, 0x24,
	0x2f, 0xc2, 0x8e, 0x85, 0x9a, 0x04, 0x6f, 0xc0, 0x97, 0x5f, 0xa2, 0x22, 0x15, 0x7a, 0x51, 0x27,
	0x0e, 0xa6, 0x21, 0x33, 0xcb, 0x59, 0x33, 0x97, 0xad, 0x4d, 0xeb, 0x79, 0x5f, 0xa6, 0x54, 0xf3,
	0x96, 0xc3, 0x5b, 0x08, 0x8a, 0x44, 0xe9, 0xf7, 0xbd, 0xcc, 0x12, 0x4d, 0x61, 0x37, 0x72, 0xe2,
	0x60, 0x3a, 0x62, 0x8d, 0x2d, 0x3b, 0xda, 0xb2, 0xcd, 0xd1, 0x96, 0x83, 0xc1, 0x5f, 0x2d, 0x3d,
	0x7a, 0x81, 0xe0, 0x64, 0x26, 0x0e, 0xc1, 0xaf, 0x6c, 0xd4, 0x9a, 0xb7, 0x19, 0x5e, 0x81, 0xa7,
	0x0f, 0x92, 0xac, 0x7d, 0x7f, 0x3a, 0xfc, 0xeb, 0xb4, 0x39, 0x48, 0xe2, 0x96, 0x19, 0x5f, 0x43,
	0xef, 0xb7, 0x84, 0x00, 0xfe, 0x72, 0x35, 0x7f, 0x78, 0x5a, 0x0c, 0xfe, 0xe1, 0x7f, 0xf0, 0xee,
	0x57, 0xcb, 0xc5, 0xc0, 0x31, 0xd1, 0xdb, 0x8a, 0x3f, 0x0e, 0xdc, 0xf1, 0x0c, 0x82, 0xbb, 0xe6,
	0xdf, 0x73, 0x21, 0x76, 0x78, 0x09, 0xbe, 0x24, 0x21, 0x0b, 0xf3, 0x39, 0x73, 0xfd, 0xd9, 0xe9,
	0x26, 0xde, 0xf6, 0x52, 0xdf, 0x1e, 0x35, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0x00, 0x53, 0x0b,
	0x8f, 0xb0, 0x01, 0x00, 0x00,
}
