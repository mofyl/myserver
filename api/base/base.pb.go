// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: base.proto

package base

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Api int32

const (
	Api_Empty                                  Api = 0
	Api_BuildConn                              Api = 1
	Api_WebLogin                               Api = 2
	Api_AddEmployeeInfo                        Api = 3
	Api_UpdateEmployeeInfo                     Api = 4
	Api_GetEmployeeInfo                        Api = 5
	Api_AddEmployeeHealthInfo                  Api = 6
	Api_UpdateEmployeeHealthInfo               Api = 7
	Api_GetEmployeeHealthInfo                  Api = 8
	Api_AddEmployeeHealthRecord                Api = 9
	Api_AddCompanyInfo                         Api = 10
	Api_UpdateCompanyInfo                      Api = 11
	Api_GetCompanyInfo                         Api = 12
	Api_AddPreventionRecord                    Api = 13
	Api_UpdatePreventionRecord                 Api = 14
	Api_GetPreventionRecord                    Api = 15
	Api_GetQiniuToken                          Api = 16
	Api_GetProvinces                           Api = 17
	Api_GetCity                                Api = 18
	Api_GetDistrict                            Api = 19
	Api_GetStreet                              Api = 20
	Api_GetCommunity                           Api = 21
	Api_QueryEmployInfoByUser                  Api = 22
	Api_GetCompanyInfoList                     Api = 23
	Api_DeleteCompanyInfo                      Api = 24
	Api_GetPreventionRecordList                Api = 25
	Api_GetEmployeeHealthRecordList            Api = 26
	Api_DeleteEmployInfo                       Api = 27
	Api_GetEmployeeHealthRecordListByCompanyId Api = 28
	Api_ExportFinish                           Api = 29
	Api_Export                                 Api = 30
	Api_Heart                                  Api = 31
)

var Api_name = map[int32]string{
	0:  "Empty",
	1:  "BuildConn",
	2:  "WebLogin",
	3:  "AddEmployeeInfo",
	4:  "UpdateEmployeeInfo",
	5:  "GetEmployeeInfo",
	6:  "AddEmployeeHealthInfo",
	7:  "UpdateEmployeeHealthInfo",
	8:  "GetEmployeeHealthInfo",
	9:  "AddEmployeeHealthRecord",
	10: "AddCompanyInfo",
	11: "UpdateCompanyInfo",
	12: "GetCompanyInfo",
	13: "AddPreventionRecord",
	14: "UpdatePreventionRecord",
	15: "GetPreventionRecord",
	16: "GetQiniuToken",
	17: "GetProvinces",
	18: "GetCity",
	19: "GetDistrict",
	20: "GetStreet",
	21: "GetCommunity",
	22: "QueryEmployInfoByUser",
	23: "GetCompanyInfoList",
	24: "DeleteCompanyInfo",
	25: "GetPreventionRecordList",
	26: "GetEmployeeHealthRecordList",
	27: "DeleteEmployInfo",
	28: "GetEmployeeHealthRecordListByCompanyId",
	29: "ExportFinish",
	30: "Export",
	31: "Heart",
}

var Api_value = map[string]int32{
	"Empty":                                  0,
	"BuildConn":                              1,
	"WebLogin":                               2,
	"AddEmployeeInfo":                        3,
	"UpdateEmployeeInfo":                     4,
	"GetEmployeeInfo":                        5,
	"AddEmployeeHealthInfo":                  6,
	"UpdateEmployeeHealthInfo":               7,
	"GetEmployeeHealthInfo":                  8,
	"AddEmployeeHealthRecord":                9,
	"AddCompanyInfo":                         10,
	"UpdateCompanyInfo":                      11,
	"GetCompanyInfo":                         12,
	"AddPreventionRecord":                    13,
	"UpdatePreventionRecord":                 14,
	"GetPreventionRecord":                    15,
	"GetQiniuToken":                          16,
	"GetProvinces":                           17,
	"GetCity":                                18,
	"GetDistrict":                            19,
	"GetStreet":                              20,
	"GetCommunity":                           21,
	"QueryEmployInfoByUser":                  22,
	"GetCompanyInfoList":                     23,
	"DeleteCompanyInfo":                      24,
	"GetPreventionRecordList":                25,
	"GetEmployeeHealthRecordList":            26,
	"DeleteEmployInfo":                       27,
	"GetEmployeeHealthRecordListByCompanyId": 28,
	"ExportFinish":                           29,
	"Export":                                 30,
	"Heart":                                  31,
}

