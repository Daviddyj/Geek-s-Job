// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package apis

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

type PersonalInformationList struct {
	Items                []*PersonInformation `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PersonalInformationList) Reset()         { *m = PersonalInformationList{} }
func (m *PersonalInformationList) String() string { return proto.CompactTextString(m) }
func (*PersonalInformationList) ProtoMessage()    {}
func (*PersonalInformationList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *PersonalInformationList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonalInformationList.Unmarshal(m, b)
}
func (m *PersonalInformationList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonalInformationList.Marshal(b, m, deterministic)
}
func (m *PersonalInformationList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonalInformationList.Merge(m, src)
}
func (m *PersonalInformationList) XXX_Size() int {
	return xxx_messageInfo_PersonalInformationList.Size(m)
}
func (m *PersonalInformationList) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonalInformationList.DiscardUnknown(m)
}

var xxx_messageInfo_PersonalInformationList proto.InternalMessageInfo

func (m *PersonalInformationList) GetItems() []*PersonInformation {
	if m != nil {
		return m.Items
	}
	return nil
}

type PersonInformation struct {
	// @gotags: gorm:"primaryKey;column:id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;column:id"`
	// @gotags: gorm:"column:person_id"
	PersonId int64 `protobuf:"varint,2,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty" gorm:"column:person_id"`
	// @gotags: gorm:"column:personal_name"
	PersonalName string `protobuf:"bytes,3,opt,name=personal_name,json=personalName,proto3" json:"personal_name,omitempty" gorm:"column:personal_name"`
	// @gotags: gorm:"column:content"
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty" gorm:"column:content"`
	// @gotags: gorm:"column:show_time"
	ShowTime string `protobuf:"bytes,5,opt,name=show_time,json=showTime,proto3" json:"show_time,omitempty" gorm:"column:show_time"`
	// @gotags: gorm:"column:by_time_age"
	ByTimeAge int64 `protobuf:"varint,6,opt,name=by_time_age,json=byTimeAge,proto3" json:"by_time_age,omitempty" gorm:"column:by_time_age"`
	// @gotags: gorm:"column:by_time_height"
	ByTimeHeight float32 `protobuf:"fixed32,7,opt,name=by_time_height,json=byTimeHeight,proto3" json:"by_time_height,omitempty" gorm:"column:by_time_height"`
	// @gotags: gorm:"column:by_time_weight"
	ByTimeWeight float32 `protobuf:"fixed32,8,opt,name=by_time_weight,json=byTimeWeight,proto3" json:"by_time_weight,omitempty" gorm:"column:by_time_weight"`
	// @gotags: gorm:"column:by_time_fatrate"
	ByTimeFatrate float32 `protobuf:"fixed32,9,opt,name=by_time_fatrate,json=byTimeFatrate,proto3" json:"by_time_fatrate,omitempty" gorm:"column:by_time_fatrate"`
	// @gotags: gorm:"column:personal_sex"
	PersonalSex string `protobuf:"bytes,10,opt,name=personal_sex,json=personalSex,proto3" json:"personal_sex,omitempty" gorm:"column:personal_sex"`
	// @gotags: gorm:"column:visiable"
	Visiable             bool     `protobuf:"varint,11,opt,name=visiable,proto3" json:"visiable,omitempty" gorm:"column:visiable"`
	//XXX_NoUnkeyedLiteral struct{} `json:"-"`
	//XXX_unrecognized     []byte   `json:"-"`
	//XXX_sizecache        int32    `json:"-"`
}

func (m *PersonInformation) Reset()         { *m = PersonInformation{} }
func (m *PersonInformation) String() string { return proto.CompactTextString(m) }
func (*PersonInformation) ProtoMessage()    {}
func (*PersonInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}

func (m *PersonInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonInformation.Unmarshal(m, b)
}
func (m *PersonInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonInformation.Marshal(b, m, deterministic)
}
func (m *PersonInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonInformation.Merge(m, src)
}
func (m *PersonInformation) XXX_Size() int {
	return xxx_messageInfo_PersonInformation.Size(m)
}
func (m *PersonInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonInformation.DiscardUnknown(m)
}

var xxx_messageInfo_PersonInformation proto.InternalMessageInfo

func (m *PersonInformation) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PersonInformation) GetPersonId() int64 {
	if m != nil {
		return m.PersonId
	}
	return 0
}

func (m *PersonInformation) GetPersonalName() string {
	if m != nil {
		return m.PersonalName
	}
	return ""
}

func (m *PersonInformation) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *PersonInformation) GetShowTime() string {
	if m != nil {
		return m.ShowTime
	}
	return ""
}

func (m *PersonInformation) GetByTimeAge() int64 {
	if m != nil {
		return m.ByTimeAge
	}
	return 0
}

func (m *PersonInformation) GetByTimeHeight() float32 {
	if m != nil {
		return m.ByTimeHeight
	}
	return 0
}

