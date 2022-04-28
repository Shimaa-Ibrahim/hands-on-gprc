// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package primefactorizationpb

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

// PrimeFactorizationServiceClient is the client API for PrimeFactorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrimeFactorizationServiceClient interface {
	PrimeFactorization(ctx context.Context, in *PrimeFactorizationRequest, opts ...grpc.CallOption) (PrimeFactorizationService_PrimeFactorizationClient, error)
}

type primeFactorizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrimeFactorizationServiceClient(cc grpc.ClientConnInterface) PrimeFactorizationServiceClient {
	return &primeFactorizationServiceClient{cc}
}

func (c *primeFactorizationServiceClient) PrimeFactorization(ctx context.Context, in *PrimeFactorizationRequest, opts ...grpc.CallOption) (PrimeFactorizationService_PrimeFactorizationClient, error) {
	stream, err := c.cc.NewStream(ctx, &PrimeFactorizationService_ServiceDesc.Streams[0], "/primefactorization.PrimeFactorizationService/PrimeFactorization", opts...)
	if err != nil {
		return nil, err
	}
	x := &primeFactorizationServicePrimeFactorizationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrimeFactorizationService_PrimeFactorizationClient interface {
	Recv() (*PrimeFactorizationResponse, error)
	grpc.ClientStream
}

type primeFactorizationServicePrimeFactorizationClient struct {
	grpc.ClientStream
}

func (x *primeFactorizationServicePrimeFactorizationClient) Recv() (*PrimeFactorizationResponse, error) {
	m := new(PrimeFactorizationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrimeFactorizationServiceServer is the server API for PrimeFactorizationService service.
// All implementations must embed UnimplementedPrimeFactorizationServiceServer
// for forward compatibility
type PrimeFactorizationServiceServer interface {
	PrimeFactorization(*PrimeFactorizationRequest, PrimeFactorizationService_PrimeFactorizationServer) error
	mustEmbedUnimplementedPrimeFactorizationServiceServer()
}

// UnimplementedPrimeFactorizationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrimeFactorizationServiceServer struct {
}

func (UnimplementedPrimeFactorizationServiceServer) PrimeFactorization(*PrimeFactorizationRequest, PrimeFactorizationService_PrimeFactorizationServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeFactorization not implemented")
}
func (UnimplementedPrimeFactorizationServiceServer) mustEmbedUnimplementedPrimeFactorizationServiceServer() {
}

// UnsafePrimeFactorizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrimeFactorizationServiceServer will
// result in compilation errors.
type UnsafePrimeFactorizationServiceServer interface {
	mustEmbedUnimplementedPrimeFactorizationServiceServer()
}

func RegisterPrimeFactorizationServiceServer(s grpc.ServiceRegistrar, srv PrimeFactorizationServiceServer) {
	s.RegisterService(&PrimeFactorizationService_ServiceDesc, srv)
}

func _PrimeFactorizationService_PrimeFactorization_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeFactorizationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrimeFactorizationServiceServer).PrimeFactorization(m, &primeFactorizationServicePrimeFactorizationServer{stream})
}

type PrimeFactorizationService_PrimeFactorizationServer interface {
	Send(*PrimeFactorizationResponse) error
	grpc.ServerStream
}

type primeFactorizationServicePrimeFactorizationServer struct {
	grpc.ServerStream
}

func (x *primeFactorizationServicePrimeFactorizationServer) Send(m *PrimeFactorizationResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PrimeFactorizationService_ServiceDesc is the grpc.ServiceDesc for PrimeFactorizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrimeFactorizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "primefactorization.PrimeFactorizationService",
	HandlerType: (*PrimeFactorizationServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeFactorization",
			Handler:       _PrimeFactorizationService_PrimeFactorization_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stream-server-api/primefactorizationpb/prime-factorization.proto",
}