func (x Api) String() string {
	return proto.EnumName(Api_name, int32(x))
}

func (Api) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}

type Head struct {
	// int32 ApiID = 1;
	ConnID   string `protobuf:"bytes,1,opt,name=ConnID,proto3" json:"ConnID,omitempty"`
	GroupID  string `protobuf:"bytes,2,opt,name=GroupID,proto3" json:"GroupID,omitempty"`
	AppKey   string `protobuf:"bytes,3,opt,name=AppKey,proto3" json:"AppKey,omitempty"`
	ServerID string `protobuf:"bytes,4,opt,name=ServerID,proto3" json:"ServerID,omitempty"`
}

func (m *Head) Reset()         { *m = Head{} }
func (m *Head) String() string { return proto.CompactTextString(m) }
func (*Head) ProtoMessage()    {}
func (*Head) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}
func (m *Head) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Head) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Head.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Head) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Head.Merge(m, src)
}
func (m *Head) XXX_Size() int {
	return m.Size()
}
func (m *Head) XXX_DiscardUnknown() {
	xxx_messageInfo_Head.DiscardUnknown(m)
}

var xxx_messageInfo_Head proto.InternalMessageInfo

func (m *Head) GetConnID() string {
	if m != nil {
		return m.ConnID
	}
	return ""
}

func (m *Head) GetGroupID() string {
	if m != nil {
		return m.GroupID
	}
	return ""
}

func (m *Head) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *Head) GetServerID() string {
	if m != nil {
		return m.ServerID
	}
	return ""
}

type BaseRes struct {
	ErrorCode int64 `protobuf:"varint,1,opt,name=ErrorCode,proto3" json:"ErrorCode,omitempty"`
}

func (m *BaseRes) Reset()         { *m = BaseRes{} }
func (m *BaseRes) String() string { return proto.CompactTextString(m) }
func (*BaseRes) ProtoMessage()    {}
func (*BaseRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{1}
}
func (m *BaseRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseRes.Merge(m, src)
}
func (m *BaseRes) XXX_Size() int {
	return m.Size()
}
func (m *BaseRes) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseRes.DiscardUnknown(m)
}

var xxx_messageInfo_BaseRes proto.InternalMessageInfo

func (m *BaseRes) GetErrorCode() int64 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type Msg struct {
	ApiId Api    `protobuf:"varint,1,opt,name=api_id,json=apiId,proto3,enum=base.Api" json:"api_id,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{2}
}
func (m *Msg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return m.Size()
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetApiId() Api {
	if m != nil {
		return m.ApiId
	}
	return Api_Empty
}

func (m *Msg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type HeartReq struct {
	Head       *Head `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	UpdateTime int64 `protobuf:"varint,2,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (m *HeartReq) Reset()         { *m = HeartReq{} }
func (m *HeartReq) String() string { return proto.CompactTextString(m) }
func (*HeartReq) ProtoMessage()    {}
func (*HeartReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{3}
}
func (m *HeartReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeartReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeartReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeartReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartReq.Merge(m, src)
}
func (m *HeartReq) XXX_Size() int {
	return m.Size()
}
func (m *HeartReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartReq.DiscardUnknown(m)
}

var xxx_messageInfo_HeartReq proto.InternalMessageInfo

func (m *HeartReq) GetHead() *Head {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *HeartReq) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type HeartRes struct {
	Base *BaseRes `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (m *HeartRes) Reset()         { *m = HeartRes{} }
func (m *HeartRes) String() string { return proto.CompactTextString(m) }
func (*HeartRes) ProtoMessage()    {}
func (*HeartRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{4}
}
func (m *HeartRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeartRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeartRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeartRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartRes.Merge(m, src)
}
func (m *HeartRes) XXX_Size() int {
	return m.Size()
}
func (m *HeartRes) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartRes.DiscardUnknown(m)
}

