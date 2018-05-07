// Code generated by protoc-gen-go. DO NOT EDIT.
// source: smsg/gt_msg.proto

/*
Package smsg is a generated protocol buffer package.

It is generated from these files:
	smsg/gt_msg.proto

It has these top-level messages:
	GtLsReqAuth
	GtLsRespAuth
	GtGsReqLogin
	GtGsRespLogin
*/
package smsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import gamedef "github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GtLsReqAuth struct {
	Account  string `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *GtLsReqAuth) Reset()                    { *m = GtLsReqAuth{} }
func (m *GtLsReqAuth) String() string            { return proto.CompactTextString(m) }
func (*GtLsReqAuth) ProtoMessage()               {}
func (*GtLsReqAuth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GtLsReqAuth) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *GtLsReqAuth) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GtLsRespAuth struct {
	ErrCode    uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg     string `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	Account    string `protobuf:"bytes,3,opt,name=account" json:"account,omitempty"`
	UserID     uint64 `protobuf:"varint,4,opt,name=userID" json:"userID,omitempty"`
	Sign       string `protobuf:"bytes,5,opt,name=sign" json:"sign,omitempty"`
	UnlockTime int64  `protobuf:"varint,6,opt,name=unlockTime" json:"unlockTime,omitempty"`
}

func (m *GtLsRespAuth) Reset()                    { *m = GtLsRespAuth{} }
func (m *GtLsRespAuth) String() string            { return proto.CompactTextString(m) }
func (*GtLsRespAuth) ProtoMessage()               {}
func (*GtLsRespAuth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GtLsRespAuth) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *GtLsRespAuth) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *GtLsRespAuth) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *GtLsRespAuth) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *GtLsRespAuth) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *GtLsRespAuth) GetUnlockTime() int64 {
	if m != nil {
		return m.UnlockTime
	}
	return 0
}

type GtGsReqLogin struct {
	UserID uint64 `protobuf:"varint,1,opt,name=userID" json:"userID,omitempty"`
}