func (m *PersonInformation) GetByTimeWeight() float32 {
	if m != nil {
		return m.ByTimeWeight
	}
	return 0
}

func (m *PersonInformation) GetByTimeFatrate() float32 {
	if m != nil {
		return m.ByTimeFatrate
	}
	return 0
}

func (m *PersonInformation) GetPersonalSex() string {
	if m != nil {
		return m.PersonalSex
	}
	return ""
}

func (m *PersonInformation) GetVisiable() bool {
	if m != nil {
		return m.Visiable
	}
	return false
}

func init() {
	proto.RegisterType((*PersonalInformationList)(nil), "apis.PersonalInformationList")
	proto.RegisterType((*PersonInformation)(nil), "apis.PersonInformation")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xdd, 0x4a, 0xc3, 0x30,
	0x14, 0x80, 0x69, 0xf7, 0xd7, 0x9e, 0x6e, 0x13, 0x73, 0xb3, 0xa0, 0x20, 0x75, 0x8a, 0xf4, 0xc6,
	0x5d, 0xe8, 0x13, 0x78, 0x23, 0x1b, 0x88, 0x48, 0x15, 0xbc, 0x2c, 0xa9, 0x3d, 0xdb, 0x0e, 0xac,
	0x49, 0x69, 0x82, 0xdb, 0x9e, 0xda, 0x57, 0x90, 0x26, 0x76, 0x14, 0xbc, 0x3c, 0xdf, 0xf7, 0xb5,
	0x39, 0x21, 0x10, 0x99, 0x63, 0x85, 0x7a, 0x51, 0xd5, 0xca, 0x28, 0xd6, 0x17, 0x15, 0xe9, 0xf9,
	0x12, 0x66, 0x6f, 0x58, 0x6b, 0x25, 0xc5, 0x6e, 0x25, 0xd7, 0xaa, 0x2e, 0x85, 0x21, 0x25, 0x5f,
	0x48, 0x1b, 0x76, 0x0f, 0x03, 0x32, 0x58, 0x6a, 0xee, 0xc5, 0xbd, 0x24, 0x7a, 0x98, 0x2d, 0x9a,
	0x0f, 0x16, 0xae, 0xee, 0xb4, 0xa9, 0xab, 0xe6, 0x3f, 0x3e, 0x9c, 0xff, 0x93, 0x6c, 0x0a, 0x3e,
	0x15, 0xdc, 0x8b, 0xbd, 0xa4, 0x97, 0xfa, 0x54, 0xb0, 0x4b, 0x08, 0x2b, 0x1b, 0x65, 0x54, 0x70,
	0xdf, 0xe2, 0xc0, 0x81, 0x55, 0xc1, 0x6e, 0x60, 0x52, 0xfd, 0x2d, 0x93, 0x49, 0x51, 0x22, 0xef,
	0xc5, 0x5e, 0x12, 0xa6, 0xe3, 0x16, 0xbe, 0x8a, 0x12, 0x19, 0x87, 0xd1, 0x97, 0x92, 0x06, 0xa5,
	0xe1, 0x7d, 0xab, 0xdb, 0xb1, 0xf9, 0xb7, 0xde, 0xaa, 0x7d, 0x66, 0xa8, 0x44, 0x3e, 0xb0, 0x2e,
	0x68, 0xc0, 0x07, 0x95, 0xc8, 0xae, 0x20, 0xca, 0x8f, 0x56, 0x65, 0x62, 0x83, 0x7c, 0x68, 0x8f,
	0x0e, 0xf3, 0x63, 0x23, 0x9f, 0x36, 0xc8, 0x6e, 0x61, 0xda, 0xfa, 0x2d, 0xd2, 0x66, 0x6b, 0xf8,
	0x28, 0xf6, 0x12, 0x3f, 0x1d, 0xbb, 0x64, 0x69, 0x59, 0xb7, 0xda, 0xbb, 0x2a, 0xe8, 0x56, 0x9f,
	0xae, 0xba, 0x83, 0xb3, 0xb6, 0x5a, 0x0b, 0x53, 0x0b, 0x83, 0x3c, 0xb4, 0xd9, 0xc4, 0x65, 0xcf,
	0x0e, 0xb2, 0x6b, 0x38, 0x5d, 0x2d, 0xd3, 0x78, 0xe0, 0x60, 0x77, 0x8e, 0x5a, 0xf6, 0x8e, 0x07,
	0x76, 0x01, 0xc1, 0x37, 0x69, 0x12, 0xf9, 0x0e, 0x79, 0x14, 0x7b, 0x49, 0x90, 0x9e, 0xe6, 0x7c,
	0x68, 0x1f, 0xf2, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x32, 0xf9, 0xa7, 0x73, 0xd7, 0x01, 0x00,
	0x00,
}
