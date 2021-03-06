// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpc/psi.proto

package mpc

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

//VLReEncIDsRequest is a message sent in Vertical-Learning-PSI Re-Encrypt-IDSet requesting.
type VLPsiReEncIDsRequest struct {
	TaskID               string   `protobuf:"bytes,2,opt,name=taskID,proto3" json:"taskID,omitempty"`
	EncIDs               []byte   `protobuf:"bytes,3,opt,name=EncIDs,proto3" json:"EncIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VLPsiReEncIDsRequest) Reset()         { *m = VLPsiReEncIDsRequest{} }
func (m *VLPsiReEncIDsRequest) String() string { return proto.CompactTextString(m) }
func (*VLPsiReEncIDsRequest) ProtoMessage()    {}
func (*VLPsiReEncIDsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae5ab50a0bdfb628, []int{0}
}

func (m *VLPsiReEncIDsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VLPsiReEncIDsRequest.Unmarshal(m, b)
}
func (m *VLPsiReEncIDsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VLPsiReEncIDsRequest.Marshal(b, m, deterministic)
}
func (m *VLPsiReEncIDsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VLPsiReEncIDsRequest.Merge(m, src)
}
func (m *VLPsiReEncIDsRequest) XXX_Size() int {
	return xxx_messageInfo_VLPsiReEncIDsRequest.Size(m)
}
func (m *VLPsiReEncIDsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VLPsiReEncIDsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VLPsiReEncIDsRequest proto.InternalMessageInfo

func (m *VLPsiReEncIDsRequest) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

func (m *VLPsiReEncIDsRequest) GetEncIDs() []byte {
	if m != nil {
		return m.EncIDs
	}
	return nil
}

type VLPsiReEncIDsResponse struct {
	TaskID               string   `protobuf:"bytes,2,opt,name=taskID,proto3" json:"taskID,omitempty"`
	ReEncIDs             []byte   `protobuf:"bytes,3,opt,name=ReEncIDs,proto3" json:"ReEncIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VLPsiReEncIDsResponse) Reset()         { *m = VLPsiReEncIDsResponse{} }
func (m *VLPsiReEncIDsResponse) String() string { return proto.CompactTextString(m) }
func (*VLPsiReEncIDsResponse) ProtoMessage()    {}
func (*VLPsiReEncIDsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae5ab50a0bdfb628, []int{1}
}

func (m *VLPsiReEncIDsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VLPsiReEncIDsResponse.Unmarshal(m, b)
}
func (m *VLPsiReEncIDsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VLPsiReEncIDsResponse.Marshal(b, m, deterministic)
}
func (m *VLPsiReEncIDsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VLPsiReEncIDsResponse.Merge(m, src)
}
func (m *VLPsiReEncIDsResponse) XXX_Size() int {
	return xxx_messageInfo_VLPsiReEncIDsResponse.Size(m)
}
func (m *VLPsiReEncIDsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VLPsiReEncIDsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VLPsiReEncIDsResponse proto.InternalMessageInfo

func (m *VLPsiReEncIDsResponse) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

func (m *VLPsiReEncIDsResponse) GetReEncIDs() []byte {
	if m != nil {
		return m.ReEncIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*VLPsiReEncIDsRequest)(nil), "mpc.VLPsiReEncIDsRequest")
	proto.RegisterType((*VLPsiReEncIDsResponse)(nil), "mpc.VLPsiReEncIDsResponse")
}

func init() { proto.RegisterFile("mpc/psi.proto", fileDescriptor_ae5ab50a0bdfb628) }

var fileDescriptor_ae5ab50a0bdfb628 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x2d, 0x48, 0xd6,
	0x2f, 0x28, 0xce, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x48, 0x56, 0x72,
	0xe3, 0x12, 0x09, 0xf3, 0x09, 0x28, 0xce, 0x0c, 0x4a, 0x75, 0xcd, 0x4b, 0xf6, 0x74, 0x29, 0x0e,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe3, 0x62, 0x2b, 0x49, 0x2c, 0xce, 0xf6, 0x74,
	0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x40, 0xe2, 0x10, 0x85, 0x12, 0xcc, 0x0a,
	0x8c, 0x1a, 0x3c, 0x41, 0x50, 0x9e, 0x92, 0x37, 0x97, 0x28, 0x9a, 0x39, 0xc5, 0x05, 0xf9, 0x79,
	0xc5, 0xa9, 0x38, 0x0d, 0x92, 0xe2, 0xe2, 0x80, 0xa9, 0x85, 0x1a, 0x05, 0xe7, 0x3b, 0x19, 0x45,
	0x19, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x07, 0x24, 0xa6, 0xa4,
	0xe4, 0xa4, 0x42, 0x48, 0x28, 0xc7, 0x25, 0x24, 0x42, 0x3f, 0x25, 0x31, 0x53, 0x1f, 0xec, 0x91,
	0x62, 0xfd, 0xdc, 0x82, 0xe4, 0x24, 0x36, 0x30, 0xdb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc4,
	0x1b, 0x05, 0x0d, 0xe5, 0x00, 0x00, 0x00,
}