var xxx_messageInfo_HeartRes proto.InternalMessageInfo

func (m *HeartRes) GetBase() *BaseRes {
	if m != nil {
		return m.Base
	}
	return nil
}

func init() {
	proto.RegisterEnum("base.Api", Api_name, Api_value)
	proto.RegisterType((*Head)(nil), "base.Head")
	proto.RegisterType((*BaseRes)(nil), "base.BaseRes")
	proto.RegisterType((*Msg)(nil), "base.Msg")
	proto.RegisterType((*HeartReq)(nil), "base.HeartReq")
	proto.RegisterType((*HeartRes)(nil), "base.HeartRes")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150) }

var fileDescriptor_db1b6b0986796150 = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4d, 0x4f, 0xdb, 0x40,
	0x10, 0x4d, 0xc8, 0xf7, 0x24, 0x81, 0x65, 0x80, 0x60, 0x3e, 0x6a, 0x68, 0x0e, 0x6d, 0x85, 0x54,
	0x0e, 0xf4, 0xd8, 0x53, 0xbe, 0x1a, 0x22, 0xa8, 0x54, 0x0c, 0xa8, 0x47, 0x64, 0xd8, 0x29, 0xac,
	0x1a, 0x7b, 0xdd, 0xf5, 0x06, 0xd5, 0xff, 0xa2, 0x3f, 0xaa, 0x87, 0x1e, 0x39, 0xf6, 0x58, 0xc1,
	0x1f, 0xa9, 0x76, 0x1d, 0xc0, 0x51, 0xaa, 0xde, 0x3c, 0xef, 0xcd, 0x7b, 0x3b, 0xf3, 0x76, 0x65,
	0x80, 0x4b, 0x3f, 0xa6, 0xfd, 0x48, 0x49, 0x2d, 0xb1, 0x68, 0xbe, 0xdb, 0x63, 0x28, 0x1e, 0x92,
	0xcf, 0xb1, 0x05, 0xe5, 0x9e, 0x0c, 0xc3, 0x51, 0xdf, 0xc9, 0xef, 0xe6, 0xdf, 0xd4, 0xbc, 0x69,
	0x85, 0x0e, 0x54, 0x86, 0x4a, 0x4e, 0xa2, 0x51, 0xdf, 0x59, 0xb0, 0xc4, 0x63, 0x69, 0x14, 0x9d,
	0x28, 0x3a, 0xa2, 0xc4, 0x29, 0xa4, 0x8a, 0xb4, 0xc2, 0x4d, 0xa8, 0x9e, 0x92, 0xba, 0x25, 0x35,
	0xea, 0x3b, 0x45, 0xcb, 0x3c, 0xd5, 0xed, 0xd7, 0x50, 0xe9, 0xfa, 0x31, 0x79, 0x14, 0xe3, 0x36,
	0xd4, 0x06, 0x4a, 0x49, 0xd5, 0x93, 0x9c, 0xec, 0x99, 0x05, 0xef, 0x19, 0x68, 0xbf, 0x87, 0xc2,
	0xc7, 0xf8, 0x1a, 0x77, 0xa1, 0xec, 0x47, 0xe2, 0x42, 0x70, 0xdb, 0xb1, 0x78, 0x50, 0xdb, 0xb7,
	0x0b, 0x74, 0x22, 0xe1, 0x95, 0xfc, 0x48, 0x8c, 0x38, 0x22, 0x14, 0xfb, 0xbe, 0xf6, 0xed, 0x70,
	0x0d, 0xcf, 0x7e, 0xb7, 0x8f, 0xa0, 0x7a, 0x48, 0xbe, 0xd2, 0x1e, 0x7d, 0x43, 0x17, 0x8a, 0x37,
	0xe4, 0xa7, 0xfa, 0xfa, 0x01, 0xa4, 0x7a, 0xb3, 0xb1, 0x67, 0x71, 0xdc, 0x81, 0xfa, 0x24, 0xe2,
	0xbe, 0xa6, 0x0b, 0x2d, 0x02, 0xb2, 0x36, 0x05, 0x0f, 0x52, 0xe8, 0x4c, 0x04, 0xd4, 0x7e, 0xfb,
	0x64, 0x16, 0xe3, 0x4b, 0xb0, 0xa1, 0x4d, 0xcd, 0x9a, 0xa9, 0xd9, 0x74, 0x21, 0xcf, 0x52, 0x7b,
	0x3f, 0x4b, 0x50, 0xe8, 0x44, 0x02, 0x6b, 0x50, 0x1a, 0x04, 0x91, 0x4e, 0x58, 0x0e, 0x9b, 0x50,
	0xeb, 0x4e, 0xc4, 0x98, 0x9b, 0x44, 0x59, 0x1e, 0x1b, 0x50, 0xfd, 0x4c, 0x97, 0xc7, 0xf2, 0x5a,
	0x84, 0x6c, 0x01, 0x57, 0x60, 0xa9, 0xc3, 0xf9, 0x20, 0x88, 0xc6, 0x32, 0x21, 0x1a, 0x85, 0x5f,
	0x24, 0x2b, 0x60, 0x0b, 0xf0, 0xdc, 0x4e, 0x30, 0x83, 0x17, 0x4d, 0xf3, 0x90, 0xf4, 0x0c, 0x58,
	0xc2, 0x0d, 0x58, 0xcb, 0x38, 0x1c, 0x92, 0x3f, 0xd6, 0x37, 0x96, 0x2a, 0xe3, 0x36, 0x38, 0xb3,
	0x3e, 0x19, 0xb6, 0x62, 0x84, 0x19, 0xb7, 0x0c, 0x55, 0xc5, 0x2d, 0x58, 0x9f, 0xf3, 0xf4, 0xe8,
	0x4a, 0x2a, 0xce, 0x6a, 0x88, 0xb0, 0xd8, 0xe1, 0xbc, 0x27, 0x83, 0xc8, 0x0f, 0x13, 0x2b, 0x00,
	0x5c, 0x83, 0xe5, 0xf4, 0xa4, 0x2c, 0x5c, 0x37, 0xad, 0x43, 0xd2, 0x59, 0xac, 0x81, 0xeb, 0xb0,
	0xd2, 0xe1, 0xfc, 0x93, 0xa2, 0x5b, 0x0a, 0xb5, 0x90, 0xe1, 0xd4, 0xb7, 0x89, 0x9b, 0xd0, 0x4a,
	0x3d, 0xe6, 0xb8, 0x45, 0x23, 0x1a, 0x92, 0x9e, 0x23, 0x96, 0x70, 0x19, 0x9a, 0x43, 0xd2, 0x27,
	0x22, 0x14, 0x93, 0x33, 0xf9, 0x95, 0x42, 0xc6, 0x90, 0x41, 0xc3, 0xf6, 0xca, 0x5b, 0x11, 0x5e,
	0x51, 0xcc, 0x96, 0xb1, 0x0e, 0x15, 0x33, 0x86, 0xd0, 0x09, 0x43, 0x5c, 0x82, 0xfa, 0x90, 0x74,
	0x5f, 0xc4, 0x5a, 0x89, 0x2b, 0xcd, 0x56, 0xcc, 0xfd, 0x0c, 0x49, 0x9f, 0x6a, 0x45, 0xa4, 0xd9,
	0xea, 0x54, 0xde, 0x93, 0x41, 0x30, 0x09, 0x8d, 0x62, 0xcd, 0x04, 0x75, 0x32, 0x21, 0x95, 0xa4,
	0x79, 0x98, 0x35, 0xba, 0xc9, 0x79, 0x4c, 0x8a, 0xb5, 0xcc, 0x4d, 0xcd, 0x2e, 0x78, 0x2c, 0x62,
	0xcd, 0xd6, 0x4d, 0x1e, 0x7d, 0x1a, 0xd3, 0x6c, 0x1e, 0x8e, 0xc9, 0xf5, 0x1f, 0x6b, 0x58, 0xcd,
	0x06, 0xee, 0xc0, 0xd6, 0xdc, 0x7d, 0x64, 0x1a, 0x36, 0x71, 0x15, 0x58, 0x6a, 0xfa, 0x3c, 0x08,
	0xdb, 0xc2, 0x3d, 0x78, 0xf5, 0x1f, 0x59, 0x37, 0x79, 0x9c, 0x80, 0xb3, 0x6d, 0xb3, 0xdb, 0xe0,
	0x7b, 0x24, 0x95, 0xfe, 0x20, 0x42, 0x11, 0xdf, 0xb0, 0x17, 0x08, 0x50, 0x4e, 0x11, 0xe6, 0x9a,
	0x37, 0x6b, 0x9f, 0x3a, 0xdb, 0xe9, 0x3a, 0xbf, 0xee, 0xdd, 0xfc, 0xdd, 0xbd, 0x9b, 0xff, 0x73,
	0xef, 0xe6, 0x7f, 0x3c, 0xb8, 0xb9, 0xbb, 0x07, 0x37, 0xf7, 0xfb, 0xc1, 0xcd, 0x5d, 0x96, 0xed,
	0xdf, 0xe3, 0xdd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0x62, 0xe3, 0x67, 0x4b, 0x04, 0x00,
	0x00,
}

