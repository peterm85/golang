// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package main

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

// ClientStreamServiceClient is the client API for ClientStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientStreamServiceClient interface {
	FetchResponse(ctx context.Context, opts ...grpc.CallOption) (ClientStreamService_FetchResponseClient, error)
}

type clientStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientStreamServiceClient(cc grpc.ClientConnInterface) ClientStreamServiceClient {
	return &clientStreamServiceClient{cc}
}

func (c *clientStreamServiceClient) FetchResponse(ctx context.Context, opts ...grpc.CallOption) (ClientStreamService_FetchResponseClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientStreamService_ServiceDesc.Streams[0], "/client.ClientStreamService/FetchResponse", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientStreamServiceFetchResponseClient{stream}
	return x, nil
}

type ClientStreamService_FetchResponseClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type clientStreamServiceFetchResponseClient struct {
	grpc.ClientStream
}

func (x *clientStreamServiceFetchResponseClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientStreamServiceFetchResponseClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStreamServiceServer is the server API for ClientStreamService service.
// All implementations must embed UnimplementedClientStreamServiceServer
// for forward compatibility
type ClientStreamServiceServer interface {
	FetchResponse(ClientStreamService_FetchResponseServer) error
	mustEmbedUnimplementedClientStreamServiceServer()
}

// UnimplementedClientStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientStreamServiceServer struct {
}

func (UnimplementedClientStreamServiceServer) FetchResponse(ClientStreamService_FetchResponseServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchResponse not implemented")
}
func (UnimplementedClientStreamServiceServer) mustEmbedUnimplementedClientStreamServiceServer() {}

// UnsafeClientStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientStreamServiceServer will
// result in compilation errors.
type UnsafeClientStreamServiceServer interface {
	mustEmbedUnimplementedClientStreamServiceServer()
}

func RegisterClientStreamServiceServer(s grpc.ServiceRegistrar, srv ClientStreamServiceServer) {
	s.RegisterService(&ClientStreamService_ServiceDesc, srv)
}

func _ClientStreamService_FetchResponse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientStreamServiceServer).FetchResponse(&clientStreamServiceFetchResponseServer{stream})
}

type ClientStreamService_FetchResponseServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type clientStreamServiceFetchResponseServer struct {
	grpc.ServerStream
}

func (x *clientStreamServiceFetchResponseServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientStreamServiceFetchResponseServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStreamService_ServiceDesc is the grpc.ServiceDesc for ClientStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "client.ClientStreamService",
	HandlerType: (*ClientStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchResponse",
			Handler:       _ClientStreamService_FetchResponse_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/client-stream/clientstr.proto",
}
