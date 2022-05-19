// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: financial-tracker/service.proto

package fintracker

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

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	// AddTransaction adds a new transaction.
	AddTransaction(ctx context.Context, in *AddTransactionRequest, opts ...grpc.CallOption) (*AddTransactionResponse, error)
	// FetchTransaction fetches a transaction.
	FetchTransaction(ctx context.Context, in *FetchTransactionRequest, opts ...grpc.CallOption) (*FetchTransactionResponse, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) AddTransaction(ctx context.Context, in *AddTransactionRequest, opts ...grpc.CallOption) (*AddTransactionResponse, error) {
	out := new(AddTransactionResponse)
	err := c.cc.Invoke(ctx, "/fintracker.v1.TransactionService/AddTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) FetchTransaction(ctx context.Context, in *FetchTransactionRequest, opts ...grpc.CallOption) (*FetchTransactionResponse, error) {
	out := new(FetchTransactionResponse)
	err := c.cc.Invoke(ctx, "/fintracker.v1.TransactionService/FetchTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations must embed UnimplementedTransactionServiceServer
// for forward compatibility
type TransactionServiceServer interface {
	// AddTransaction adds a new transaction.
	AddTransaction(context.Context, *AddTransactionRequest) (*AddTransactionResponse, error)
	// FetchTransaction fetches a transaction.
	FetchTransaction(context.Context, *FetchTransactionRequest) (*FetchTransactionResponse, error)
	mustEmbedUnimplementedTransactionServiceServer()
}

// UnimplementedTransactionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (UnimplementedTransactionServiceServer) AddTransaction(context.Context, *AddTransactionRequest) (*AddTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) FetchTransaction(context.Context, *FetchTransactionRequest) (*FetchTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_AddTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).AddTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fintracker.v1.TransactionService/AddTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).AddTransaction(ctx, req.(*AddTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_FetchTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).FetchTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fintracker.v1.TransactionService/FetchTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).FetchTransaction(ctx, req.(*FetchTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fintracker.v1.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTransaction",
			Handler:    _TransactionService_AddTransaction_Handler,
		},
		{
			MethodName: "FetchTransaction",
			Handler:    _TransactionService_FetchTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "financial-tracker/service.proto",
}
