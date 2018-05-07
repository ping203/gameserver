// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmsg/client_msg.proto

/*
Package cmsg is a generated protocol buffer package.

It is generated from these files:
	cmsg/client_msg.proto

It has these top-level messages:
	CReqSyncSystemTime
	CRespSyncSystemTime
	CReqAuth
	CRespAuth
	CReqLogin
	CRespLogin
	CReqLogout
	CRespLogout
	CNotifyLogout
	CReqUserLoginState
	CRespUserLoginState
	CReqNotifyUserData
	CRespNotifyUserData
*/
package cmsg

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

// 同步系统时间
type CReqSyncSystemTime struct {
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *CReqSyncSystemTime) Reset()                    { *m = CReqSyncSystemTime{} }
func (m *CReqSyncSystemTime) String() string            { return proto.CompactTextString(m) }
func (*CReqSyncSystemTime) ProtoMessage()               {}
func (*CReqSyncSystemTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CReqSyncSystemTime) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type CRespSyncSystemTime struct {
	Timestamp       int64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	ServerTimestamp int64 `protobuf:"varint,2,opt,name=serverTimestamp" json:"serverTimestamp,omitempty"`
}

func (m *CRespSyncSystemTime) Reset()                    { *m = CRespSyncSystemTime{} }
func (m *CRespSyncSystemTime) String() string            { return proto.CompactTextString(m) }
func (*CRespSyncSystemTime) ProtoMessage()               {}
func (*CRespSyncSystemTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CRespSyncSystemTime) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *CRespSyncSystemTime) GetServerTimestamp() int64 {
	if m != nil {
		return m.ServerTimestamp
	}
	return 0
}

// 玩家认证请求
type CReqAuth struct {
	Account  string                    `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Password string                    `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Extra    *gamedef.ExtraAccountInfo `protobuf:"bytes,3,opt,name=extra" json:"extra,omitempty"`
}

func (m *CReqAuth) Reset()                    { *m = CReqAuth{} }
func (m *CReqAuth) String() string            { return proto.CompactTextString(m) }
func (*CReqAuth) ProtoMessage()               {}
func (*CReqAuth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CReqAuth) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *CReqAuth) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CReqAuth) GetExtra() *gamedef.ExtraAccountInfo {
	if m != nil {
		return m.Extra
	}
	return nil
}

type CRespAuth struct {
	ErrCode    uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg     string `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	Account    string `protobuf:"bytes,3,opt,name=account" json:"account,omitempty"`
	UserID     uint64 `protobuf:"varint,4,opt,name=userID" json:"userID,omitempty"`
	Sign       string `protobuf:"bytes,5,opt,name=sign" json:"sign,omitempty"`
	UnlockTime int64  `protobuf:"varint,6,opt,name=unlockTime" json:"unlockTime,omitempty"`
}

func (m *CRespAuth) Reset()                    { *m = CRespAuth{} }
func (m *CRespAuth) String() string            { return proto.CompactTextString(m) }
func (*CRespAuth) ProtoMessage()               {}
func (*CRespAuth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CRespAuth) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *CRespAuth) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *CRespAuth) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *CRespAuth) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CRespAuth) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *CRespAuth) GetUnlockTime() int64 {
	if m != nil {
		return m.UnlockTime
	}
	return 0
}

// 玩家登录请求
type CReqLogin struct {
	UserID uint64 `protobuf:"varint,1,opt,name=userID" json:"userID,omitempty"`
	Sign   string `protobuf:"bytes,2,opt,name=sign" json:"sign,omitempty"`
}

func (m *CReqLogin) Reset()                    { *m = CReqLogin{} }
func (m *CReqLogin) String() string            { return proto.CompactTextString(m) }
func (*CReqLogin) ProtoMessage()               {}
func (*CReqLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CReqLogin) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CReqLogin) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

type CRespLogin struct {
	ErrCode     uint32        `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg      string        `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	UserID      uint64        `protobuf:"varint,3,opt,name=userID" json:"userID,omitempty"`
	User        *gamedef.User `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	IsReconnect bool          `protobuf:"varint,5,opt,name=isReconnect" json:"isReconnect,omitempty"`
}

