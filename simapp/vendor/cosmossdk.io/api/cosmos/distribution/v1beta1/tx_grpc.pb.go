// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cosmos/distribution/v1beta1/tx.proto

package distributionv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Msg_SetWithdrawAddress_FullMethodName          = "/cosmos.distribution.v1beta1.Msg/SetWithdrawAddress"
	Msg_WithdrawDelegatorReward_FullMethodName     = "/cosmos.distribution.v1beta1.Msg/WithdrawDelegatorReward"
	Msg_WithdrawValidatorCommission_FullMethodName = "/cosmos.distribution.v1beta1.Msg/WithdrawValidatorCommission"
	Msg_FundCommunityPool_FullMethodName           = "/cosmos.distribution.v1beta1.Msg/FundCommunityPool"
	Msg_UpdateParams_FullMethodName                = "/cosmos.distribution.v1beta1.Msg/UpdateParams"
	Msg_CommunityPoolSpend_FullMethodName          = "/cosmos.distribution.v1beta1.Msg/CommunityPoolSpend"
	Msg_DepositValidatorRewardsPool_FullMethodName = "/cosmos.distribution.v1beta1.Msg/DepositValidatorRewardsPool"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Msg defines the distribution Msg service.
type MsgClient interface {
	// SetWithdrawAddress defines a method to change the withdraw address
	// for a delegator (or validator self-delegation).
	SetWithdrawAddress(ctx context.Context, in *MsgSetWithdrawAddress, opts ...grpc.CallOption) (*MsgSetWithdrawAddressResponse, error)
	// WithdrawDelegatorReward defines a method to withdraw rewards of delegator
	// from a single validator.
	WithdrawDelegatorReward(ctx context.Context, in *MsgWithdrawDelegatorReward, opts ...grpc.CallOption) (*MsgWithdrawDelegatorRewardResponse, error)
	// WithdrawValidatorCommission defines a method to withdraw the
	// full commission to the validator address.
	WithdrawValidatorCommission(ctx context.Context, in *MsgWithdrawValidatorCommission, opts ...grpc.CallOption) (*MsgWithdrawValidatorCommissionResponse, error)
	// Deprecated: Do not use.
	// FundCommunityPool defines a method to allow an account to directly
	// fund the community pool.
	//
	// Deprecated: Use x/protocolpool module's FundCommunityPool instead.
	FundCommunityPool(ctx context.Context, in *MsgFundCommunityPool, opts ...grpc.CallOption) (*MsgFundCommunityPoolResponse, error)
	// UpdateParams defines a governance operation for updating the x/distribution
	// module parameters. The authority is defined in the keeper.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// CommunityPoolSpend defines a governance operation for sending tokens from
	// the community pool in the x/distribution module to another account, which
	// could be the governance module itself. The authority is defined in the
	// keeper.
	//
	// Deprecated: Use x/protocolpool module's CommunityPoolSpend instead.
	CommunityPoolSpend(ctx context.Context, in *MsgCommunityPoolSpend, opts ...grpc.CallOption) (*MsgCommunityPoolSpendResponse, error)
	// DepositValidatorRewardsPool defines a method to provide additional rewards
	// to delegators to a specific validator.
	DepositValidatorRewardsPool(ctx context.Context, in *MsgDepositValidatorRewardsPool, opts ...grpc.CallOption) (*MsgDepositValidatorRewardsPoolResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetWithdrawAddress(ctx context.Context, in *MsgSetWithdrawAddress, opts ...grpc.CallOption) (*MsgSetWithdrawAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgSetWithdrawAddressResponse)
	err := c.cc.Invoke(ctx, Msg_SetWithdrawAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawDelegatorReward(ctx context.Context, in *MsgWithdrawDelegatorReward, opts ...grpc.CallOption) (*MsgWithdrawDelegatorRewardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgWithdrawDelegatorRewardResponse)
	err := c.cc.Invoke(ctx, Msg_WithdrawDelegatorReward_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawValidatorCommission(ctx context.Context, in *MsgWithdrawValidatorCommission, opts ...grpc.CallOption) (*MsgWithdrawValidatorCommissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgWithdrawValidatorCommissionResponse)
	err := c.cc.Invoke(ctx, Msg_WithdrawValidatorCommission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *msgClient) FundCommunityPool(ctx context.Context, in *MsgFundCommunityPool, opts ...grpc.CallOption) (*MsgFundCommunityPoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgFundCommunityPoolResponse)
	err := c.cc.Invoke(ctx, Msg_FundCommunityPool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CommunityPoolSpend(ctx context.Context, in *MsgCommunityPoolSpend, opts ...grpc.CallOption) (*MsgCommunityPoolSpendResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgCommunityPoolSpendResponse)
	err := c.cc.Invoke(ctx, Msg_CommunityPoolSpend_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DepositValidatorRewardsPool(ctx context.Context, in *MsgDepositValidatorRewardsPool, opts ...grpc.CallOption) (*MsgDepositValidatorRewardsPoolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgDepositValidatorRewardsPoolResponse)
	err := c.cc.Invoke(ctx, Msg_DepositValidatorRewardsPool_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility.
//
// Msg defines the distribution Msg service.
type MsgServer interface {
	// SetWithdrawAddress defines a method to change the withdraw address
	// for a delegator (or validator self-delegation).
	SetWithdrawAddress(context.Context, *MsgSetWithdrawAddress) (*MsgSetWithdrawAddressResponse, error)
	// WithdrawDelegatorReward defines a method to withdraw rewards of delegator
	// from a single validator.
	WithdrawDelegatorReward(context.Context, *MsgWithdrawDelegatorReward) (*MsgWithdrawDelegatorRewardResponse, error)
	// WithdrawValidatorCommission defines a method to withdraw the
	// full commission to the validator address.
	WithdrawValidatorCommission(context.Context, *MsgWithdrawValidatorCommission) (*MsgWithdrawValidatorCommissionResponse, error)
	// Deprecated: Do not use.
	// FundCommunityPool defines a method to allow an account to directly
	// fund the community pool.
	//
	// Deprecated: Use x/protocolpool module's FundCommunityPool instead.
	FundCommunityPool(context.Context, *MsgFundCommunityPool) (*MsgFundCommunityPoolResponse, error)
	// UpdateParams defines a governance operation for updating the x/distribution
	// module parameters. The authority is defined in the keeper.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	// CommunityPoolSpend defines a governance operation for sending tokens from
	// the community pool in the x/distribution module to another account, which
	// could be the governance module itself. The authority is defined in the
	// keeper.
	//
	// Deprecated: Use x/protocolpool module's CommunityPoolSpend instead.
	CommunityPoolSpend(context.Context, *MsgCommunityPoolSpend) (*MsgCommunityPoolSpendResponse, error)
	// DepositValidatorRewardsPool defines a method to provide additional rewards
	// to delegators to a specific validator.
	DepositValidatorRewardsPool(context.Context, *MsgDepositValidatorRewardsPool) (*MsgDepositValidatorRewardsPoolResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMsgServer struct{}

func (UnimplementedMsgServer) SetWithdrawAddress(context.Context, *MsgSetWithdrawAddress) (*MsgSetWithdrawAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWithdrawAddress not implemented")
}
func (UnimplementedMsgServer) WithdrawDelegatorReward(context.Context, *MsgWithdrawDelegatorReward) (*MsgWithdrawDelegatorRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawDelegatorReward not implemented")
}
func (UnimplementedMsgServer) WithdrawValidatorCommission(context.Context, *MsgWithdrawValidatorCommission) (*MsgWithdrawValidatorCommissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawValidatorCommission not implemented")
}
func (UnimplementedMsgServer) FundCommunityPool(context.Context, *MsgFundCommunityPool) (*MsgFundCommunityPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundCommunityPool not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CommunityPoolSpend(context.Context, *MsgCommunityPoolSpend) (*MsgCommunityPoolSpendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommunityPoolSpend not implemented")
}
func (UnimplementedMsgServer) DepositValidatorRewardsPool(context.Context, *MsgDepositValidatorRewardsPool) (*MsgDepositValidatorRewardsPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositValidatorRewardsPool not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}
func (UnimplementedMsgServer) testEmbeddedByValue()             {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	// If the following call pancis, it indicates UnimplementedMsgServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_SetWithdrawAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetWithdrawAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetWithdrawAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SetWithdrawAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetWithdrawAddress(ctx, req.(*MsgSetWithdrawAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawDelegatorReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawDelegatorReward)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawDelegatorReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_WithdrawDelegatorReward_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawDelegatorReward(ctx, req.(*MsgWithdrawDelegatorReward))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawValidatorCommission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawValidatorCommission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawValidatorCommission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_WithdrawValidatorCommission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawValidatorCommission(ctx, req.(*MsgWithdrawValidatorCommission))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_FundCommunityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFundCommunityPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FundCommunityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_FundCommunityPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FundCommunityPool(ctx, req.(*MsgFundCommunityPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CommunityPoolSpend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCommunityPoolSpend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CommunityPoolSpend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CommunityPoolSpend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CommunityPoolSpend(ctx, req.(*MsgCommunityPoolSpend))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DepositValidatorRewardsPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDepositValidatorRewardsPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DepositValidatorRewardsPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DepositValidatorRewardsPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DepositValidatorRewardsPool(ctx, req.(*MsgDepositValidatorRewardsPool))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.distribution.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetWithdrawAddress",
			Handler:    _Msg_SetWithdrawAddress_Handler,
		},
		{
			MethodName: "WithdrawDelegatorReward",
			Handler:    _Msg_WithdrawDelegatorReward_Handler,
		},
		{
			MethodName: "WithdrawValidatorCommission",
			Handler:    _Msg_WithdrawValidatorCommission_Handler,
		},
		{
			MethodName: "FundCommunityPool",
			Handler:    _Msg_FundCommunityPool_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CommunityPoolSpend",
			Handler:    _Msg_CommunityPoolSpend_Handler,
		},
		{
			MethodName: "DepositValidatorRewardsPool",
			Handler:    _Msg_DepositValidatorRewardsPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/distribution/v1beta1/tx.proto",
}