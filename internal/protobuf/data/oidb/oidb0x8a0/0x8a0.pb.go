// Code generated by yaprotoc. DO NOT EDIT.
// source: oidb\0x8a0.proto

package oidb0x8a0

import (
	"github.com/pkg/errors"
	"github.com/segmentio/encoding/proto"
)

type KickMemberInfo struct {
	Operate   *uint32 `protobuf:"varint,1,opt"`
	MemberUin *uint64 `protobuf:"varint,2,opt"`
	Flag      *uint32 `protobuf:"varint,3,opt"`
	Msg       []byte  `protobuf:"bytes,4,opt"`
}

func (x *KickMemberInfo) GetOperate() uint32 {
	if x != nil && x.Operate != nil {
		return *x.Operate
	}
	return 0
}

func (x *KickMemberInfo) GetMemberUin() uint64 {
	if x != nil && x.MemberUin != nil {
		return *x.MemberUin
	}
	return 0
}

func (x *KickMemberInfo) GetFlag() uint32 {
	if x != nil && x.Flag != nil {
		return *x.Flag
	}
	return 0
}

func (x *KickMemberInfo) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type KickResult struct {
	Result    *uint32 `protobuf:"varint,1,opt"`
	MemberUin *uint64 `protobuf:"varint,2,opt"`
}

func (x *KickResult) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *KickResult) GetMemberUin() uint64 {
	if x != nil && x.MemberUin != nil {
		return *x.MemberUin
	}
	return 0
}

func (x *KickResult) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type ReqBody struct {
	GroupCode   *uint64           `protobuf:"varint,1,opt"`
	MsgKickList []*KickMemberInfo `protobuf:"bytes,2,rep"`
	KickList    []uint64          `protobuf:"varint,3,rep"`
	KickFlag    *uint32           `protobuf:"varint,4,opt"`
	KickMsg     []byte            `protobuf:"bytes,5,opt"`
}

func (x *ReqBody) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *ReqBody) GetKickFlag() uint32 {
	if x != nil && x.KickFlag != nil {
		return *x.KickFlag
	}
	return 0
}

func (x *ReqBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type RspBody struct {
	GroupCode  *uint64       `protobuf:"varint,1,opt"`
	KickResult []*KickResult `protobuf:"bytes,2,rep"`
}

func (x *RspBody) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *RspBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}