func (m *CRespLogin) Reset()                    { *m = CRespLogin{} }
func (m *CRespLogin) String() string            { return proto.CompactTextString(m) }
func (*CRespLogin) ProtoMessage()               {}
func (*CRespLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CRespLogin) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *CRespLogin) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *CRespLogin) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CRespLogin) GetUser() *gamedef.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CRespLogin) GetIsReconnect() bool {
	if m != nil {
		return m.IsReconnect
	}
	return false
}

// 玩家登出请求
type CReqLogout struct {
}

func (m *CReqLogout) Reset()                    { *m = CReqLogout{} }
func (m *CReqLogout) String() string            { return proto.CompactTextString(m) }
func (*CReqLogout) ProtoMessage()               {}
func (*CReqLogout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type CRespLogout struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *CRespLogout) Reset()                    { *m = CRespLogout{} }
func (m *CRespLogout) String() string            { return proto.CompactTextString(m) }
func (*CRespLogout) ProtoMessage()               {}
func (*CRespLogout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CRespLogout) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *CRespLogout) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

// 玩家被提出
type CNotifyLogout struct {
	LoginInfo *gamedef.LoginInfo `protobuf:"bytes,1,opt,name=loginInfo" json:"loginInfo,omitempty"`
}

func (m *CNotifyLogout) Reset()                    { *m = CNotifyLogout{} }
func (m *CNotifyLogout) String() string            { return proto.CompactTextString(m) }
func (*CNotifyLogout) ProtoMessage()               {}
func (*CNotifyLogout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CNotifyLogout) GetLoginInfo() *gamedef.LoginInfo {
	if m != nil {
		return m.LoginInfo
	}
	return nil
}

// 玩家登录状态请求
type CReqUserLoginState struct {
	UserID uint64 `protobuf:"varint,1,opt,name=userID" json:"userID,omitempty"`
	Sign   string `protobuf:"bytes,2,opt,name=sign" json:"sign,omitempty"`
}

func (m *CReqUserLoginState) Reset()                    { *m = CReqUserLoginState{} }
func (m *CReqUserLoginState) String() string            { return proto.CompactTextString(m) }
func (*CReqUserLoginState) ProtoMessage()               {}
func (*CReqUserLoginState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CReqUserLoginState) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CReqUserLoginState) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

type CRespUserLoginState struct {
	ErrCode   uint32             `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg    string             `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	UserID    uint64             `protobuf:"varint,3,opt,name=userID" json:"userID,omitempty"`
	IsLogin   bool               `protobuf:"varint,4,opt,name=isLogin" json:"isLogin,omitempty"`
	LoginInfo *gamedef.LoginInfo `protobuf:"bytes,5,opt,name=loginInfo" json:"loginInfo,omitempty"`
}

func (m *CRespUserLoginState) Reset()                    { *m = CRespUserLoginState{} }
func (m *CRespUserLoginState) String() string            { return proto.CompactTextString(m) }
func (*CRespUserLoginState) ProtoMessage()               {}
func (*CRespUserLoginState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CRespUserLoginState) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *CRespUserLoginState) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *CRespUserLoginState) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CRespUserLoginState) GetIsLogin() bool {
	if m != nil {
		return m.IsLogin
	}
	return false
}

func (m *CRespUserLoginState) GetLoginInfo() *gamedef.LoginInfo {
	if m != nil {
		return m.LoginInfo
	}
	return nil
}

// 请求玩家数据
type CReqNotifyUserData struct {
}

