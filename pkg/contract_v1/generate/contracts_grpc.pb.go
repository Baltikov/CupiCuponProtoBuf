// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: contracts.proto

package contracts_v1

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

// ContractsClient is the client API for Contracts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContractsClient interface {
	CreateContract(ctx context.Context, in *CreateContractRequest, opts ...grpc.CallOption) (*CreateContractResponse, error)
	GetContract(ctx context.Context, in *GetContractRequest, opts ...grpc.CallOption) (*GetContractResponse, error)
}

type contractsClient struct {
	cc grpc.ClientConnInterface
}

func NewContractsClient(cc grpc.ClientConnInterface) ContractsClient {
	return &contractsClient{cc}
}

func (c *contractsClient) CreateContract(ctx context.Context, in *CreateContractRequest, opts ...grpc.CallOption) (*CreateContractResponse, error) {
	out := new(CreateContractResponse)
	err := c.cc.Invoke(ctx, "/Contracts/CreateContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractsClient) GetContract(ctx context.Context, in *GetContractRequest, opts ...grpc.CallOption) (*GetContractResponse, error) {
	out := new(GetContractResponse)
	err := c.cc.Invoke(ctx, "/Contracts/GetContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContractsServer is the server API for Contracts service.
// All implementations must embed UnimplementedContractsServer
// for forward compatibility
type ContractsServer interface {
	CreateContract(context.Context, *CreateContractRequest) (*CreateContractResponse, error)
	GetContract(context.Context, *GetContractRequest) (*GetContractResponse, error)
	mustEmbedUnimplementedContractsServer()
}

// UnimplementedContractsServer must be embedded to have forward compatible implementations.
type UnimplementedContractsServer struct {
}

func (UnimplementedContractsServer) CreateContract(context.Context, *CreateContractRequest) (*CreateContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContract not implemented")
}
func (UnimplementedContractsServer) GetContract(context.Context, *GetContractRequest) (*GetContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContract not implemented")
}
func (UnimplementedContractsServer) mustEmbedUnimplementedContractsServer() {}

// UnsafeContractsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContractsServer will
// result in compilation errors.
type UnsafeContractsServer interface {
	mustEmbedUnimplementedContractsServer()
}

func RegisterContractsServer(s grpc.ServiceRegistrar, srv ContractsServer) {
	s.RegisterService(&Contracts_ServiceDesc, srv)
}

func _Contracts_CreateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractsServer).CreateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Contracts/CreateContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractsServer).CreateContract(ctx, req.(*CreateContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contracts_GetContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractsServer).GetContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Contracts/GetContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractsServer).GetContract(ctx, req.(*GetContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Contracts_ServiceDesc is the grpc.ServiceDesc for Contracts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Contracts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Contracts",
	HandlerType: (*ContractsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContract",
			Handler:    _Contracts_CreateContract_Handler,
		},
		{
			MethodName: "GetContract",
			Handler:    _Contracts_GetContract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contracts.proto",
}
