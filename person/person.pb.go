// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person/person.proto

package person

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	date "github.com/rhass99/st_pb/date"
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

type Person_EyeColour int32

const (
	Person_UNKNOWN Person_EyeColour = 0
	Person_GREEN   Person_EyeColour = 1
	Person_BROWN   Person_EyeColour = 2
)

var Person_EyeColour_name = map[int32]string{
	0: "UNKNOWN",
	1: "GREEN",
	2: "BROWN",
}

var Person_EyeColour_value = map[string]int32{
	"UNKNOWN": 0,
	"GREEN":   1,
	"BROWN":   2,
}

func (x Person_EyeColour) String() string {
	return proto.EnumName(Person_EyeColour_name, int32(x))
}

func (Person_EyeColour) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e25b6e050ad8ae9a, []int{0, 0}
}

type Person struct {
	Age                  int32            `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	FirstName            string           `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string           `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	SmallPicture         []byte           `protobuf:"bytes,4,opt,name=smallPicture,proto3" json:"smallPicture,omitempty"`
	IsProfileVerified    bool             `protobuf:"varint,5,opt,name=isProfileVerified,proto3" json:"isProfileVerified,omitempty"`
	Height               float32          `protobuf:"fixed32,6,opt,name=height,proto3" json:"height,omitempty"`
	PhoneNumber          []string         `protobuf:"bytes,7,rep,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	EyeColour            Person_EyeColour `protobuf:"varint,8,opt,name=eyeColour,proto3,enum=person.Person_EyeColour" json:"eyeColour,omitempty"`
	Date                 *date.Date       `protobuf:"bytes,9,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_e25b6e050ad8ae9a, []int{0}
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

func (m *Person) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Person) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Person) GetSmallPicture() []byte {
	if m != nil {
		return m.SmallPicture
	}
	return nil
}

func (m *Person) GetIsProfileVerified() bool {
	if m != nil {
		return m.IsProfileVerified
	}
	return false
}

func (m *Person) GetHeight() float32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Person) GetPhoneNumber() []string {
	if m != nil {
		return m.PhoneNumber
	}
	return nil
}

func (m *Person) GetEyeColour() Person_EyeColour {
	if m != nil {
		return m.EyeColour
	}
	return Person_UNKNOWN
}

func (m *Person) GetDate() *date.Date {
	if m != nil {
		return m.Date
	}
	return nil
}

func init() {
	proto.RegisterEnum("person.Person_EyeColour", Person_EyeColour_name, Person_EyeColour_value)
	proto.RegisterType((*Person)(nil), "person.Person")
}

func init() { proto.RegisterFile("person/person.proto", fileDescriptor_e25b6e050ad8ae9a) }

var fileDescriptor_e25b6e050ad8ae9a = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x86, 0xef, 0xa4, 0x6d, 0xda, 0x9c, 0x96, 0x7b, 0x7b, 0x47, 0x90, 0xa1, 0x88, 0x8e, 0x5d,
	0xcd, 0x42, 0x12, 0xa8, 0x20, 0x74, 0x5b, 0x2d, 0x2e, 0x84, 0x58, 0x06, 0x54, 0x70, 0x23, 0x49,
	0x7b, 0xda, 0x0c, 0x24, 0x9d, 0x30, 0x99, 0x2c, 0xfc, 0xc7, 0xfe, 0x0c, 0xc9, 0x87, 0xad, 0xe2,
	0x26, 0x39, 0xef, 0xf3, 0xcc, 0x2c, 0xe6, 0x3d, 0x70, 0x92, 0xa3, 0x29, 0xf4, 0x3e, 0x68, 0x7e,
	0x7e, 0x6e, 0xb4, 0xd5, 0xd4, 0x6d, 0xd2, 0xe4, 0xdf, 0x26, 0xb2, 0x18, 0x54, 0x9f, 0x46, 0x4c,
	0x3f, 0x1c, 0x70, 0x57, 0xb5, 0xa3, 0x63, 0xe8, 0x44, 0x3b, 0x64, 0x84, 0x13, 0xd1, 0x93, 0xd5,
	0x48, 0xcf, 0xc0, 0xdb, 0x2a, 0x53, 0xd8, 0x30, 0xca, 0x90, 0x39, 0x9c, 0x08, 0x4f, 0x1e, 0x01,
	0x9d, 0xc0, 0x20, 0x8d, 0x5a, 0xd9, 0xa9, 0xe5, 0x21, 0xd3, 0x29, 0x8c, 0x8a, 0x2c, 0x4a, 0xd3,
	0x95, 0x5a, 0xdb, 0xd2, 0x20, 0xeb, 0x72, 0x22, 0x46, 0xf2, 0x07, 0xa3, 0x57, 0xf0, 0x5f, 0x15,
	0x2b, 0xa3, 0xb7, 0x2a, 0xc5, 0x67, 0x34, 0x6a, 0xab, 0x70, 0xc3, 0x7a, 0x9c, 0x88, 0x81, 0xfc,
	0x2d, 0xe8, 0x29, 0xb8, 0x09, 0xaa, 0x5d, 0x62, 0x99, 0xcb, 0x89, 0x70, 0x64, 0x9b, 0x28, 0x87,
	0x61, 0x9e, 0xe8, 0x3d, 0x86, 0x65, 0x16, 0xa3, 0x61, 0x7d, 0xde, 0x11, 0x9e, 0xfc, 0x8e, 0xe8,
	0x0d, 0x78, 0xf8, 0x8e, 0xb7, 0x3a, 0xd5, 0xa5, 0x61, 0x03, 0x4e, 0xc4, 0xdf, 0x19, 0xf3, 0xdb,
	0x76, 0x9a, 0xa7, 0xfb, 0xcb, 0x2f, 0x2f, 0x8f, 0x47, 0xe9, 0x39, 0x74, 0xab, 0xa2, 0x98, 0xc7,
	0x89, 0x18, 0xce, 0xc0, 0xaf, 0x5b, 0xbb, 0x8b, 0x2c, 0xca, 0x9a, 0x4f, 0x7d, 0xf0, 0x0e, 0xf7,
	0xe8, 0x10, 0xfa, 0x4f, 0xe1, 0x43, 0xf8, 0xf8, 0x12, 0x8e, 0xff, 0x50, 0x0f, 0x7a, 0xf7, 0x72,
	0xb9, 0x0c, 0xc7, 0xa4, 0x1a, 0x17, 0xb2, 0xa2, 0xce, 0xe2, 0xf2, 0xf5, 0x62, 0xa7, 0x6c, 0x52,
	0xc6, 0xfe, 0x5a, 0x67, 0x81, 0x49, 0xa2, 0xa2, 0x98, 0xcf, 0x83, 0xc2, 0xbe, 0xe5, 0x71, 0xbb,
	0xac, 0xd8, 0xad, 0x97, 0x72, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xda, 0x16, 0x55, 0xc4,
	0x01, 0x00, 0x00,
}