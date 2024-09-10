// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cometbft/types/v1/evidence.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Evidence is a generic type for wrapping evidence of misbehavior by a validator.
type Evidence struct {
	// The type of evidence.
	//
	// Types that are valid to be assigned to Sum:
	//	*Evidence_DuplicateVoteEvidence
	//	*Evidence_LightClientAttackEvidence
	Sum isEvidence_Sum `protobuf_oneof:"sum"`
}

func (m *Evidence) Reset()         { *m = Evidence{} }
func (m *Evidence) String() string { return proto.CompactTextString(m) }
func (*Evidence) ProtoMessage()    {}
func (*Evidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c96acf1a4e66b9a, []int{0}
}
func (m *Evidence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Evidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Evidence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Evidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Evidence.Merge(m, src)
}
func (m *Evidence) XXX_Size() int {
	return m.Size()
}
func (m *Evidence) XXX_DiscardUnknown() {
	xxx_messageInfo_Evidence.DiscardUnknown(m)
}

var xxx_messageInfo_Evidence proto.InternalMessageInfo

type isEvidence_Sum interface {
	isEvidence_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Evidence_DuplicateVoteEvidence struct {
	DuplicateVoteEvidence *DuplicateVoteEvidence `protobuf:"bytes,1,opt,name=duplicate_vote_evidence,json=duplicateVoteEvidence,proto3,oneof" json:"duplicate_vote_evidence,omitempty"`
}
type Evidence_LightClientAttackEvidence struct {
	LightClientAttackEvidence *LightClientAttackEvidence `protobuf:"bytes,2,opt,name=light_client_attack_evidence,json=lightClientAttackEvidence,proto3,oneof" json:"light_client_attack_evidence,omitempty"`
}

func (*Evidence_DuplicateVoteEvidence) isEvidence_Sum()     {}
func (*Evidence_LightClientAttackEvidence) isEvidence_Sum() {}

func (m *Evidence) GetSum() isEvidence_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Evidence) GetDuplicateVoteEvidence() *DuplicateVoteEvidence {
	if x, ok := m.GetSum().(*Evidence_DuplicateVoteEvidence); ok {
		return x.DuplicateVoteEvidence
	}
	return nil
}

func (m *Evidence) GetLightClientAttackEvidence() *LightClientAttackEvidence {
	if x, ok := m.GetSum().(*Evidence_LightClientAttackEvidence); ok {
		return x.LightClientAttackEvidence
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Evidence) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Evidence_DuplicateVoteEvidence)(nil),
		(*Evidence_LightClientAttackEvidence)(nil),
	}
}