func (m *CReqNotifyUserData) Reset()                    { *m = CReqNotifyUserData{} }
func (m *CReqNotifyUserData) String() string            { return proto.CompactTextString(m) }
func (*CReqNotifyUserData) ProtoMessage()               {}
func (*CReqNotifyUserData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

// 返回玩家数据
type CRespNotifyUserData struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *CRespNotifyUserData) Reset()                    { *m = CRespNotifyUserData{} }
func (m *CRespNotifyUserData) String() string            { return proto.CompactTextString(m) }
func (*CRespNotifyUserData) ProtoMessage()               {}
func (*CRespNotifyUserData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *CRespNotifyUserData) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *CRespNotifyUserData) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*CReqSyncSystemTime)(nil), "cmsg.CReqSyncSystemTime")
	proto.RegisterType((*CRespSyncSystemTime)(nil), "cmsg.CRespSyncSystemTime")
	proto.RegisterType((*CReqAuth)(nil), "cmsg.CReqAuth")
	proto.RegisterType((*CRespAuth)(nil), "cmsg.CRespAuth")
	proto.RegisterType((*CReqLogin)(nil), "cmsg.CReqLogin")
	proto.RegisterType((*CRespLogin)(nil), "cmsg.CRespLogin")
	proto.RegisterType((*CReqLogout)(nil), "cmsg.CReqLogout")
	proto.RegisterType((*CRespLogout)(nil), "cmsg.CRespLogout")
	proto.RegisterType((*CNotifyLogout)(nil), "cmsg.CNotifyLogout")
	proto.RegisterType((*CReqUserLoginState)(nil), "cmsg.CReqUserLoginState")
	proto.RegisterType((*CRespUserLoginState)(nil), "cmsg.CRespUserLoginState")
	proto.RegisterType((*CReqNotifyUserData)(nil), "cmsg.CReqNotifyUserData")
	proto.RegisterType((*CRespNotifyUserData)(nil), "cmsg.CRespNotifyUserData")
}

