// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: service.proto

package __

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

const (
	FactorService_Factor_FullMethodName = "/factors.FactorService/Factor"
)

// FactorServiceClient is the client API for FactorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FactorServiceClient interface {
	Factor(ctx context.Context, in *FactorReq, opts ...grpc.CallOption) (FactorService_FactorClient, error)
}

type factorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFactorServiceClient(cc grpc.ClientConnInterface) FactorServiceClient {
	return &factorServiceClient{cc}
}

func (c *factorServiceClient) Factor(ctx context.Context, in *FactorReq, opts ...grpc.CallOption) (FactorService_FactorClient, error) {
	stream, err := c.cc.NewStream(ctx, &FactorService_ServiceDesc.Streams[0], FactorService_Factor_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &factorServiceFactorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FactorService_FactorClient interface {
	Recv() (*FactorResp, error)
	grpc.ClientStream
}

type factorServiceFactorClient struct {
	grpc.ClientStream
}

func (x *factorServiceFactorClient) Recv() (*FactorResp, error) {
	m := new(FactorResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FactorServiceServer is the server API for FactorService service.
// All implementations must embed UnimplementedFactorServiceServer
// for forward compatibility
type FactorServiceServer interface {
	Factor(*FactorReq, FactorService_FactorServer) error
	mustEmbedUnimplementedFactorServiceServer()
}

// UnimplementedFactorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFactorServiceServer struct {
}

func (UnimplementedFactorServiceServer) Factor(*FactorReq, FactorService_FactorServer) error {
	return status.Errorf(codes.Unimplemented, "method Factor not implemented")
}
func (UnimplementedFactorServiceServer) mustEmbedUnimplementedFactorServiceServer() {}

// UnsafeFactorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FactorServiceServer will
// result in compilation errors.
type UnsafeFactorServiceServer interface {
	mustEmbedUnimplementedFactorServiceServer()
}

func RegisterFactorServiceServer(s grpc.ServiceRegistrar, srv FactorServiceServer) {
	s.RegisterService(&FactorService_ServiceDesc, srv)
}

func _FactorService_Factor_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FactorReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FactorServiceServer).Factor(m, &factorServiceFactorServer{stream})
}

type FactorService_FactorServer interface {
	Send(*FactorResp) error
	grpc.ServerStream
}

type factorServiceFactorServer struct {
	grpc.ServerStream
}

func (x *factorServiceFactorServer) Send(m *FactorResp) error {
	return x.ServerStream.SendMsg(m)
}

// FactorService_ServiceDesc is the grpc.ServiceDesc for FactorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FactorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "factors.FactorService",
	HandlerType: (*FactorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Factor",
			Handler:       _FactorService_Factor_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}