func (m *GtGsReqLogin) Reset()                    { *m = GtGsReqLogin{} }
func (m *GtGsReqLogin) String() string            { return proto.CompactTextString(m) }
func (*GtGsReqLogin) ProtoMessage()               {}
func (*GtGsReqLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GtGsReqLogin) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type GtGsRespLogin struct {
	ErrCode     uint32        `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg      string        `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	UserID      uint64        `protobuf:"varint,3,opt,name=userID" json:"userID,omitempty"`
	User        *gamedef.User `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	IsReconnect bool          `protobuf:"varint,5,opt,name=isReconnect" json:"isReconnect,omitempty"`
}

func (m *GtGsRespLogin) Reset()                    { *m = GtGsRespLogin{} }
func (m *GtGsRespLogin) String() string            { return proto.CompactTextString(m) }
func (*GtGsRespLogin) ProtoMessage()               {}
func (*GtGsRespLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GtGsRespLogin) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *GtGsRespLogin) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *GtGsRespLogin) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *GtGsRespLogin) GetUser() *gamedef.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GtGsRespLogin) GetIsReconnect() bool {
	if m != nil {
		return m.IsReconnect
	}
	return false
}

func init() {
	proto.RegisterType((*GtLsReqAuth)(nil), "smsg.GtLsReqAuth")
	proto.RegisterType((*GtLsRespAuth)(nil), "smsg.GtLsRespAuth")
	proto.RegisterType((*GtGsReqLogin)(nil), "smsg.GtGsReqLogin")
	proto.RegisterType((*GtGsRespLogin)(nil), "smsg.GtGsRespLogin")
}

func init() { proto.RegisterFile("smsg/gt_msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x15, 0x1a, 0x42, 0xb9, 0xd2, 0x01, 0x0f, 0x55, 0xd4, 0x01, 0x85, 0x0c, 0xa8, 0x53,
	0x22, 0xca, 0xd4, 0x11, 0x15, 0xa9, 0x42, 0x2a, 0x8b, 0x05, 0x73, 0xd5, 0x3a, 0x57, 0x37, 0x82,
	0xd8, 0xc1, 0xe7, 0x50, 0x7e, 0x0d, 0x23, 0xbf, 0x13, 0xc5, 0x49, 0x21, 0x1d, 0x99, 0xf2, 0xde,
	0xf3, 0xe5, 0xb3, 0x9f, 0x0e, 0x2e, 0xa9, 0x20, 0x99, 0x4a, 0xbb, 0x2a, 0x48, 0x26, 0xa5, 0xd1,
	0x56, 0x33, 0xbf, 0x8e, 0xc6, 0x4b, 0x99, 0xdb, 0x5d, 0xb5, 0x49, 0x84, 0x2e, 0xd2, 0x3d, 0xaa,
	0xcf, 0xbc, 0x9a, 0xde, 0xce, 0x66, 0xa9, 0x5c, 0x17, 0x48, 0x68, 0x3e, 0xd0, 0xa4, 0x64, 0x44,
	0xda, 0xca, 0x3a, 0x75, 0x3f, 0x3b, 0x95, 0xe1, 0xd6, 0x7d, 0x57, 0x19, 0x6e, 0x1b, 0x66, 0x3c,
	0x87, 0xc1, 0xc2, 0x2e, 0x89, 0xe3, 0xfb, 0x7d, 0x65, 0x77, 0x2c, 0x84, 0xb3, 0xb5, 0x10, 0xba,
	0x52, 0x36, 0xf4, 0x22, 0x6f, 0x72, 0xce, 0x0f, 0x96, 0x8d, 0xa1, 0x5f, 0xae, 0x89, 0xf6, 0xda,
	0x64, 0xe1, 0x89, 0x3b, 0xfa, 0xf5, 0xf1, 0xb7, 0x07, 0x17, 0x0d, 0x85, 0xca, 0x03, 0x06, 0x8d,
	0x99, 0xeb, 0x0c, 0x1d, 0x66, 0xc8, 0x0f, 0x96, 0x8d, 0x20, 0x40, 0x63, 0x9e, 0x48, 0xb6, 0x90,
	0xd6, 0x75, 0x2f, 0xee, 0x1d, 0x5f, 0x3c, 0x82, 0xa0, 0x22, 0x34, 0x8f, 0x0f, 0xa1, 0x1f, 0x79,
	0x13, 0x9f, 0xb7, 0x8e, 0x31, 0xf0, 0x29, 0x97, 0x2a, 0x3c, 0x75, 0xe3, 0x4e, 0xb3, 0x2b, 0x80,
	0x4a, 0xbd, 0x69, 0xf1, 0xfa, 0x9c, 0x17, 0x18, 0x06, 0x91, 0x37, 0xe9, 0xf1, 0x4e, 0x12, 0xdf,
	0xd4, 0xef, 0x5c, 0xd4, 0x6d, 0x97, 0x5a, 0xe6, 0xaa, 0xc3, 0xf6, 0xba, 0xec, 0xf8, 0xcb, 0x83,
	0x61, 0x33, 0x48, 0x65, 0x33, 0xf9, 0xff, 0x46, 0x7f, 0xec, 0xde, 0xd1, 0xbb, 0xaf, 0xc1, 0xaf,
	0x95, 0x6b, 0x33, 0x98, 0x0e, 0x93, 0x76, 0x31, 0xc9, 0x0b, 0xa1, 0xe1, 0xee, 0x88, 0x45, 0x30,
	0xc8, 0x89, 0xa3, 0xd0, 0x4a, 0xa1, 0xb0, 0xae, 0x61, 0x9f, 0x77, 0xa3, 0x4d, 0xe0, 0xb6, 0x77,
	0xf7, 0x13, 0x00, 0x00, 0xff, 0xff, 0x61, 0xb2, 0x6c, 0x85, 0x26, 0x02, 0x00, 0x00,
}