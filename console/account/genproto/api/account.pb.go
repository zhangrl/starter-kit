// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/account.proto

package go_micro_api_console_account

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/stack-labs/stack-rpc/api/proto"
	srv "github.com/stack-labs/starter-kit/console/account/genproto/srv"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f8c2a44ebe43c87, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Response struct {
	Code                 int64              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string             `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	Data                 *srv.LoginResponse `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f8c2a44ebe43c87, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *Response) GetData() *srv.LoginResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "go.micro.api.console.account.LoginRequest")
	proto.RegisterType((*Response)(nil), "go.micro.api.console.account.Response")
}

func init() { proto.RegisterFile("api/account.proto", fileDescriptor_3f8c2a44ebe43c87) }

var fileDescriptor_3f8c2a44ebe43c87 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0x87, 0xdf, 0xbc, 0x8d, 0xb1, 0x5d, 0x05, 0xed, 0x22, 0x52, 0x8a, 0x87, 0x50, 0x51, 0x82,
	0x92, 0x0d, 0xb4, 0x07, 0xcf, 0xed, 0x4d, 0xe8, 0x69, 0xcf, 0x7a, 0xd8, 0x6e, 0xd6, 0xb8, 0x98,
	0xee, 0xac, 0xbb, 0x9b, 0x68, 0xbf, 0x9f, 0x9f, 0xca, 0x93, 0xe4, 0x4f, 0x63, 0x45, 0x28, 0xde,
	0x7e, 0xc9, 0x3c, 0xf3, 0x64, 0x32, 0x83, 0x86, 0x4c, 0xcb, 0x84, 0x71, 0x0e, 0x85, 0x72, 0x44,
	0x1b, 0x70, 0x80, 0x2f, 0x32, 0x20, 0x6b, 0xc9, 0x0d, 0x10, 0xa6, 0x25, 0xe1, 0xa0, 0x2c, 0xe4,
	0x82, 0xb4, 0xcc, 0x78, 0x96, 0x49, 0xf7, 0x5c, 0xac, 0x08, 0x87, 0x75, 0x62, 0x1d, 0xe3, 0x2f,
	0x71, 0xce, 0x56, 0xb6, 0x8d, 0x46, 0xf3, 0xa4, 0x12, 0xd6, 0xa2, 0x2a, 0x35, 0xca, 0xf1, 0x7c,
	0xa7, 0x49, 0xa8, 0x12, 0x36, 0xda, 0xc0, 0xfb, 0xa6, 0xc1, 0x78, 0x9c, 0x09, 0x15, 0x97, 0x2c,
	0x97, 0x29, 0x73, 0x22, 0xf9, 0x15, 0x5a, 0xc5, 0xd0, 0x9a, 0xf2, 0xe7, 0xa0, 0x93, 0x07, 0x74,
	0xbc, 0x84, 0x4c, 0x2a, 0x2a, 0x5e, 0x0b, 0x61, 0x1d, 0xbe, 0x42, 0xfd, 0xc2, 0x0a, 0xa3, 0xd8,
	0x5a, 0x8c, 0xbc, 0xd0, 0x8b, 0x06, 0x8b, 0xc1, 0xe7, 0x22, 0x30, 0x7e, 0xe8, 0x47, 0x21, 0xed,
	0x4a, 0x15, 0xa6, 0x99, 0xb5, 0x6f, 0x60, 0xd2, 0xd1, 0xff, 0x5d, 0x2c, 0x88, 0xce, 0x68, 0x57,
	0x9a, 0x00, 0xea, 0x53, 0x61, 0x35, 0x28, 0x2b, 0x30, 0x46, 0x3e, 0x87, 0xb4, 0xb1, 0xf6, 0x68,
	0x9d, 0xf1, 0x39, 0x0a, 0x52, 0xe1, 0x98, 0xcc, 0x1b, 0x09, 0x6d, 0x9f, 0xf0, 0x1d, 0xf2, 0x53,
	0xe6, 0xd8, 0xa8, 0x17, 0x7a, 0xd1, 0xd1, 0xf4, 0x92, 0x74, 0xdb, 0xb4, 0xa6, 0xdc, 0x6e, 0x91,
	0xb4, 0x73, 0x37, 0x7a, 0x5a, 0x37, 0x4c, 0x3f, 0x3c, 0x74, 0x38, 0x6f, 0xea, 0xf8, 0x11, 0x1d,
	0xd4, 0x08, 0xbe, 0x21, 0xfb, 0xae, 0x41, 0x76, 0xff, 0x7f, 0x7c, 0xbd, 0x9f, 0xdd, 0x7e, 0x6e,
	0xf2, 0x0f, 0xc7, 0x28, 0x58, 0x42, 0x06, 0x85, 0xc3, 0x27, 0x55, 0x4f, 0x45, 0x6f, 0x25, 0xa7,
	0xdf, 0x2f, 0x3a, 0xfc, 0x16, 0xf9, 0xf7, 0xea, 0x09, 0xfe, 0x04, 0xaf, 0x82, 0xfa, 0x38, 0xb3,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0x87, 0x19, 0x32, 0x5a, 0x02, 0x00, 0x00,
}