// DuplicateVoteEvidence contains evidence of a validator signed two conflicting votes.
type DuplicateVoteEvidence struct {
	VoteA            *Vote     `protobuf:"bytes,1,opt,name=vote_a,json=voteA,proto3" json:"vote_a,omitempty"`
	VoteB            *Vote     `protobuf:"bytes,2,opt,name=vote_b,json=voteB,proto3" json:"vote_b,omitempty"`
	TotalVotingPower int64     `protobuf:"varint,3,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
	ValidatorPower   int64     `protobuf:"varint,4,opt,name=validator_power,json=validatorPower,proto3" json:"validator_power,omitempty"`
	Timestamp        time.Time `protobuf:"bytes,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *DuplicateVoteEvidence) Reset()         { *m = DuplicateVoteEvidence{} }
func (m *DuplicateVoteEvidence) String() string { return proto.CompactTextString(m) }
func (*DuplicateVoteEvidence) ProtoMessage()    {}
func (*DuplicateVoteEvidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c96acf1a4e66b9a, []int{1}
}
func (m *DuplicateVoteEvidence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DuplicateVoteEvidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DuplicateVoteEvidence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DuplicateVoteEvidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DuplicateVoteEvidence.Merge(m, src)
}
func (m *DuplicateVoteEvidence) XXX_Size() int {
	return m.Size()
}
func (m *DuplicateVoteEvidence) XXX_DiscardUnknown() {
	xxx_messageInfo_DuplicateVoteEvidence.DiscardUnknown(m)
}

var xxx_messageInfo_DuplicateVoteEvidence proto.InternalMessageInfo

func (m *DuplicateVoteEvidence) GetVoteA() *Vote {
	if m != nil {
		return m.VoteA
	}
	return nil
}

func (m *DuplicateVoteEvidence) GetVoteB() *Vote {
	if m != nil {
		return m.VoteB
	}
	return nil
}

func (m *DuplicateVoteEvidence) GetTotalVotingPower() int64 {
	if m != nil {
		return m.TotalVotingPower
	}
	return 0
}

func (m *DuplicateVoteEvidence) GetValidatorPower() int64 {
	if m != nil {
		return m.ValidatorPower
	}
	return 0
}

func (m *DuplicateVoteEvidence) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

// LightClientAttackEvidence contains evidence of a set of validators attempting to mislead a light client.
type LightClientAttackEvidence struct {
	ConflictingBlock    *LightBlock  `protobuf:"bytes,1,opt,name=conflicting_block,json=conflictingBlock,proto3" json:"conflicting_block,omitempty"`
	CommonHeight        int64        `protobuf:"varint,2,opt,name=common_height,json=commonHeight,proto3" json:"common_height,omitempty"`
	ByzantineValidators []*Validator `protobuf:"bytes,3,rep,name=byzantine_validators,json=byzantineValidators,proto3" json:"byzantine_validators,omitempty"`
	TotalVotingPower    int64        `protobuf:"varint,4,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
	Timestamp           time.Time    `protobuf:"bytes,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *LightClientAttackEvidence) Reset()         { *m = LightClientAttackEvidence{} }
func (m *LightClientAttackEvidence) String() string { return proto.CompactTextString(m) }
func (*LightClientAttackEvidence) ProtoMessage()    {}
func (*LightClientAttackEvidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c96acf1a4e66b9a, []int{2}
}
func (m *LightClientAttackEvidence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LightClientAttackEvidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LightClientAttackEvidence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LightClientAttackEvidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightClientAttackEvidence.Merge(m, src)
}
func (m *LightClientAttackEvidence) XXX_Size() int {
	return m.Size()
}
func (m *LightClientAttackEvidence) XXX_DiscardUnknown() {
	xxx_messageInfo_LightClientAttackEvidence.DiscardUnknown(m)
}

var xxx_messageInfo_LightClientAttackEvidence proto.InternalMessageInfo

func (m *LightClientAttackEvidence) GetConflictingBlock() *LightBlock {
	if m != nil {
		return m.ConflictingBlock
	}
	return nil
}

func (m *LightClientAttackEvidence) GetCommonHeight() int64 {
	if m != nil {
		return m.CommonHeight
	}
	return 0
}

func (m *LightClientAttackEvidence) GetByzantineValidators() []*Validator {
	if m != nil {
		return m.ByzantineValidators
	}
	return nil
}

func (m *LightClientAttackEvidence) GetTotalVotingPower() int64 {
	if m != nil {
		return m.TotalVotingPower
	}
	return 0
}

func (m *LightClientAttackEvidence) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

// EvidenceList is a list of evidence.
type EvidenceList struct {
	Evidence []Evidence `protobuf:"bytes,1,rep,name=evidence,proto3" json:"evidence"`
}

func (m *EvidenceList) Reset()         { *m = EvidenceList{} }
func (m *EvidenceList) String() string { return proto.CompactTextString(m) }
func (*EvidenceList) ProtoMessage()    {}
func (*EvidenceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c96acf1a4e66b9a, []int{3}
}
func (m *EvidenceList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EvidenceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EvidenceList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EvidenceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvidenceList.Merge(m, src)
}
func (m *EvidenceList) XXX_Size() int {
	return m.Size()
}
func (m *EvidenceList) XXX_DiscardUnknown() {
	xxx_messageInfo_EvidenceList.DiscardUnknown(m)
}

var xxx_messageInfo_EvidenceList proto.InternalMessageInfo

func (m *EvidenceList) GetEvidence() []Evidence {
	if m != nil {
		return m.Evidence
	}
	return nil
}

func init() {
	proto.RegisterType((*Evidence)(nil), "cometbft.types.v1.Evidence")
	proto.RegisterType((*DuplicateVoteEvidence)(nil), "cometbft.types.v1.DuplicateVoteEvidence")
	proto.RegisterType((*LightClientAttackEvidence)(nil), "cometbft.types.v1.LightClientAttackEvidence")
	proto.RegisterType((*EvidenceList)(nil), "cometbft.types.v1.EvidenceList")
}

func init() { proto.RegisterFile("cometbft/types/v1/evidence.proto", fileDescriptor_4c96acf1a4e66b9a) }

var fileDescriptor_4c96acf1a4e66b9a = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe3, 0x38, 0xa9, 0xc2, 0xb6, 0x40, 0xbb, 0xb4, 0x6a, 0x1a, 0x5a, 0x27, 0x84, 0x03,
	0x39, 0x54, 0xb6, 0x1a, 0xce, 0x1c, 0x6a, 0x40, 0xaa, 0x50, 0x10, 0xc8, 0x42, 0x3d, 0x70, 0xb1,
	0xd6, 0xce, 0xc6, 0x59, 0xd5, 0xf6, 0x5a, 0xf1, 0x24, 0xa8, 0x3c, 0x45, 0xc5, 0x53, 0xf5, 0x46,
	0x8f, 0x9c, 0x00, 0x25, 0x12, 0xcf, 0x81, 0x76, 0x6d, 0x6f, 0x23, 0xc5, 0xe6, 0xc4, 0x6d, 0x33,
	0xf3, 0xfd, 0x3b, 0x33, 0xff, 0xc4, 0x8b, 0x7a, 0x3e, 0x8f, 0x28, 0x78, 0x13, 0xb0, 0xe0, 0x3a,
	0xa1, 0xa9, 0xb5, 0x38, 0xb3, 0xe8, 0x82, 0x8d, 0x69, 0xec, 0x53, 0x33, 0x99, 0x71, 0xe0, 0x78,
	0xaf, 0x20, 0x4c, 0x49, 0x98, 0x8b, 0xb3, 0xce, 0xc9, 0xa6, 0x28, 0xcb, 0x49, 0x45, 0xe7, 0xd9,
	0x66, 0x7a, 0x41, 0x42, 0x36, 0x26, 0xc0, 0x67, 0x39, 0xb2, 0x1f, 0xf0, 0x80, 0xcb, 0xa3, 0x25,
	0x4e, 0x79, 0xb4, 0x1b, 0x70, 0x1e, 0x84, 0xd4, 0x92, 0xbf, 0xbc, 0xf9, 0xc4, 0x02, 0x16, 0xd1,
	0x14, 0x48, 0x94, 0x64, 0x40, 0xff, 0x8f, 0x86, 0x5a, 0x6f, 0xf3, 0xf6, 0xb0, 0x87, 0x0e, 0xc7,
	0xf3, 0x24, 0x64, 0x3e, 0x01, 0xea, 0x2e, 0x38, 0x50, 0xb7, 0xe8, 0xbc, 0xad, 0xf5, 0xb4, 0xc1,
	0xf6, 0x70, 0x60, 0x6e, 0xb4, 0x6e, 0xbe, 0x29, 0x14, 0x97, 0x1c, 0x68, 0x71, 0xd5, 0x45, 0xcd,
	0x39, 0x18, 0x97, 0x25, 0x30, 0x47, 0xc7, 0x21, 0x0b, 0xa6, 0xe0, 0xfa, 0x21, 0xa3, 0x31, 0xb8,
	0x04, 0x80, 0xf8, 0x57, 0xf7, 0x85, 0xea, 0xb2, 0xd0, 0x69, 0x49, 0xa1, 0x91, 0x90, 0xbd, 0x96,
	0xaa, 0x73, 0x29, 0x5a, 0x2b, 0x76, 0x14, 0x56, 0x25, 0xed, 0x26, 0xd2, 0xd3, 0x79, 0xd4, 0xff,
	0x56, 0x47, 0x07, 0xa5, 0xad, 0x62, 0x13, 0x6d, 0xc9, 0x59, 0x49, 0x3e, 0xe4, 0x61, 0x49, 0x6d,
	0x21, 0x70, 0x9a, 0x02, 0x3b, 0x57, 0xbc, 0x97, 0xf7, 0xfa, 0x6f, 0xde, 0xc6, 0xa7, 0x08, 0x03,
	0x07, 0x12, 0x0a, 0x47, 0x59, 0x1c, 0xb8, 0x09, 0xff, 0x42, 0x67, 0x6d, 0xbd, 0xa7, 0x0d, 0x74,
	0x67, 0x57, 0x66, 0x2e, 0x65, 0xe2, 0xa3, 0x88, 0xe3, 0x17, 0xe8, 0xb1, 0x5a, 0x6d, 0x8e, 0x36,
	0x24, 0xfa, 0x48, 0x85, 0x33, 0xd0, 0x46, 0x0f, 0xd4, 0x32, 0xdb, 0x4d, 0xd9, 0x49, 0xc7, 0xcc,
	0xd6, 0x6d, 0x16, 0xeb, 0x36, 0x3f, 0x15, 0x84, 0xdd, 0xba, 0xfd, 0xd9, 0xad, 0xdd, 0xfc, 0xea,
	0x6a, 0xce, 0xbd, 0xac, 0xff, 0xbd, 0x8e, 0x8e, 0x2a, 0x6d, 0xc5, 0xef, 0xd0, 0x9e, 0xcf, 0xe3,
	0x49, 0xc8, 0x7c, 0xd9, 0xb7, 0x17, 0x72, 0xff, 0x2a, 0xf7, 0xe8, 0xa4, 0x6a, 0x3f, 0xb6, 0x80,
	0x9c, 0xdd, 0x35, 0x9d, 0x8c, 0xe0, 0xe7, 0xe8, 0xa1, 0xcf, 0xa3, 0x88, 0xc7, 0xee, 0x94, 0x0a,
	0x4e, 0x7a, 0xa7, 0x3b, 0x3b, 0x59, 0xf0, 0x42, 0xc6, 0xf0, 0x07, 0xb4, 0xef, 0x5d, 0x7f, 0x25,
	0x31, 0xb0, 0x98, 0xba, 0x6a, 0xdc, 0xb4, 0xad, 0xf7, 0xf4, 0xc1, 0xf6, 0xf0, 0xb8, 0xcc, 0xe7,
	0x02, 0x72, 0x9e, 0x28, 0xa5, 0x8a, 0xa5, 0x15, 0xd6, 0x37, 0x2a, 0xac, 0xff, 0x1f, 0x8e, 0xbe,
	0x47, 0x3b, 0x85, 0x7f, 0x23, 0x96, 0x02, 0x7e, 0x85, 0x5a, 0x6b, 0xdf, 0x90, 0x18, 0xe3, 0x69,
	0xc9, 0x18, 0xea, 0xcf, 0xda, 0x10, 0x77, 0x3a, 0x4a, 0x62, 0x8f, 0x6e, 0x97, 0x86, 0x76, 0xb7,
	0x34, 0xb4, 0xdf, 0x4b, 0x43, 0xbb, 0x59, 0x19, 0xb5, 0xbb, 0x95, 0x51, 0xfb, 0xb1, 0x32, 0x6a,
	0x9f, 0x87, 0x01, 0x83, 0xe9, 0xdc, 0x13, 0x97, 0x59, 0xea, 0x75, 0x50, 0x07, 0x92, 0x30, 0x6b,
	0xe3, 0xcd, 0xf0, 0xb6, 0xe4, 0x14, 0x2f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x50, 0xcc, 0xfa,
	0x44, 0xa3, 0x04, 0x00, 0x00,
}

func (m *Evidence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Evidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Evidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Evidence_DuplicateVoteEvidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Evidence_DuplicateVoteEvidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DuplicateVoteEvidence != nil {
		{
			size, err := m.DuplicateVoteEvidence.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvidence(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Evidence_LightClientAttackEvidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Evidence_LightClientAttackEvidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.LightClientAttackEvidence != nil {
		{
			size, err := m.LightClientAttackEvidence.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvidence(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *DuplicateVoteEvidence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DuplicateVoteEvidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DuplicateVoteEvidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintEvidence(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x2a
	if m.ValidatorPower != 0 {
		i = encodeVarintEvidence(dAtA, i, uint64(m.ValidatorPower))
		i--
		dAtA[i] = 0x20
	}
	if m.TotalVotingPower != 0 {
		i = encodeVarintEvidence(dAtA, i, uint64(m.TotalVotingPower))
		i--
		dAtA[i] = 0x18
	}
	if m.VoteB != nil {
		{
			size, err := m.VoteB.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvidence(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.VoteA != nil {
		{
			size, err := m.VoteA.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvidence(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LightClientAttackEvidence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LightClientAttackEvidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LightClientAttackEvidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n6, err6 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err6 != nil {
		return 0, err6
	}
	i -= n6
	i = encodeVarintEvidence(dAtA, i, uint64(n6))
	i--
	dAtA[i] = 0x2a
	if m.TotalVotingPower != 0 {
		i = encodeVarintEvidence(dAtA, i, uint64(m.TotalVotingPower))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ByzantineValidators) > 0 {
		for iNdEx := len(m.ByzantineValidators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ByzantineValidators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvidence(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.CommonHeight != 0 {
		i = encodeVarintEvidence(dAtA, i, uint64(m.CommonHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.ConflictingBlock != nil {
		{
			size, err := m.ConflictingBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvidence(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EvidenceList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EvidenceList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EvidenceList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Evidence) > 0 {
		for iNdEx := len(m.Evidence) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Evidence[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvidence(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvidence(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvidence(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Evidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Evidence_DuplicateVoteEvidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DuplicateVoteEvidence != nil {
		l = m.DuplicateVoteEvidence.Size()
		n += 1 + l + sovEvidence(uint64(l))
	}
	return n
}
func (m *Evidence_LightClientAttackEvidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LightClientAttackEvidence != nil {
		l = m.LightClientAttackEvidence.Size()
		n += 1 + l + sovEvidence(uint64(l))
	}
	return n
}
func (m *DuplicateVoteEvidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VoteA != nil {
		l = m.VoteA.Size()
		n += 1 + l + sovEvidence(uint64(l))
	}
	if m.VoteB != nil {
		l = m.VoteB.Size()
		n += 1 + l + sovEvidence(uint64(l))
	}
	if m.TotalVotingPower != 0 {
		n += 1 + sovEvidence(uint64(m.TotalVotingPower))
	}
	if m.ValidatorPower != 0 {
		n += 1 + sovEvidence(uint64(m.ValidatorPower))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovEvidence(uint64(l))
	return n
}

func (m *LightClientAttackEvidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConflictingBlock != nil {
		l = m.ConflictingBlock.Size()
		n += 1 + l + sovEvidence(uint64(l))
	}
	if m.CommonHeight != 0 {
		n += 1 + sovEvidence(uint64(m.CommonHeight))
	}
	if len(m.ByzantineValidators) > 0 {
		for _, e := range m.ByzantineValidators {
			l = e.Size()
			n += 1 + l + sovEvidence(uint64(l))
		}
	}
	if m.TotalVotingPower != 0 {
		n += 1 + sovEvidence(uint64(m.TotalVotingPower))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovEvidence(uint64(l))
	return n
}

func (m *EvidenceList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Evidence) > 0 {
		for _, e := range m.Evidence {
			l = e.Size()
			n += 1 + l + sovEvidence(uint64(l))
		}
	}
	return n
}

func sovEvidence(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvidence(x uint64) (n int) {
	return sovEvidence(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Evidence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvidence
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
			return fmt.Errorf("proto: Evidence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Evidence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DuplicateVoteEvidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DuplicateVoteEvidence{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Evidence_DuplicateVoteEvidence{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LightClientAttackEvidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &LightClientAttackEvidence{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Evidence_LightClientAttackEvidence{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvidence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvidence
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
func (m *DuplicateVoteEvidence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvidence
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
			return fmt.Errorf("proto: DuplicateVoteEvidence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DuplicateVoteEvidence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteA", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VoteA == nil {
				m.VoteA = &Vote{}
			}
			if err := m.VoteA.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteB", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VoteB == nil {
				m.VoteB = &Vote{}
			}
			if err := m.VoteB.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVotingPower", wireType)
			}
			m.TotalVotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalVotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorPower", wireType)
			}
			m.ValidatorPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidatorPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvidence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvidence
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
func (m *LightClientAttackEvidence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvidence
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
			return fmt.Errorf("proto: LightClientAttackEvidence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LightClientAttackEvidence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConflictingBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConflictingBlock == nil {
				m.ConflictingBlock = &LightBlock{}
			}
			if err := m.ConflictingBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonHeight", wireType)
			}
			m.CommonHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommonHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ByzantineValidators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ByzantineValidators = append(m.ByzantineValidators, &Validator{})
			if err := m.ByzantineValidators[len(m.ByzantineValidators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVotingPower", wireType)
			}
			m.TotalVotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalVotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvidence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvidence
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
func (m *EvidenceList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvidence
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
			return fmt.Errorf("proto: EvidenceList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EvidenceList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Evidence = append(m.Evidence, Evidence{})
			if err := m.Evidence[len(m.Evidence)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvidence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvidence
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
func skipEvidence(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvidence
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
					return 0, ErrIntOverflowEvidence
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
					return 0, ErrIntOverflowEvidence
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
				return 0, ErrInvalidLengthEvidence
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvidence
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvidence
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvidence        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvidence          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvidence = fmt.Errorf("proto: unexpected end of group")
)
