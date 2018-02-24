// Code generated by protoc-gen-go. DO NOT EDIT.
// source: defs/session.proto

package defs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Session struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=AccessToken" json:"AccessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=RefreshToken" json:"RefreshToken,omitempty"`
	CreatedTime  string `protobuf:"bytes,3,opt,name=CreatedTime" json:"CreatedTime,omitempty"`
	ExpireTime   string `protobuf:"bytes,4,opt,name=ExpireTime" json:"ExpireTime,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Session) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Session) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *Session) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

func (m *Session) GetExpireTime() string {
	if m != nil {
		return m.ExpireTime
	}
	return ""
}

func init() {
	proto.RegisterType((*Session)(nil), "defs.Session")
}

func init() { proto.RegisterFile("defs/session.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x65, 0xa8, 0x40, 0x98, 0x4e, 0x9e, 0x3a, 0xa1, 0x2a, 0x13, 0x4b, 0xed, 0x81, 0xbe,
	0x00, 0x45, 0xec, 0x55, 0xc9, 0xc4, 0xe6, 0x9f, 0x9b, 0xe4, 0x92, 0xc4, 0x8e, 0x7c, 0x13, 0x09,
	0xf1, 0x16, 0xbc, 0x31, 0xb2, 0x0d, 0x52, 0xd8, 0xec, 0x4f, 0xdf, 0x39, 0x3a, 0x97, 0x0b, 0x07,
	0x0d, 0x29, 0x02, 0x22, 0x0c, 0x5e, 0x4e, 0x31, 0xcc, 0x41, 0x6c, 0x12, 0xab, 0xbe, 0x19, 0xbf,
	0x7d, 0x2b, 0x5c, 0xec, 0xf9, 0xfd, 0xb3, 0xb5, 0x40, 0x54, 0x87, 0x1e, 0xfc, 0x8e, 0xed, 0xd9,
	0xe3, 0xdd, 0x65, 0x8d, 0x44, 0xc5, 0xb7, 0x17, 0x68, 0x22, 0x50, 0x57, 0x94, 0xab, 0xac, 0xfc,
	0x63, 0xa9, 0xe5, 0x25, 0x82, 0x9e, 0xc1, 0xd5, 0x38, 0xc2, 0xee, 0xba, 0xb4, 0xac, 0x90, 0x78,
	0xe0, 0xfc, 0xf5, 0x73, 0xc2, 0x08, 0x59, 0xd8, 0x64, 0x61, 0x45, 0x4e, 0x96, 0x57, 0x1e, 0xfd,
	0x87, 0x96, 0xa4, 0x7b, 0x34, 0x32, 0xea, 0x69, 0xf1, 0x5f, 0x30, 0x48, 0xed, 0x5d, 0x0c, 0xe8,
	0xca, 0xfe, 0xd3, 0xf6, 0x77, 0xf6, 0x39, 0xfd, 0xce, 0xec, 0xfd, 0xd0, 0xe2, 0xdc, 0x2d, 0x46,
	0xda, 0x30, 0x2a, 0x3a, 0xf6, 0x68, 0xe8, 0x38, 0xa2, 0xfa, 0x0b, 0x1f, 0xcc, 0x10, 0x5a, 0x95,
	0x93, 0x2a, 0x1d, 0x6e, 0x6e, 0xf2, 0xfb, 0xe9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x24, 0x5c, 0x77,
	0x9b, 0x1b, 0x01, 0x00, 0x00,
}