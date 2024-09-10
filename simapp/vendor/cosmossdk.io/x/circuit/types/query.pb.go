// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/circuit/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryAccountRequest is the request type for the Query/Account RPC method.
type QueryAccountRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryAccountRequest) Reset()         { *m = QueryAccountRequest{} }
func (m *QueryAccountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAccountRequest) ProtoMessage()    {}
func (*QueryAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87c65073a3d3c1e1, []int{0}
}
func (m *QueryAccountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAccountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAccountRequest.Merge(m, src)
}
func (m *QueryAccountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAccountRequest proto.InternalMessageInfo

func (m *QueryAccountRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// AccountResponse is the response type for the Query/Account RPC method.
type AccountResponse struct {
	Permission *Permissions `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (m *AccountResponse) Reset()         { *m = AccountResponse{} }
func (m *AccountResponse) String() string { return proto.CompactTextString(m) }
func (*AccountResponse) ProtoMessage()    {}
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87c65073a3d3c1e1, []int{1}
}
func (m *AccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountResponse.Merge(m, src)
}
func (m *AccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *AccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountResponse proto.InternalMessageInfo

func (m *AccountResponse) GetPermission() *Permissions {
	if m != nil {
		return m.Permission
	}
	return nil
}

// QueryAccountsRequest is the request type for the Query/Accounts RPC method.
type QueryAccountsRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAccountsRequest) Reset()         { *m = QueryAccountsRequest{} }
func (m *QueryAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAccountsRequest) ProtoMessage()    {}
func (*QueryAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87c65073a3d3c1e1, []int{2}
}
func (m *QueryAccountsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAccountsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAccountsRequest.Merge(m, src)
}
func (m *QueryAccountsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAccountsRequest proto.InternalMessageInfo

func (m *QueryAccountsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// AccountsResponse is the response type for the Query/Accounts RPC method.
type AccountsResponse struct {
	Accounts []*GenesisAccountPermissions `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *AccountsResponse) Reset()         { *m = AccountsResponse{} }
func (m *AccountsResponse) String() string { return proto.CompactTextString(m) }
func (*AccountsResponse) ProtoMessage()    {}
func (*AccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87c65073a3d3c1e1, []int{3}
}
func (m *AccountsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountsResponse.Merge(m, src)
}
func (m *AccountsResponse) XXX_Size() int {
	return m.Size()
}
func (m *AccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountsResponse proto.InternalMessageInfo

func (m *AccountsResponse) GetAccounts() []*GenesisAccountPermissions {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *AccountsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryDisabledListRequest is the request type for the Query/DisabledList RPC method.
type QueryDisabledListRequest struct {
}

func (m *QueryDisabledListRequest) Reset()         { *m = QueryDisabledListRequest{} }
func (m *QueryDisabledListRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDisabledListRequest) ProtoMessage()    {}
func (*QueryDisabledListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_87c65073a3d3c1e1, []int{4}
}
func (m *QueryDisabledListRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDisabledListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDisabledListRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDisabledListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDisabledListRequest.Merge(m, src)
}
func (m *QueryDisabledListRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDisabledListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDisabledListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDisabledListRequest proto.InternalMessageInfo

// DisabledListResponse is the response type for the Query/DisabledList RPC method.
type DisabledListResponse struct {
	DisabledList []string `protobuf:"bytes,1,rep,name=disabled_list,json=disabledList,proto3" json:"disabled_list,omitempty"`
}

func (m *DisabledListResponse) Reset()         { *m = DisabledListResponse{} }
func (m *DisabledListResponse) String() string { return proto.CompactTextString(m) }
func (*DisabledListResponse) ProtoMessage()    {}
func (*DisabledListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_87c65073a3d3c1e1, []int{5}
}
func (m *DisabledListResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DisabledListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DisabledListResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DisabledListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisabledListResponse.Merge(m, src)
}
func (m *DisabledListResponse) XXX_Size() int {
	return m.Size()
}
func (m *DisabledListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DisabledListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DisabledListResponse proto.InternalMessageInfo

func (m *DisabledListResponse) GetDisabledList() []string {
	if m != nil {
		return m.DisabledList
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAccountRequest)(nil), "cosmos.circuit.v1.QueryAccountRequest")
	proto.RegisterType((*AccountResponse)(nil), "cosmos.circuit.v1.AccountResponse")
	proto.RegisterType((*QueryAccountsRequest)(nil), "cosmos.circuit.v1.QueryAccountsRequest")
	proto.RegisterType((*AccountsResponse)(nil), "cosmos.circuit.v1.AccountsResponse")
	proto.RegisterType((*QueryDisabledListRequest)(nil), "cosmos.circuit.v1.QueryDisabledListRequest")
	proto.RegisterType((*DisabledListResponse)(nil), "cosmos.circuit.v1.DisabledListResponse")
}

func init() { proto.RegisterFile("cosmos/circuit/v1/query.proto", fileDescriptor_87c65073a3d3c1e1) }

var fileDescriptor_87c65073a3d3c1e1 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xc1, 0x6b, 0x13, 0x4f,
	0x14, 0xc7, 0x33, 0x2d, 0xbf, 0x5f, 0xdb, 0xd7, 0x8a, 0x3a, 0xf6, 0x10, 0xb6, 0xed, 0x5a, 0xb7,
	0xd8, 0x84, 0x5a, 0x76, 0x48, 0x04, 0x2f, 0x82, 0xa0, 0x88, 0xf5, 0xe0, 0xa1, 0xdd, 0xa3, 0x07,
	0x65, 0xb2, 0x3b, 0x84, 0xc1, 0x74, 0x67, 0xbb, 0x6f, 0x13, 0x2c, 0xe2, 0xa5, 0x27, 0xbd, 0x89,
	0xfe, 0x0d, 0x1e, 0x05, 0xff, 0x0c, 0x8f, 0x05, 0x2f, 0x1e, 0x25, 0x11, 0xfc, 0x37, 0xa4, 0xb3,
	0x33, 0x9b, 0x8d, 0x9d, 0xda, 0xe3, 0xcc, 0xbc, 0xef, 0xcb, 0xe7, 0xfb, 0xbe, 0x2f, 0x0b, 0x1b,
	0xb1, 0xc2, 0x43, 0x85, 0x2c, 0x96, 0x79, 0x3c, 0x94, 0x05, 0x1b, 0x75, 0xd8, 0xd1, 0x50, 0xe4,
	0xc7, 0x61, 0x96, 0xab, 0x42, 0xd1, 0xeb, 0xe5, 0x73, 0x68, 0x9e, 0xc3, 0x51, 0xc7, 0xdb, 0x31,
	0x8a, 0x1e, 0x47, 0x51, 0xd6, 0xb2, 0x51, 0xa7, 0x27, 0x0a, 0xde, 0x61, 0x19, 0xef, 0xcb, 0x94,
	0x17, 0x52, 0xa5, 0xa5, 0xdc, 0x73, 0x74, 0x2f, 0x8e, 0x33, 0x81, 0xe6, 0x79, 0xbd, 0xaf, 0x54,
	0x7f, 0x20, 0x18, 0xcf, 0x24, 0xe3, 0x69, 0xaa, 0x0a, 0xad, 0xb5, 0xaf, 0x6b, 0x46, 0x6c, 0x7f,
	0xa3, 0x0e, 0x16, 0x30, 0xb8, 0x71, 0x70, 0x76, 0x7c, 0x18, 0xc7, 0x6a, 0x98, 0x16, 0x91, 0x38,
	0x1a, 0x0a, 0x2c, 0x68, 0x13, 0x16, 0x78, 0x92, 0xe4, 0x02, 0xb1, 0x49, 0x36, 0x49, 0x7b, 0x29,
	0xb2, 0xc7, 0xe0, 0x00, 0xae, 0x56, 0xb5, 0x98, 0xa9, 0x14, 0x05, 0x7d, 0x00, 0x90, 0x89, 0xfc,
	0x50, 0x22, 0x4a, 0x95, 0xea, 0xfa, 0xe5, 0xae, 0x1f, 0x9e, 0x73, 0x1c, 0xee, 0x57, 0x45, 0x18,
	0xd5, 0x14, 0xc1, 0x0b, 0x58, 0xad, 0x33, 0xa0, 0x85, 0x78, 0x02, 0x30, 0x9d, 0x84, 0xe9, 0xbb,
	0x6d, 0xfb, 0x9e, 0x8d, 0x2d, 0x2c, 0x9d, 0x98, 0xb1, 0x85, 0xfb, 0xbc, 0x2f, 0x8c, 0x36, 0xaa,
	0x29, 0x83, 0xcf, 0x04, 0xae, 0x4d, 0x7b, 0x1b, 0xe8, 0xa7, 0xb0, 0xc8, 0xcd, 0x5d, 0x93, 0x6c,
	0xce, 0xb7, 0x97, 0xbb, 0xbb, 0x0e, 0xe4, 0x3d, 0x91, 0x0a, 0x94, 0x68, 0xd4, 0x75, 0x03, 0x95,
	0x9a, 0xee, 0xcd, 0x60, 0xce, 0x69, 0xcc, 0xd6, 0xa5, 0x98, 0x25, 0xc6, 0x0c, 0xa7, 0x07, 0x4d,
	0x3d, 0x87, 0xc7, 0x12, 0x79, 0x6f, 0x20, 0x92, 0x67, 0x12, 0x6d, 0x20, 0xc1, 0x7d, 0x58, 0x9d,
	0xbd, 0x36, 0x36, 0xb6, 0xe0, 0x4a, 0x62, 0xee, 0x5f, 0x0e, 0x24, 0x16, 0xda, 0xcb, 0x52, 0xb4,
	0x92, 0xd4, 0x8a, 0xbb, 0x5f, 0xe6, 0xe1, 0x3f, 0xdd, 0x99, 0xbe, 0x27, 0xb0, 0x60, 0xcc, 0xd0,
	0x6d, 0x87, 0x5f, 0xc7, 0x2e, 0x78, 0x81, 0xa3, 0xee, 0xaf, 0x15, 0x08, 0xba, 0xef, 0x7e, 0x7f,
	0xdd, 0x21, 0x27, 0xdf, 0x7f, 0x7d, 0x9a, 0x6b, 0xd1, 0xdb, 0xec, 0xfc, 0xba, 0xda, 0x69, 0xb1,
	0x37, 0x66, 0x91, 0xde, 0xd2, 0x13, 0x02, 0x8b, 0x36, 0x16, 0xda, 0xba, 0x04, 0xc6, 0x2e, 0x85,
	0xb7, 0x75, 0x31, 0x4d, 0x15, 0x6e, 0xd0, 0x9e, 0xe2, 0x6c, 0xd0, 0xb5, 0x7f, 0xe0, 0xd0, 0x8f,
	0x04, 0x56, 0xea, 0x83, 0xa5, 0x77, 0x2e, 0x02, 0x71, 0xa4, 0xe2, 0xb9, 0xa8, 0x5d, 0x31, 0x05,
	0xbb, 0x53, 0xa0, 0x5b, 0xf4, 0xa6, 0x03, 0xc8, 0xe4, 0xa5, 0x33, 0x7c, 0x74, 0xef, 0xdb, 0xd8,
	0x27, 0xa7, 0x63, 0x9f, 0xfc, 0x1c, 0xfb, 0xe4, 0xc3, 0xc4, 0x6f, 0x9c, 0x4e, 0xfc, 0xc6, 0x8f,
	0x89, 0xdf, 0x78, 0xbe, 0x5e, 0x2a, 0x31, 0x79, 0x15, 0x4a, 0xc5, 0x5e, 0x57, 0x1d, 0xf4, 0xd7,
	0xa0, 0xf7, 0xbf, 0xfe, 0x4f, 0xdf, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xdd, 0xdb, 0x03,
	0x8d, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Account returns account permissions.
	Account(ctx context.Context, in *QueryAccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	// Accounts returns multiple accounts permissions.
	Accounts(ctx context.Context, in *QueryAccountsRequest, opts ...grpc.CallOption) (*AccountsResponse, error)
	// DisabledList returns a list of disabled message urls
	DisabledList(ctx context.Context, in *QueryDisabledListRequest, opts ...grpc.CallOption) (*DisabledListResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Account(ctx context.Context, in *QueryAccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/cosmos.circuit.v1.Query/Account", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Accounts(ctx context.Context, in *QueryAccountsRequest, opts ...grpc.CallOption) (*AccountsResponse, error) {
	out := new(AccountsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.circuit.v1.Query/Accounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DisabledList(ctx context.Context, in *QueryDisabledListRequest, opts ...grpc.CallOption) (*DisabledListResponse, error) {
	out := new(DisabledListResponse)
	err := c.cc.Invoke(ctx, "/cosmos.circuit.v1.Query/DisabledList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Account returns account permissions.
	Account(context.Context, *QueryAccountRequest) (*AccountResponse, error)
	// Accounts returns multiple accounts permissions.
	Accounts(context.Context, *QueryAccountsRequest) (*AccountsResponse, error)
	// DisabledList returns a list of disabled message urls
	DisabledList(context.Context, *QueryDisabledListRequest) (*DisabledListResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Account(ctx context.Context, req *QueryAccountRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Account not implemented")
}
func (*UnimplementedQueryServer) Accounts(ctx context.Context, req *QueryAccountsRequest) (*AccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Accounts not implemented")
}
func (*UnimplementedQueryServer) DisabledList(ctx context.Context, req *QueryDisabledListRequest) (*DisabledListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisabledList not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Account_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Account(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.circuit.v1.Query/Account",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Account(ctx, req.(*QueryAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Accounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Accounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.circuit.v1.Query/Accounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Accounts(ctx, req.(*QueryAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DisabledList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDisabledListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DisabledList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.circuit.v1.Query/DisabledList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DisabledList(ctx, req.(*QueryDisabledListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.circuit.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Account",
			Handler:    _Query_Account_Handler,
		},
		{
			MethodName: "Accounts",
			Handler:    _Query_Accounts_Handler,
		},
		{
			MethodName: "DisabledList",
			Handler:    _Query_DisabledList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/circuit/v1/query.proto",
}

func (m *QueryAccountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAccountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAccountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Permission != nil {
		{
			size, err := m.Permission.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAccountsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAccountsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAccountsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AccountsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Accounts) > 0 {
		for iNdEx := len(m.Accounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Accounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryDisabledListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDisabledListRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDisabledListRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *DisabledListResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DisabledListResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DisabledListResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DisabledList) > 0 {
		for iNdEx := len(m.DisabledList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DisabledList[iNdEx])
			copy(dAtA[i:], m.DisabledList[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.DisabledList[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAccountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *AccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Permission != nil {
		l = m.Permission.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAccountsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *AccountsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Accounts) > 0 {
		for _, e := range m.Accounts {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDisabledListRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *DisabledListResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DisabledList) > 0 {
		for _, s := range m.DisabledList {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAccountRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAccountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAccountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *AccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: AccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Permission == nil {
				m.Permission = &Permissions{}
			}
			if err := m.Permission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAccountsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAccountsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAccountsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *AccountsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: AccountsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accounts = append(m.Accounts, &GenesisAccountPermissions{})
			if err := m.Accounts[len(m.Accounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryDisabledListRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDisabledListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDisabledListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *DisabledListResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: DisabledListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DisabledListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisabledList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisabledList = append(m.DisabledList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
