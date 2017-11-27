// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addr.proto

/*
Package addrBook is a generated protocol buffer package.

It is generated from these files:
	addr.proto

It has these top-level messages:
	Person
	AddressBook
*/
package addrBook

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PersonPhoneType int32

const (
	Person_MOBILE PersonPhoneType = 0
	Person_HOME   PersonPhoneType = 1
	Person_WORK   PersonPhoneType = 2
)

var PersonPhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var PersonPhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x PersonPhoneType) String() string {
	return proto.EnumName(PersonPhoneType_name, int32(x))
}
func (PersonPhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Person struct {
	Name    string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id      int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email   string                `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Number  []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=number" json:"number,omitempty"`
	Hashtbl map[string]string     `protobuf:"bytes,5,rep,name=hashtbl" json:"hashtbl,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

func (m *Person) GetNumber() []*Person_PhoneNumber {
	if m != nil {
		return m.Number
	}
	return nil
}

func (m *Person) GetHashtbl() map[string]string {
	if m != nil {
		return m.Hashtbl
	}
	return nil
}

type Person_PhoneNumber struct {
	Number string          `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Ptype  PersonPhoneType `protobuf:"varint,2,opt,name=ptype,enum=addrBook.PersonPhoneType" json:"ptype,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetPtype() PersonPhoneType {
	if m != nil {
		return m.Ptype
	}
	return Person_MOBILE
}

type AddressBook struct {
	Ppl []*Person `protobuf:"bytes,1,rep,name=ppl" json:"ppl,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddressBook) GetPpl() []*Person {
	if m != nil {
		return m.Ppl
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "addrBook.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "addrBook.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "addrBook.AddressBook")
	proto.RegisterEnum("addrBook.PersonPhoneType", PersonPhoneType_name, PersonPhoneType_value)
}

func init() { proto.RegisterFile("addr.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xd1, 0x6b, 0xc2, 0x30,
	0x10, 0xc6, 0x97, 0xd4, 0x76, 0x7a, 0x0e, 0x29, 0xc7, 0x18, 0x41, 0x36, 0x90, 0x3e, 0x09, 0x83,
	0xb2, 0xb9, 0xc1, 0x86, 0x6f, 0x13, 0x04, 0xc7, 0xe6, 0x94, 0x30, 0xf0, 0xb9, 0x92, 0x80, 0x62,
	0x6d, 0x42, 0x5a, 0x07, 0xfd, 0xbb, 0xf7, 0x0f, 0x8c, 0x24, 0x2a, 0x32, 0xdf, 0xbe, 0xe3, 0xbe,
	0xfb, 0xdd, 0x97, 0x1c, 0x40, 0x26, 0x84, 0x49, 0xb5, 0x51, 0x95, 0xc2, 0xa6, 0xd5, 0x23, 0xa5,
	0x36, 0xc9, 0x2f, 0x85, 0x68, 0x2e, 0x4d, 0xa9, 0x0a, 0x44, 0x68, 0x14, 0xd9, 0x56, 0x32, 0xd2,
	0x23, 0xfd, 0x16, 0x77, 0x1a, 0x3b, 0x40, 0xd7, 0x82, 0xd1, 0x1e, 0xe9, 0x87, 0x9c, 0xae, 0x05,
	0x5e, 0x43, 0x28, 0xb7, 0xd9, 0x3a, 0x67, 0x81, 0x33, 0xf9, 0x02, 0x9f, 0x21, 0x2a, 0x76, 0xdb,
	0xa5, 0x34, 0xac, 0xd1, 0x0b, 0xfa, 0xed, 0xc1, 0x6d, 0x7a, 0xe0, 0xa7, 0x9e, 0x9d, 0xce, 0x57,
	0xaa, 0x90, 0x5f, 0xce, 0xc3, 0xf7, 0x5e, 0x7c, 0x81, 0xcb, 0x55, 0x56, 0xae, 0xaa, 0x65, 0xce,
	0x42, 0x37, 0x76, 0x77, 0x36, 0x36, 0xf1, 0xfd, 0x71, 0x51, 0x99, 0x9a, 0x1f, 0xdc, 0xdd, 0x05,
	0xb4, 0x4f, 0x78, 0x78, 0x73, 0xdc, 0xee, 0x93, 0x1f, 0xf8, 0x0f, 0x10, 0xea, 0xaa, 0xd6, 0xd2,
	0xc5, 0xef, 0x0c, 0xba, 0x67, 0x74, 0x6d, 0x21, 0xdf, 0xb5, 0x96, 0xdc, 0x1b, 0xbb, 0x43, 0xb8,
	0x3a, 0xdd, 0x88, 0x31, 0x04, 0x1b, 0x59, 0xef, 0xb1, 0x56, 0xda, 0xf7, 0xff, 0x64, 0xf9, 0xce,
	0x33, 0x5b, 0xdc, 0x17, 0x43, 0xfa, 0x4a, 0x92, 0x7b, 0x68, 0x1d, 0x79, 0x08, 0x10, 0x4d, 0x67,
	0xa3, 0xf7, 0xcf, 0x71, 0x7c, 0x81, 0x4d, 0x68, 0x4c, 0x66, 0xd3, 0x71, 0x4c, 0xac, 0x5a, 0xcc,
	0xf8, 0x47, 0x4c, 0x93, 0x47, 0x68, 0xbf, 0x09, 0x61, 0x64, 0x59, 0xda, 0x3c, 0x98, 0x40, 0xa0,
	0x75, 0xce, 0x88, 0xfb, 0x85, 0xf8, 0x7f, 0x4e, 0x6e, 0x9b, 0xcb, 0xc8, 0x5d, 0xee, 0xe9, 0x2f,
	0x00, 0x00, 0xff, 0xff, 0x24, 0x7b, 0x8d, 0x28, 0xc7, 0x01, 0x00, 0x00,
}