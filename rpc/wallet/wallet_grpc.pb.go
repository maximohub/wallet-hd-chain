// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: savourrpc/wallet.proto

package wallet

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServiceClient interface {
	GetSupportCoins(ctx context.Context, in *SupportCoinsRequest, opts ...grpc.CallOption) (*SupportCoinsResponse, error)
	GetNonce(ctx context.Context, in *NonceRequest, opts ...grpc.CallOption) (*NonceResponse, error)
	GetGasPrice(ctx context.Context, in *GasPriceRequest, opts ...grpc.CallOption) (*GasPriceResponse, error)
	SendTx(ctx context.Context, in *SendTxRequest, opts ...grpc.CallOption) (*SendTxResponse, error)
	GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	GetTxByAddress(ctx context.Context, in *TxAddressRequest, opts ...grpc.CallOption) (*TxAddressResponse, error)
	GetTxByHash(ctx context.Context, in *TxHashRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	GetUtxo(ctx context.Context, in *UtxoRequest, opts ...grpc.CallOption) (*UtxoResponse, error)
	GetMinRent(ctx context.Context, in *MinRentRequest, opts ...grpc.CallOption) (*MinRentResponse, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) GetSupportCoins(ctx context.Context, in *SupportCoinsRequest, opts ...grpc.CallOption) (*SupportCoinsResponse, error) {
	out := new(SupportCoinsResponse)
	err := c.cc.Invoke(ctx, "/savourrpc.wallet.WalletService/getSupportCoins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetNonce(ctx context.Context, in *NonceRequest, opts ...grpc.CallOption) (*NonceResponse, error) {
	out := new(NonceResponse)
	err := c.cc.Invoke(ctx, "/savourrpc.wallet.WalletService/getNonce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetGasPrice(ctx context.Context, in *GasPriceRequest, opts ...grpc.CallOption) (*GasPriceResponse, error) {
	out := new(GasPriceResponse)
	err := c.cc.Invoke(ctx, "/savourrpc.wallet.WalletService/getGasPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) SendTx(ctx context.Context, in *SendTxRequest, opts ...grpc.CallOption) (*SendTxResponse, error) {
	out := new(SendTxResponse)
	err := c.cc.Invoke(ctx, "/savourrpc.wallet.WalletService/SendTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/savourrpc.wallet.WalletService/getBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetTxByAddress(ctx context.Context, in *TxAddressRequest, opts ...grpc.CallOption) (*TxAddressResponse, error) {
	out := new(TxAddressResponse)
	err := c.cc.Invoke(ctx, "/savourrpc.wallet.WalletService/getTxByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetTxByHash(ctx context.Context, in *TxHashRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, "/savourrpc.wallet.WalletService/getTxByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/savourrpc.wallet.WalletService/getAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetUtxo(ctx context.Context, in *UtxoRequest, opts ...grpc.CallOption) (*UtxoResponse, error) {
	out := new(UtxoResponse)
	err := c.cc.Invoke(ctx, "/savourrpc.wallet.WalletService/getUtxo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetMinRent(ctx context.Context, in *MinRentRequest, opts ...grpc.CallOption) (*MinRentResponse, error) {
	out := new(MinRentResponse)
	err := c.cc.Invoke(ctx, "/savourrpc.wallet.WalletService/getMinRent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations must embed UnimplementedWalletServiceServer
// for forward compatibility
type WalletServiceServer interface {
	GetSupportCoins(context.Context, *SupportCoinsRequest) (*SupportCoinsResponse, error)
	GetNonce(context.Context, *NonceRequest) (*NonceResponse, error)
	GetGasPrice(context.Context, *GasPriceRequest) (*GasPriceResponse, error)
	SendTx(context.Context, *SendTxRequest) (*SendTxResponse, error)
	GetBalance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	GetTxByAddress(context.Context, *TxAddressRequest) (*TxAddressResponse, error)
	GetTxByHash(context.Context, *TxHashRequest) (*TxHashResponse, error)
	GetAccount(context.Context, *AccountRequest) (*AccountResponse, error)
	GetUtxo(context.Context, *UtxoRequest) (*UtxoResponse, error)
	GetMinRent(context.Context, *MinRentRequest) (*MinRentResponse, error)
}

// UnimplementedWalletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServiceServer struct {
}

func (UnimplementedWalletServiceServer) GetSupportCoins(context.Context, *SupportCoinsRequest) (*SupportCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportCoins not implemented")
}
func (UnimplementedWalletServiceServer) GetNonce(context.Context, *NonceRequest) (*NonceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNonce not implemented")
}
func (UnimplementedWalletServiceServer) GetGasPrice(context.Context, *GasPriceRequest) (*GasPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGasPrice not implemented")
}
func (UnimplementedWalletServiceServer) SendTx(context.Context, *SendTxRequest) (*SendTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTx not implemented")
}
func (UnimplementedWalletServiceServer) GetBalance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedWalletServiceServer) GetTxByAddress(context.Context, *TxAddressRequest) (*TxAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxByAddress not implemented")
}
func (UnimplementedWalletServiceServer) GetTxByHash(context.Context, *TxHashRequest) (*TxHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxByHash not implemented")
}
func (UnimplementedWalletServiceServer) GetAccount(context.Context, *AccountRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedWalletServiceServer) GetUtxo(context.Context, *UtxoRequest) (*UtxoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUtxo not implemented")
}
func (UnimplementedWalletServiceServer) GetMinRent(context.Context, *MinRentRequest) (*MinRentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinRent not implemented")
}
func (UnimplementedWalletServiceServer) mustEmbedUnimplementedWalletServiceServer() {}

// UnsafeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServiceServer will
// result in compilation errors.
type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	s.RegisterService(&WalletService_ServiceDesc, srv)
}

func _WalletService_GetSupportCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportCoinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetSupportCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.wallet.WalletService/getSupportCoins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetSupportCoins(ctx, req.(*SupportCoinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetNonce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetNonce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.wallet.WalletService/getNonce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetNonce(ctx, req.(*NonceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetGasPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GasPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetGasPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.wallet.WalletService/getGasPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetGasPrice(ctx, req.(*GasPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_SendTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).SendTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.wallet.WalletService/SendTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).SendTx(ctx, req.(*SendTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.wallet.WalletService/getBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetBalance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetTxByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetTxByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.wallet.WalletService/getTxByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetTxByAddress(ctx, req.(*TxAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetTxByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetTxByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.wallet.WalletService/getTxByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetTxByHash(ctx, req.(*TxHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.wallet.WalletService/getAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetUtxo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UtxoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetUtxo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.wallet.WalletService/getUtxo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetUtxo(ctx, req.(*UtxoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetMinRent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinRentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetMinRent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.wallet.WalletService/getMinRent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetMinRent(ctx, req.(*MinRentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "savourrpc.wallet.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSupportCoins",
			Handler:    _WalletService_GetSupportCoins_Handler,
		},
		{
			MethodName: "getNonce",
			Handler:    _WalletService_GetNonce_Handler,
		},
		{
			MethodName: "getGasPrice",
			Handler:    _WalletService_GetGasPrice_Handler,
		},
		{
			MethodName: "SendTx",
			Handler:    _WalletService_SendTx_Handler,
		},
		{
			MethodName: "getBalance",
			Handler:    _WalletService_GetBalance_Handler,
		},
		{
			MethodName: "getTxByAddress",
			Handler:    _WalletService_GetTxByAddress_Handler,
		},
		{
			MethodName: "getTxByHash",
			Handler:    _WalletService_GetTxByHash_Handler,
		},
		{
			MethodName: "getAccount",
			Handler:    _WalletService_GetAccount_Handler,
		},
		{
			MethodName: "getUtxo",
			Handler:    _WalletService_GetUtxo_Handler,
		},
		{
			MethodName: "getMinRent",
			Handler:    _WalletService_GetMinRent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "savourrpc/wallet.proto",
}
