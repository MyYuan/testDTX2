// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpc/learners/logic_reg_vl/logic_reg_vl.proto

package logic_reg_vl

import (
	fmt "fmt"
	common "github.com/PaddlePaddle/PaddleDTX/dai/protos/common"
	mpc "github.com/PaddlePaddle/PaddleDTX/dai/protos/mpc"
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

//MessageType defines the type of message with which communicate with nodes in cluster,
// and in some way it indicates the phase of learning
//Some types are for local message which is not passed between nodes
type MessageType int32

const (
	MessageType_MsgPsiEnc                MessageType = 0
	MessageType_MsgPsiAskReEnc           MessageType = 1
	MessageType_MsgPsiReEnc              MessageType = 2
	MessageType_MsgPsiIntersect          MessageType = 3
	MessageType_MsgTrainHup              MessageType = 4
	MessageType_MsgHomoPubkey            MessageType = 5
	MessageType_MsgTrainLoop             MessageType = 6
	MessageType_MsgTrainCalLocalGradCost MessageType = 7
	MessageType_MsgTrainPartBytes        MessageType = 8
	MessageType_MsgTrainCalEncGradCost   MessageType = 9
	MessageType_MsgTrainEncGradCost      MessageType = 10
	MessageType_MsgTrainDecLocalGradCost MessageType = 11
	MessageType_MsgTrainGradAndCost      MessageType = 12
	MessageType_MsgTrainUpdCostGrad      MessageType = 13
	MessageType_MsgTrainStatus           MessageType = 14
	MessageType_MsgTrainCheckStatus      MessageType = 15
	MessageType_MsgTrainModels           MessageType = 16
	MessageType_MsgCheckPauseRound       MessageType = 17
	MessageType_MsgTrainSet              MessageType = 18
	MessageType_MsgContinueLoop          MessageType = 19
	MessageType_MsgPredictHup            MessageType = 51
	MessageType_MsgPredictPart           MessageType = 52
	MessageType_MsgPredictFinal          MessageType = 53
)

var MessageType_name = map[int32]string{
	0:  "MsgPsiEnc",
	1:  "MsgPsiAskReEnc",
	2:  "MsgPsiReEnc",
	3:  "MsgPsiIntersect",
	4:  "MsgTrainHup",
	5:  "MsgHomoPubkey",
	6:  "MsgTrainLoop",
	7:  "MsgTrainCalLocalGradCost",
	8:  "MsgTrainPartBytes",
	9:  "MsgTrainCalEncGradCost",
	10: "MsgTrainEncGradCost",
	11: "MsgTrainDecLocalGradCost",
	12: "MsgTrainGradAndCost",
	13: "MsgTrainUpdCostGrad",
	14: "MsgTrainStatus",
	15: "MsgTrainCheckStatus",
	16: "MsgTrainModels",
	17: "MsgCheckPauseRound",
	18: "MsgTrainSet",
	19: "MsgContinueLoop",
	51: "MsgPredictHup",
	52: "MsgPredictPart",
	53: "MsgPredictFinal",
}

var MessageType_value = map[string]int32{
	"MsgPsiEnc":                0,
	"MsgPsiAskReEnc":           1,
	"MsgPsiReEnc":              2,
	"MsgPsiIntersect":          3,
	"MsgTrainHup":              4,
	"MsgHomoPubkey":            5,
	"MsgTrainLoop":             6,
	"MsgTrainCalLocalGradCost": 7,
	"MsgTrainPartBytes":        8,
	"MsgTrainCalEncGradCost":   9,
	"MsgTrainEncGradCost":      10,
	"MsgTrainDecLocalGradCost": 11,
	"MsgTrainGradAndCost":      12,
	"MsgTrainUpdCostGrad":      13,
	"MsgTrainStatus":           14,
	"MsgTrainCheckStatus":      15,
	"MsgTrainModels":           16,
	"MsgCheckPauseRound":       17,
	"MsgTrainSet":              18,
	"MsgContinueLoop":          19,
	"MsgPredictHup":            51,
	"MsgPredictPart":           52,
	"MsgPredictFinal":          53,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cba41b5f67b9a4c9, []int{0}
}

type Message struct {
	Type                 MessageType                       `protobuf:"varint,1,opt,name=type,proto3,enum=logic_reg_vl.MessageType" json:"type,omitempty"`
	To                   string                            `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	From                 string                            `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	LoopRound            uint64                            `protobuf:"varint,4,opt,name=loopRound,proto3" json:"loopRound,omitempty"`
	VlLPsiReEncIDsReq    *mpc.VLPsiReEncIDsRequest         `protobuf:"bytes,5,opt,name=vlLPsiReEncIDsReq,proto3" json:"vlLPsiReEncIDsReq,omitempty"`
	VlLPsiReEncIDsResp   *mpc.VLPsiReEncIDsResponse        `protobuf:"bytes,6,opt,name=vlLPsiReEncIDsResp,proto3" json:"vlLPsiReEncIDsResp,omitempty"`
	HomoPubkey           []byte                            `protobuf:"bytes,7,opt,name=homoPubkey,proto3" json:"homoPubkey,omitempty"`
	PartBytes            []byte                            `protobuf:"bytes,8,opt,name=PartBytes,proto3" json:"PartBytes,omitempty"`
	EncGradFromOther     []byte                            `protobuf:"bytes,9,opt,name=encGradFromOther,proto3" json:"encGradFromOther,omitempty"`
	EncCostFromOther     []byte                            `protobuf:"bytes,10,opt,name=encCostFromOther,proto3" json:"encCostFromOther,omitempty"`
	GradBytes            []byte                            `protobuf:"bytes,11,opt,name=gradBytes,proto3" json:"gradBytes,omitempty"`
	CostBytes            []byte                            `protobuf:"bytes,12,opt,name=costBytes,proto3" json:"costBytes,omitempty"`
	Stopped              bool                              `protobuf:"varint,13,opt,name=stopped,proto3" json:"stopped,omitempty"`
	TrainSet             []*common.TrainTaskResult_FileRow `protobuf:"bytes,14,rep,name=trainSet,proto3" json:"trainSet,omitempty"`
	PauseRound           uint64                            `protobuf:"varint,15,opt,name=pauseRound,proto3" json:"pauseRound,omitempty"`
	TriggerRound         uint64                            `protobuf:"varint,16,opt,name=triggerRound,proto3" json:"triggerRound,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_cba41b5f67b9a4c9, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_MsgPsiEnc
}

func (m *Message) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetLoopRound() uint64 {
	if m != nil {
		return m.LoopRound
	}
	return 0
}

func (m *Message) GetVlLPsiReEncIDsReq() *mpc.VLPsiReEncIDsRequest {
	if m != nil {
		return m.VlLPsiReEncIDsReq
	}
	return nil
}

func (m *Message) GetVlLPsiReEncIDsResp() *mpc.VLPsiReEncIDsResponse {
	if m != nil {
		return m.VlLPsiReEncIDsResp
	}
	return nil
}

func (m *Message) GetHomoPubkey() []byte {
	if m != nil {
		return m.HomoPubkey
	}
	return nil
}

func (m *Message) GetPartBytes() []byte {
	if m != nil {
		return m.PartBytes
	}
	return nil
}

func (m *Message) GetEncGradFromOther() []byte {
	if m != nil {
		return m.EncGradFromOther
	}
	return nil
}

func (m *Message) GetEncCostFromOther() []byte {
	if m != nil {
		return m.EncCostFromOther
	}
	return nil
}

func (m *Message) GetGradBytes() []byte {
	if m != nil {
		return m.GradBytes
	}
	return nil
}

func (m *Message) GetCostBytes() []byte {
	if m != nil {
		return m.CostBytes
	}
	return nil
}

func (m *Message) GetStopped() bool {
	if m != nil {
		return m.Stopped
	}
	return false
}

func (m *Message) GetTrainSet() []*common.TrainTaskResult_FileRow {
	if m != nil {
		return m.TrainSet
	}
	return nil
}

func (m *Message) GetPauseRound() uint64 {
	if m != nil {
		return m.PauseRound
	}
	return 0
}

func (m *Message) GetTriggerRound() uint64 {
	if m != nil {
		return m.TriggerRound
	}
	return 0
}

type PredictMessage struct {
	Type                 MessageType                `protobuf:"varint,1,opt,name=type,proto3,enum=logic_reg_vl.MessageType" json:"type,omitempty"`
	To                   string                     `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	From                 string                     `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	VlLPsiReEncIDsReq    *mpc.VLPsiReEncIDsRequest  `protobuf:"bytes,4,opt,name=vlLPsiReEncIDsReq,proto3" json:"vlLPsiReEncIDsReq,omitempty"`
	VlLPsiReEncIDsResp   *mpc.VLPsiReEncIDsResponse `protobuf:"bytes,5,opt,name=vlLPsiReEncIDsResp,proto3" json:"vlLPsiReEncIDsResp,omitempty"`
	PredictPart          []float64                  `protobuf:"fixed64,6,rep,packed,name=predictPart,proto3" json:"predictPart,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PredictMessage) Reset()         { *m = PredictMessage{} }
func (m *PredictMessage) String() string { return proto.CompactTextString(m) }
func (*PredictMessage) ProtoMessage()    {}
func (*PredictMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cba41b5f67b9a4c9, []int{1}
}

func (m *PredictMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictMessage.Unmarshal(m, b)
}
func (m *PredictMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictMessage.Marshal(b, m, deterministic)
}
func (m *PredictMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictMessage.Merge(m, src)
}
func (m *PredictMessage) XXX_Size() int {
	return xxx_messageInfo_PredictMessage.Size(m)
}
func (m *PredictMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PredictMessage proto.InternalMessageInfo

func (m *PredictMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_MsgPsiEnc
}

func (m *PredictMessage) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *PredictMessage) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *PredictMessage) GetVlLPsiReEncIDsReq() *mpc.VLPsiReEncIDsRequest {
	if m != nil {
		return m.VlLPsiReEncIDsReq
	}
	return nil
}

func (m *PredictMessage) GetVlLPsiReEncIDsResp() *mpc.VLPsiReEncIDsResponse {
	if m != nil {
		return m.VlLPsiReEncIDsResp
	}
	return nil
}

func (m *PredictMessage) GetPredictPart() []float64 {
	if m != nil {
		return m.PredictPart
	}
	return nil
}

func init() {
	proto.RegisterEnum("logic_reg_vl.MessageType", MessageType_name, MessageType_value)
	proto.RegisterType((*Message)(nil), "logic_reg_vl.Message")
	proto.RegisterType((*PredictMessage)(nil), "logic_reg_vl.PredictMessage")
}

func init() {
	proto.RegisterFile("mpc/learners/logic_reg_vl/logic_reg_vl.proto", fileDescriptor_cba41b5f67b9a4c9)
}

var fileDescriptor_cba41b5f67b9a4c9 = []byte{
	// 716 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0x4f, 0x6f, 0xda, 0x4a,
	0x14, 0xc5, 0x9f, 0xc1, 0xe1, 0xcf, 0xf0, 0x6f, 0x18, 0xf4, 0xf2, 0x1c, 0x14, 0xbd, 0x67, 0x65,
	0x65, 0x45, 0xaf, 0x20, 0x25, 0xed, 0xaa, 0xab, 0x84, 0x84, 0x24, 0x55, 0x50, 0x91, 0x43, 0xab,
	0xaa, 0x9b, 0xc8, 0xb1, 0x6f, 0x8d, 0x15, 0xe3, 0x99, 0x7a, 0xc6, 0xa9, 0x58, 0xf6, 0x6b, 0xf4,
	0xb3, 0x76, 0x51, 0xcd, 0x8c, 0x01, 0x93, 0xa4, 0x9b, 0xaa, 0xdd, 0x00, 0xf3, 0x3b, 0xe7, 0xfa,
	0xe2, 0x7b, 0x8f, 0x01, 0xfd, 0xbf, 0x60, 0xfe, 0x30, 0x06, 0x2f, 0x4d, 0x20, 0xe5, 0xc3, 0x98,
	0x86, 0x91, 0x7f, 0x9b, 0x42, 0x78, 0xfb, 0x10, 0x6f, 0x1d, 0x06, 0x2c, 0xa5, 0x82, 0x92, 0x66,
	0x91, 0xf5, 0x5b, 0xb2, 0x96, 0xf1, 0x48, 0x8b, 0xfd, 0x9e, 0x4f, 0x17, 0x0b, 0x9a, 0x0c, 0xf5,
	0x9b, 0x86, 0x07, 0xdf, 0x4d, 0x54, 0x9d, 0x00, 0xe7, 0x5e, 0x08, 0xe4, 0x05, 0x32, 0xc5, 0x92,
	0x81, 0x65, 0xd8, 0x86, 0xd3, 0x3e, 0xda, 0x1b, 0x6c, 0x35, 0xc8, 0x4d, 0xb3, 0x25, 0x03, 0x57,
	0xd9, 0x48, 0x1b, 0x95, 0x04, 0xb5, 0x4a, 0xb6, 0xe1, 0xd4, 0xdd, 0x92, 0xa0, 0x84, 0x20, 0xf3,
	0x53, 0x4a, 0x17, 0x56, 0x59, 0x11, 0xf5, 0x99, 0xec, 0xa3, 0x7a, 0x4c, 0x29, 0x73, 0x69, 0x96,
	0x04, 0x96, 0x69, 0x1b, 0x8e, 0xe9, 0x6e, 0x00, 0xb9, 0x40, 0xdd, 0x87, 0xf8, 0x7a, 0xca, 0x23,
	0x17, 0xce, 0x13, 0xff, 0xea, 0x8c, 0xbb, 0xf0, 0xd9, 0xda, 0xb1, 0x0d, 0xa7, 0x71, 0xb4, 0x37,
	0x58, 0x30, 0x7f, 0xf0, 0xfe, 0x91, 0x98, 0x01, 0x17, 0xee, 0xd3, 0x1a, 0xf2, 0x06, 0x91, 0xc7,
	0x90, 0x33, 0xab, 0xa2, 0xae, 0xd4, 0x7f, 0xee, 0x4a, 0x9c, 0xd1, 0x84, 0x83, 0xfb, 0x4c, 0x15,
	0xf9, 0x17, 0xa1, 0x39, 0x5d, 0xd0, 0x69, 0x76, 0x77, 0x0f, 0x4b, 0xab, 0x6a, 0x1b, 0x4e, 0xd3,
	0x2d, 0x10, 0x79, 0x4b, 0x53, 0x2f, 0x15, 0xa7, 0x4b, 0x01, 0xdc, 0xaa, 0x29, 0x79, 0x03, 0xc8,
	0x21, 0xc2, 0x90, 0xf8, 0x17, 0xa9, 0x17, 0x8c, 0x53, 0xba, 0x78, 0x2b, 0xe6, 0x90, 0x5a, 0x75,
	0x65, 0x7a, 0xc2, 0x73, 0xef, 0x88, 0x72, 0xb1, 0xf1, 0xa2, 0xb5, 0x77, 0x8b, 0xcb, 0xae, 0x61,
	0xea, 0x05, 0xba, 0x6b, 0x43, 0x77, 0x5d, 0x03, 0xa9, 0xfa, 0x94, 0xe7, 0xdf, 0xa9, 0xa9, 0xd5,
	0x35, 0x20, 0x16, 0xaa, 0x72, 0x41, 0x19, 0x83, 0xc0, 0x6a, 0xd9, 0x86, 0x53, 0x73, 0x57, 0x47,
	0xf2, 0x1a, 0xd5, 0x44, 0xea, 0x45, 0xc9, 0x0d, 0x08, 0xab, 0x6d, 0x97, 0x9d, 0xc6, 0xd1, 0x7f,
	0x83, 0x3c, 0x1e, 0x33, 0xc9, 0x67, 0x1e, 0xbf, 0x77, 0x81, 0x67, 0xb1, 0x18, 0x8c, 0xa3, 0x18,
	0x5c, 0xfa, 0xc5, 0x5d, 0x17, 0xc8, 0x41, 0x31, 0x2f, 0xe3, 0xa0, 0x97, 0xdb, 0x51, 0xcb, 0x2d,
	0x10, 0x72, 0x80, 0x9a, 0x22, 0x8d, 0xc2, 0x10, 0x52, 0xed, 0xc0, 0xca, 0xb1, 0xc5, 0x0e, 0xbe,
	0x95, 0x50, 0x7b, 0x9a, 0x42, 0x10, 0xf9, 0xe2, 0x0f, 0xa6, 0xf0, 0xd9, 0x9c, 0x99, 0xbf, 0x2d,
	0x67, 0x3b, 0xbf, 0x94, 0x33, 0x1b, 0x35, 0x98, 0xbe, 0x73, 0x99, 0x1e, 0xab, 0x62, 0x97, 0x1d,
	0xc3, 0x2d, 0xa2, 0xc3, 0xaf, 0x26, 0x6a, 0x14, 0x6e, 0x98, 0xb4, 0x50, 0x7d, 0xc2, 0xc3, 0x29,
	0x8f, 0xce, 0x13, 0x1f, 0xff, 0x45, 0x08, 0x6a, 0xeb, 0xe3, 0x89, 0x5c, 0x92, 0x64, 0x06, 0xe9,
	0xa0, 0x86, 0x66, 0x1a, 0x94, 0x48, 0x0f, 0x75, 0x34, 0xb8, 0x4a, 0x04, 0xa4, 0x1c, 0x7c, 0x81,
	0xcb, 0xb9, 0x4b, 0x6d, 0xf8, 0x32, 0x63, 0xd8, 0x24, 0x5d, 0xd4, 0x9a, 0xf0, 0xf0, 0x72, 0x1d,
	0x72, 0xbc, 0x43, 0x30, 0x6a, 0xae, 0x3c, 0xd7, 0x94, 0x32, 0x5c, 0x21, 0xfb, 0xc8, 0x5a, 0x91,
	0x91, 0x17, 0x5f, 0x53, 0xdf, 0x8b, 0x65, 0x9e, 0x65, 0x4e, 0x71, 0x95, 0xfc, 0x8d, 0xba, 0x2b,
	0x75, 0xfd, 0x34, 0xe0, 0x1a, 0xe9, 0xa3, 0xdd, 0x42, 0xd1, 0xb9, 0x7e, 0x04, 0x54, 0x49, 0x9d,
	0xfc, 0x83, 0x7a, 0x2b, 0xad, 0x28, 0xa0, 0x62, 0xa7, 0x33, 0xf0, 0xb7, 0x3b, 0x35, 0x8a, 0x65,
	0x92, 0x9e, 0x24, 0x5a, 0x68, 0x16, 0x85, 0x77, 0x4c, 0x41, 0xa9, 0xe3, 0x56, 0x3e, 0x29, 0x25,
	0xdc, 0x08, 0x4f, 0x64, 0x1c, 0xb7, 0x8b, 0xe6, 0xd1, 0x1c, 0xfc, 0xfb, 0x5c, 0xe8, 0x14, 0xcd,
	0x13, 0x1a, 0x40, 0xcc, 0x31, 0x26, 0xbb, 0x88, 0x4c, 0x78, 0xa8, 0x7c, 0xd3, 0x75, 0xc0, 0x71,
	0xb7, 0x38, 0xc8, 0x1b, 0x10, 0x98, 0xe4, 0xe3, 0x1e, 0xd1, 0x44, 0x44, 0x49, 0x06, 0x6a, 0x70,
	0xbd, 0x7c, 0xba, 0x79, 0xcc, 0xe5, 0xc0, 0x8f, 0x57, 0xbb, 0xdb, 0x2c, 0x1b, 0xbf, 0x5c, 0xad,
	0x4a, 0xb3, 0x71, 0x94, 0x78, 0x31, 0x7e, 0x75, 0x7a, 0xf9, 0x71, 0x1c, 0x46, 0x62, 0x9e, 0xdd,
	0xc9, 0xe7, 0x72, 0x38, 0xf5, 0x82, 0x20, 0x06, 0xfd, 0x9a, 0x1f, 0xce, 0x66, 0x1f, 0x86, 0x81,
	0x17, 0x0d, 0xd5, 0xcf, 0x39, 0x1f, 0xfe, 0xf4, 0xef, 0xe2, 0xae, 0xa2, 0x1c, 0xc7, 0x3f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0x6a, 0x0a, 0xa0, 0x52, 0x06, 0x00, 0x00,
}