func (m *Head) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Head) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Head) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ServerID) > 0 {
		i -= len(m.ServerID)
		copy(dAtA[i:], m.ServerID)
		i = encodeVarintBase(dAtA, i, uint64(len(m.ServerID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AppKey) > 0 {
		i -= len(m.AppKey)
		copy(dAtA[i:], m.AppKey)
		i = encodeVarintBase(dAtA, i, uint64(len(m.AppKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.GroupID) > 0 {
		i -= len(m.GroupID)
		copy(dAtA[i:], m.GroupID)
		i = encodeVarintBase(dAtA, i, uint64(len(m.GroupID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ConnID) > 0 {
		i -= len(m.ConnID)
		copy(dAtA[i:], m.ConnID)
		i = encodeVarintBase(dAtA, i, uint64(len(m.ConnID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BaseRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ErrorCode != 0 {
		i = encodeVarintBase(dAtA, i, uint64(m.ErrorCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Msg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Msg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Msg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintBase(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.ApiId != 0 {
		i = encodeVarintBase(dAtA, i, uint64(m.ApiId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *HeartReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeartReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeartReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdateTime != 0 {
		i = encodeVarintBase(dAtA, i, uint64(m.UpdateTime))
		i--
		dAtA[i] = 0x10
	}
	if m.Head != nil {
		{
			size, err := m.Head.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBase(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HeartRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeartRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeartRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Base != nil {
		{
			size, err := m.Base.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBase(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBase(dAtA []byte, offset int, v uint64) int {
	offset -= sovBase(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Head) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConnID)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	l = len(m.GroupID)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	l = len(m.AppKey)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	l = len(m.ServerID)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	return n
}

func (m *BaseRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ErrorCode != 0 {
		n += 1 + sovBase(uint64(m.ErrorCode))
	}
	return n
}

func (m *Msg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ApiId != 0 {
		n += 1 + sovBase(uint64(m.ApiId))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	return n
}

func (m *HeartReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Head != nil {
		l = m.Head.Size()
		n += 1 + l + sovBase(uint64(l))
	}
	if m.UpdateTime != 0 {
		n += 1 + sovBase(uint64(m.UpdateTime))
	}
	return n
}

func (m *HeartRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Base != nil {
		l = m.Base.Size()
		n += 1 + l + sovBase(uint64(l))
	}
	return n
}

func sovBase(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBase(x uint64) (n int) {
	return sovBase(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Head) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Head: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Head: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBase
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBase
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BaseRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BaseRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorCode", wireType)
			}
			m.ErrorCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrorCode |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBase
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBase
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Msg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Msg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Msg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiId", wireType)
			}
			m.ApiId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApiId |= Api(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBase
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBase
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HeartReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HeartReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeartReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Head", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Head == nil {
				m.Head = &Head{}
			}
			if err := m.Head.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateTime", wireType)
			}
			m.UpdateTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBase
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBase
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HeartRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HeartRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeartRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Base == nil {
				m.Base = &BaseRes{}
			}
			if err := m.Base.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBase
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBase
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBase(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBase
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBase
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBase
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthBase
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBase
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBase
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBase        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBase          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBase = fmt.Errorf("proto: unexpected end of group")
)
