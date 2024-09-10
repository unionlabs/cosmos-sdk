// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accounts/testing/counter/v1/counter.proto

package v1

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// MsgInit defines a message which initializes the counter with a given amount.
type MsgInit struct {
	// initial_value is the initial amount to set the counter to.
	InitialValue uint64 `protobuf:"varint,1,opt,name=initial_value,json=initialValue,proto3" json:"initial_value,omitempty"`
}

func (m *MsgInit) Reset()         { *m = MsgInit{} }
func (m *MsgInit) String() string { return proto.CompactTextString(m) }
func (*MsgInit) ProtoMessage()    {}
func (*MsgInit) Descriptor() ([]byte, []int) {
	return fileDescriptor_21c9320877186411, []int{0}
}
func (m *MsgInit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInit.Merge(m, src)
}
func (m *MsgInit) XXX_Size() int {
	return m.Size()
}
func (m *MsgInit) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInit.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInit proto.InternalMessageInfo

func (m *MsgInit) GetInitialValue() uint64 {
	if m != nil {
		return m.InitialValue
	}
	return 0
}

// MsgInitResponse defines the MsgInit response type.
type MsgInitResponse struct {
}

func (m *MsgInitResponse) Reset()         { *m = MsgInitResponse{} }
func (m *MsgInitResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInitResponse) ProtoMessage()    {}
func (*MsgInitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21c9320877186411, []int{1}
}
func (m *MsgInitResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitResponse.Merge(m, src)
}
func (m *MsgInitResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitResponse proto.InternalMessageInfo

// MsgIncreaseCounter defines a message which increases the counter by a given amount.
type MsgIncreaseCounter struct {
	// amount is the amount to increase the counter by.
	Amount uint64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *MsgIncreaseCounter) Reset()         { *m = MsgIncreaseCounter{} }
func (m *MsgIncreaseCounter) String() string { return proto.CompactTextString(m) }
func (*MsgIncreaseCounter) ProtoMessage()    {}
func (*MsgIncreaseCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_21c9320877186411, []int{2}
}
func (m *MsgIncreaseCounter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIncreaseCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIncreaseCounter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIncreaseCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIncreaseCounter.Merge(m, src)
}
func (m *MsgIncreaseCounter) XXX_Size() int {
	return m.Size()
}
func (m *MsgIncreaseCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIncreaseCounter.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIncreaseCounter proto.InternalMessageInfo

func (m *MsgIncreaseCounter) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// MsgIncreaseCounterResponse defines the MsgIncreaseCounter response type.
// Returns the new counter value.
type MsgIncreaseCounterResponse struct {
	// new_amount defines the new counter value after the increase.
	NewAmount uint64 `protobuf:"varint,1,opt,name=new_amount,json=newAmount,proto3" json:"new_amount,omitempty"`
}

func (m *MsgIncreaseCounterResponse) Reset()         { *m = MsgIncreaseCounterResponse{} }
func (m *MsgIncreaseCounterResponse) String() string { return proto.CompactTextString(m) }
func (*MsgIncreaseCounterResponse) ProtoMessage()    {}
func (*MsgIncreaseCounterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21c9320877186411, []int{3}
}
func (m *MsgIncreaseCounterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIncreaseCounterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIncreaseCounterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIncreaseCounterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIncreaseCounterResponse.Merge(m, src)
}
func (m *MsgIncreaseCounterResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgIncreaseCounterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIncreaseCounterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIncreaseCounterResponse proto.InternalMessageInfo

func (m *MsgIncreaseCounterResponse) GetNewAmount() uint64 {
	if m != nil {
		return m.NewAmount
	}
	return 0
}

// MsgTestDependencies is used to test the dependencies.
type MsgTestDependencies struct {
}

func (m *MsgTestDependencies) Reset()         { *m = MsgTestDependencies{} }
func (m *MsgTestDependencies) String() string { return proto.CompactTextString(m) }
func (*MsgTestDependencies) ProtoMessage()    {}
func (*MsgTestDependencies) Descriptor() ([]byte, []int) {
	return fileDescriptor_21c9320877186411, []int{4}
}
func (m *MsgTestDependencies) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTestDependencies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTestDependencies.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTestDependencies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTestDependencies.Merge(m, src)
}
func (m *MsgTestDependencies) XXX_Size() int {
	return m.Size()
}
func (m *MsgTestDependencies) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTestDependencies.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTestDependencies proto.InternalMessageInfo

// MsgTestDependenciesResponse is used to test the dependencies.
type MsgTestDependenciesResponse struct {
	// chain_id is used to test that the header service correctly works.
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// address is used to test address codec.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// before_gas is used to test the gas meter reporting.
	BeforeGas uint64 `protobuf:"varint,3,opt,name=before_gas,json=beforeGas,proto3" json:"before_gas,omitempty"`
	// after_gas is used to test gas meter increasing.
	AfterGas uint64 `protobuf:"varint,4,opt,name=after_gas,json=afterGas,proto3" json:"after_gas,omitempty"`
	// funds reports the funds from the implementation.Funds method.
	Funds github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=funds,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"funds"`
}

func (m *MsgTestDependenciesResponse) Reset()         { *m = MsgTestDependenciesResponse{} }
func (m *MsgTestDependenciesResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTestDependenciesResponse) ProtoMessage()    {}
func (*MsgTestDependenciesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21c9320877186411, []int{5}
}
func (m *MsgTestDependenciesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTestDependenciesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTestDependenciesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTestDependenciesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTestDependenciesResponse.Merge(m, src)
}
func (m *MsgTestDependenciesResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgTestDependenciesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTestDependenciesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTestDependenciesResponse proto.InternalMessageInfo

func (m *MsgTestDependenciesResponse) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *MsgTestDependenciesResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgTestDependenciesResponse) GetBeforeGas() uint64 {
	if m != nil {
		return m.BeforeGas
	}
	return 0
}

func (m *MsgTestDependenciesResponse) GetAfterGas() uint64 {
	if m != nil {
		return m.AfterGas
	}
	return 0
}

func (m *MsgTestDependenciesResponse) GetFunds() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Funds
	}
	return nil
}

// QueryCounterRequest is used to query the counter value.
type QueryCounterRequest struct {
}

func (m *QueryCounterRequest) Reset()         { *m = QueryCounterRequest{} }
func (m *QueryCounterRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCounterRequest) ProtoMessage()    {}
func (*QueryCounterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21c9320877186411, []int{6}
}
func (m *QueryCounterRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCounterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCounterRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCounterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCounterRequest.Merge(m, src)
}
func (m *QueryCounterRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCounterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCounterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCounterRequest proto.InternalMessageInfo

// QueryCounterResponse returns the counter value.
type QueryCounterResponse struct {
	// value defines the value of the counter.
	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *QueryCounterResponse) Reset()         { *m = QueryCounterResponse{} }
func (m *QueryCounterResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCounterResponse) ProtoMessage()    {}
func (*QueryCounterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21c9320877186411, []int{7}
}
func (m *QueryCounterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCounterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCounterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCounterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCounterResponse.Merge(m, src)
}
func (m *QueryCounterResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCounterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCounterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCounterResponse proto.InternalMessageInfo

func (m *QueryCounterResponse) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgInit)(nil), "cosmos.accounts.testing.counter.v1.MsgInit")
	proto.RegisterType((*MsgInitResponse)(nil), "cosmos.accounts.testing.counter.v1.MsgInitResponse")
	proto.RegisterType((*MsgIncreaseCounter)(nil), "cosmos.accounts.testing.counter.v1.MsgIncreaseCounter")
	proto.RegisterType((*MsgIncreaseCounterResponse)(nil), "cosmos.accounts.testing.counter.v1.MsgIncreaseCounterResponse")
	proto.RegisterType((*MsgTestDependencies)(nil), "cosmos.accounts.testing.counter.v1.MsgTestDependencies")
	proto.RegisterType((*MsgTestDependenciesResponse)(nil), "cosmos.accounts.testing.counter.v1.MsgTestDependenciesResponse")
	proto.RegisterType((*QueryCounterRequest)(nil), "cosmos.accounts.testing.counter.v1.QueryCounterRequest")
	proto.RegisterType((*QueryCounterResponse)(nil), "cosmos.accounts.testing.counter.v1.QueryCounterResponse")
}

func init() {
	proto.RegisterFile("cosmos/accounts/testing/counter/v1/counter.proto", fileDescriptor_21c9320877186411)
}

var fileDescriptor_21c9320877186411 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x69, 0xd3, 0x34, 0x03, 0x08, 0x61, 0x02, 0x72, 0x53, 0xe1, 0x56, 0xe6, 0x12, 0xa1,
	0xb2, 0xdb, 0xc0, 0x91, 0x13, 0x6d, 0x25, 0xd4, 0x43, 0x0f, 0x44, 0x88, 0x03, 0x97, 0x68, 0x6d,
	0x4f, 0xdc, 0x55, 0x9b, 0xdd, 0xe0, 0x59, 0xa7, 0xf4, 0x2d, 0x78, 0x0e, 0x9e, 0xa4, 0xc7, 0x1e,
	0x39, 0x01, 0x4a, 0x9e, 0x80, 0x37, 0x40, 0xfb, 0x93, 0x8a, 0x0a, 0xc4, 0xc9, 0xf3, 0x7d, 0x33,
	0xdf, 0xe7, 0xd1, 0x7c, 0x0b, 0xfb, 0x85, 0xa6, 0xa9, 0x26, 0x2e, 0x8a, 0x42, 0x37, 0xca, 0x10,
	0x37, 0x48, 0x46, 0xaa, 0x8a, 0x3b, 0x88, 0x35, 0x9f, 0x0f, 0x57, 0x25, 0x9b, 0xd5, 0xda, 0xe8,
	0x38, 0xf3, 0x0a, 0xb6, 0x52, 0xb0, 0xa0, 0x60, 0xab, 0xb1, 0xf9, 0xb0, 0x9f, 0x06, 0xd7, 0x5c,
	0x10, 0xf2, 0xf9, 0x30, 0x47, 0x23, 0xac, 0x8d, 0x54, 0xde, 0xa3, 0xdf, 0xab, 0x74, 0xa5, 0x5d,
	0xc9, 0x6d, 0xe5, 0xd9, 0x8c, 0x41, 0xe7, 0x84, 0xaa, 0x63, 0x25, 0x4d, 0xfc, 0x0c, 0xee, 0x4b,
	0x25, 0x8d, 0x14, 0xe7, 0xe3, 0xb9, 0x38, 0x6f, 0x30, 0x89, 0x76, 0xa3, 0xc1, 0xfa, 0xe8, 0x5e,
	0x20, 0x3f, 0x58, 0x2e, 0x7b, 0x08, 0x0f, 0xc2, 0xfc, 0x08, 0x69, 0xa6, 0x15, 0x61, 0xb6, 0x07,
	0xb1, 0xa3, 0x8a, 0x1a, 0x05, 0xe1, 0xa1, 0xdf, 0x28, 0x7e, 0x02, 0x1b, 0x62, 0x6a, 0xeb, 0x60,
	0x13, 0x50, 0xf6, 0x1a, 0xfa, 0x7f, 0x4f, 0xaf, 0xbc, 0xe2, 0xa7, 0x00, 0x0a, 0x2f, 0xc6, 0xb7,
	0x94, 0x5d, 0x85, 0x17, 0x6f, 0xbc, 0xf8, 0x31, 0x3c, 0x3a, 0xa1, 0xea, 0x3d, 0x92, 0x39, 0xc2,
	0x19, 0xaa, 0x12, 0x55, 0x21, 0x91, 0xb2, 0x5f, 0x11, 0x6c, 0xff, 0x83, 0xbf, 0x71, 0xdd, 0x82,
	0xcd, 0xe2, 0x54, 0x48, 0x35, 0x96, 0xa5, 0xf3, 0xec, 0x8e, 0x3a, 0x0e, 0x1f, 0x97, 0x71, 0x02,
	0x1d, 0x51, 0x96, 0x35, 0x12, 0x25, 0x77, 0x7c, 0x27, 0x40, 0xbb, 0x4a, 0x8e, 0x13, 0x5d, 0xe3,
	0xb8, 0x12, 0x94, 0xac, 0xf9, 0x55, 0x3c, 0xf3, 0x56, 0x50, 0xbc, 0x0d, 0x5d, 0x31, 0x31, 0x58,
	0xbb, 0xee, 0xba, 0xeb, 0x6e, 0x3a, 0xc2, 0x36, 0x05, 0xb4, 0x27, 0x8d, 0x2a, 0x29, 0x69, 0xef,
	0xae, 0x0d, 0xee, 0xbe, 0xdc, 0x62, 0x21, 0x3f, 0x9b, 0x0d, 0x0b, 0xd9, 0xb0, 0x43, 0x2d, 0xd5,
	0xc1, 0xfe, 0xd5, 0xf7, 0x9d, 0xd6, 0xd7, 0x1f, 0x3b, 0x83, 0x4a, 0x9a, 0xd3, 0x26, 0x67, 0x85,
	0x9e, 0xf2, 0x10, 0xa4, 0xff, 0xbc, 0xa0, 0xf2, 0x8c, 0x9b, 0xcb, 0x19, 0x92, 0x13, 0xd0, 0xc8,
	0x3b, 0xdb, 0x53, 0xbc, 0x6b, 0xb0, 0xbe, 0xbc, 0xb9, 0xe0, 0xa7, 0x06, 0xc9, 0x64, 0x7b, 0xd0,
	0xbb, 0x4d, 0x87, 0x13, 0xf4, 0xa0, 0xfd, 0x67, 0xa8, 0x1e, 0x1c, 0x1c, 0x5d, 0x2d, 0xd2, 0xe8,
	0x7a, 0x91, 0x46, 0x3f, 0x17, 0x69, 0xf4, 0x65, 0x99, 0xb6, 0xae, 0x97, 0x69, 0xeb, 0xdb, 0x32,
	0x6d, 0x7d, 0x7c, 0xee, 0xff, 0x4e, 0xe5, 0x19, 0x93, 0x9a, 0x7f, 0xfe, 0xdf, 0x5b, 0xcd, 0x37,
	0xdc, 0x53, 0x7a, 0xf5, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x99, 0xcd, 0x49, 0xa2, 0xd8, 0x02, 0x00,
	0x00,
}

func (m *MsgInit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InitialValue != 0 {
		i = encodeVarintCounter(dAtA, i, uint64(m.InitialValue))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgInitResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgIncreaseCounter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIncreaseCounter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIncreaseCounter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintCounter(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgIncreaseCounterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIncreaseCounterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIncreaseCounterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewAmount != 0 {
		i = encodeVarintCounter(dAtA, i, uint64(m.NewAmount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgTestDependencies) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTestDependencies) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTestDependencies) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgTestDependenciesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTestDependenciesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTestDependenciesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Funds) > 0 {
		for iNdEx := len(m.Funds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Funds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCounter(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.AfterGas != 0 {
		i = encodeVarintCounter(dAtA, i, uint64(m.AfterGas))
		i--
		dAtA[i] = 0x20
	}
	if m.BeforeGas != 0 {
		i = encodeVarintCounter(dAtA, i, uint64(m.BeforeGas))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCounter(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintCounter(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCounterRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCounterRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCounterRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryCounterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCounterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCounterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintCounter(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCounter(dAtA []byte, offset int, v uint64) int {
	offset -= sovCounter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgInit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InitialValue != 0 {
		n += 1 + sovCounter(uint64(m.InitialValue))
	}
	return n
}

func (m *MsgInitResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgIncreaseCounter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Amount != 0 {
		n += 1 + sovCounter(uint64(m.Amount))
	}
	return n
}

func (m *MsgIncreaseCounterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NewAmount != 0 {
		n += 1 + sovCounter(uint64(m.NewAmount))
	}
	return n
}

func (m *MsgTestDependencies) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgTestDependenciesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovCounter(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCounter(uint64(l))
	}
	if m.BeforeGas != 0 {
		n += 1 + sovCounter(uint64(m.BeforeGas))
	}
	if m.AfterGas != 0 {
		n += 1 + sovCounter(uint64(m.AfterGas))
	}
	if len(m.Funds) > 0 {
		for _, e := range m.Funds {
			l = e.Size()
			n += 1 + l + sovCounter(uint64(l))
		}
	}
	return n
}

func (m *QueryCounterRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryCounterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovCounter(uint64(m.Value))
	}
	return n
}

func sovCounter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCounter(x uint64) (n int) {
	return sovCounter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCounter
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
			return fmt.Errorf("proto: MsgInit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialValue", wireType)
			}
			m.InitialValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCounter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCounter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCounter
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
func (m *MsgInitResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCounter
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
			return fmt.Errorf("proto: MsgInitResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCounter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCounter
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
func (m *MsgIncreaseCounter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCounter
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
			return fmt.Errorf("proto: MsgIncreaseCounter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIncreaseCounter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCounter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCounter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCounter
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
func (m *MsgIncreaseCounterResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCounter
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
			return fmt.Errorf("proto: MsgIncreaseCounterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIncreaseCounterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewAmount", wireType)
			}
			m.NewAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCounter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCounter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCounter
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
func (m *MsgTestDependencies) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCounter
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
			return fmt.Errorf("proto: MsgTestDependencies: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTestDependencies: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCounter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCounter
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
func (m *MsgTestDependenciesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCounter
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
			return fmt.Errorf("proto: MsgTestDependenciesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTestDependenciesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCounter
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
				return ErrInvalidLengthCounter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCounter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCounter
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
				return ErrInvalidLengthCounter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCounter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeforeGas", wireType)
			}
			m.BeforeGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCounter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BeforeGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AfterGas", wireType)
			}
			m.AfterGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCounter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AfterGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCounter
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
				return ErrInvalidLengthCounter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCounter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Funds = append(m.Funds, types.Coin{})
			if err := m.Funds[len(m.Funds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCounter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCounter
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
func (m *QueryCounterRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCounter
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
			return fmt.Errorf("proto: QueryCounterRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCounterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCounter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCounter
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
func (m *QueryCounterResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCounter
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
			return fmt.Errorf("proto: QueryCounterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCounterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCounter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCounter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCounter
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
func skipCounter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCounter
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
					return 0, ErrIntOverflowCounter
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
					return 0, ErrIntOverflowCounter
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
				return 0, ErrInvalidLengthCounter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCounter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCounter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCounter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCounter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCounter = fmt.Errorf("proto: unexpected end of group")
)