func init() { proto.RegisterFile("cmsg/client_msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0x34, 0x4d, 0x26, 0x44, 0x48, 0x0b, 0x54, 0xa6, 0x42, 0x28, 0xec, 0x29, 0xa7,
	0x18, 0xc2, 0x01, 0xf5, 0x04, 0x51, 0x8a, 0x50, 0xa5, 0xc0, 0x61, 0x53, 0x8e, 0xa8, 0x72, 0x9d,
	0x89, 0xbb, 0x22, 0xde, 0x4d, 0x76, 0xd7, 0xb4, 0xf9, 0x2d, 0x1c, 0x39, 0xf2, 0x27, 0xd1, 0x8e,
	0xed, 0x7c, 0x54, 0x1c, 0xb0, 0xc4, 0x29, 0x33, 0x6f, 0x67, 0xdf, 0xbc, 0xf7, 0xa2, 0x35, 0x3c,
	0x4b, 0x32, 0x9b, 0x46, 0xc9, 0x52, 0xa2, 0x72, 0xd7, 0x99, 0x4d, 0x87, 0x2b, 0xa3, 0x9d, 0x66,
	0x4d, 0x0f, 0x9f, 0x4d, 0x53, 0xe9, 0x6e, 0xf3, 0x9b, 0x61, 0xa2, 0xb3, 0xe8, 0x0e, 0xd5, 0xbd,
	0xcc, 0x47, 0x6f, 0xce, 0xcf, 0xa3, 0x34, 0xce, 0xd0, 0xa2, 0xf9, 0x81, 0x26, 0xb2, 0x26, 0x89,
	0xca, 0xd2, 0xa3, 0x74, 0x99, 0xaa, 0x39, 0x2e, 0xe8, 0xf7, 0x7a, 0x8e, 0x8b, 0x82, 0x93, 0x8f,
	0x80, 0x4d, 0x04, 0xae, 0x67, 0x1b, 0x95, 0xcc, 0x36, 0xd6, 0x61, 0x76, 0x25, 0x33, 0x64, 0x2f,
	0xa0, 0xe3, 0x64, 0x86, 0xd6, 0xc5, 0xd9, 0x2a, 0x0c, 0xfa, 0xc1, 0xa0, 0x21, 0x76, 0x00, 0xff,
	0x06, 0x4f, 0x26, 0x02, 0xed, 0xaa, 0xce, 0x25, 0x36, 0x80, 0xc7, 0x85, 0xa6, 0xab, 0xed, 0xcc,
	0x11, 0xcd, 0x3c, 0x84, 0xf9, 0x1a, 0xda, 0x5e, 0xd2, 0x38, 0x77, 0xb7, 0x2c, 0x84, 0x93, 0x38,
	0x49, 0x74, 0xae, 0x1c, 0x31, 0x76, 0x44, 0xd5, 0xb2, 0x33, 0x68, 0xaf, 0x62, 0x6b, 0xef, 0xb4,
	0x99, 0x13, 0x51, 0x47, 0x6c, 0x7b, 0x16, 0xc1, 0x31, 0xde, 0x3b, 0x13, 0x87, 0x8d, 0x7e, 0x30,
	0xe8, 0x8e, 0x9e, 0x0f, 0x4b, 0xf3, 0xc3, 0x8f, 0x1e, 0x1d, 0x17, 0x0c, 0x97, 0x6a, 0xa1, 0x45,
	0x31, 0xc7, 0x7f, 0x05, 0xd0, 0x21, 0x4b, 0xd5, 0x52, 0x34, 0x66, 0xa2, 0xe7, 0x48, 0x4b, 0x7b,
	0xa2, 0x6a, 0xd9, 0x29, 0xb4, 0xd0, 0x98, 0xcf, 0x36, 0x2d, 0x57, 0x96, 0xdd, 0xbe, 0xcc, 0xc6,
	0xa1, 0xcc, 0x53, 0x68, 0xe5, 0x16, 0xcd, 0xe5, 0x45, 0xd8, 0xec, 0x07, 0x83, 0xa6, 0x28, 0x3b,
	0xc6, 0xa0, 0x69, 0x65, 0xaa, 0xc2, 0x63, 0x1a, 0xa7, 0x9a, 0xbd, 0x04, 0xc8, 0xd5, 0x52, 0x27,
	0xdf, 0x7d, 0x16, 0x61, 0x8b, 0xd2, 0xd9, 0x43, 0xf8, 0x3b, 0x12, 0xb9, 0x9e, 0xea, 0x54, 0xaa,
	0x3d, 0xe2, 0xe0, 0xaf, 0xc4, 0x47, 0x3b, 0x62, 0xfe, 0x33, 0x00, 0x20, 0x7b, 0xc5, 0xd5, 0xfa,
	0xfe, 0x76, 0xcb, 0x1a, 0x07, 0xcb, 0x5e, 0x41, 0xd3, 0x57, 0xe4, 0xad, 0x3b, 0xea, 0x6d, 0x73,
	0xfe, 0x6a, 0xd1, 0x08, 0x3a, 0x62, 0x7d, 0xe8, 0x4a, 0x2b, 0x30, 0xd1, 0x4a, 0x61, 0xe2, 0xc8,
	0x6f, 0x5b, 0xec, 0x43, 0xfc, 0x11, 0x89, 0xf3, 0xb6, 0x74, 0xee, 0xf8, 0x7b, 0xe8, 0x56, 0x52,
	0x75, 0xee, 0xea, 0x6b, 0xe5, 0x63, 0xe8, 0x4d, 0xbe, 0x68, 0x27, 0x17, 0x9b, 0x92, 0xe2, 0x35,
	0x74, 0x96, 0xde, 0xb7, 0xff, 0xc3, 0x89, 0xa4, 0x3b, 0x62, 0x5b, 0xa5, 0xd3, 0xea, 0x44, 0xec,
	0x86, 0xf8, 0x87, 0xe2, 0x51, 0x78, 0x17, 0x74, 0x3e, 0x73, 0xb1, 0xc3, 0x5a, 0x89, 0xff, 0x0e,
	0xca, 0x37, 0xf2, 0x80, 0xe3, 0xff, 0x45, 0x1f, 0xc2, 0x89, 0xb4, 0xc4, 0x4c, 0xe9, 0xb7, 0x45,
	0xd5, 0x1e, 0xfa, 0x3d, 0xfe, 0x17, 0xbf, 0x4f, 0x0b, 0xbf, 0x45, 0x6a, 0x5e, 0xf1, 0x45, 0xec,
	0x62, 0xfe, 0xa9, 0xb4, 0x70, 0x08, 0xd7, 0xb7, 0x70, 0xd3, 0xa2, 0x4f, 0xcd, 0xdb, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x20, 0xac, 0x86, 0x4e, 0xd7, 0x04, 0x00, 0x00,
}
