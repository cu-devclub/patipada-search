// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: mlGateway.proto

package proto

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

// MlGatewayServiceClient is the client API for MlGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MlGatewayServiceClient interface {
	Text2Vec(ctx context.Context, in *Text2VecRequest, opts ...grpc.CallOption) (*Text2VecResponse, error)
}

type mlGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMlGatewayServiceClient(cc grpc.ClientConnInterface) MlGatewayServiceClient {
	return &mlGatewayServiceClient{cc}
}

func (c *mlGatewayServiceClient) Text2Vec(ctx context.Context, in *Text2VecRequest, opts ...grpc.CallOption) (*Text2VecResponse, error) {
	out := new(Text2VecResponse)
	err := c.cc.Invoke(ctx, "/proto.MlGatewayService/Text2Vec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MlGatewayServiceServer is the server API for MlGatewayService service.
// All implementations must embed UnimplementedMlGatewayServiceServer
// for forward compatibility
type MlGatewayServiceServer interface {
	Text2Vec(context.Context, *Text2VecRequest) (*Text2VecResponse, error)
	mustEmbedUnimplementedMlGatewayServiceServer()
}

// UnimplementedMlGatewayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMlGatewayServiceServer struct {
}

func (UnimplementedMlGatewayServiceServer) Text2Vec(context.Context, *Text2VecRequest) (*Text2VecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Text2Vec not implemented")
}
func (UnimplementedMlGatewayServiceServer) mustEmbedUnimplementedMlGatewayServiceServer() {}

// UnsafeMlGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MlGatewayServiceServer will
// result in compilation errors.
type UnsafeMlGatewayServiceServer interface {
	mustEmbedUnimplementedMlGatewayServiceServer()
}

func RegisterMlGatewayServiceServer(s grpc.ServiceRegistrar, srv MlGatewayServiceServer) {
	s.RegisterService(&MlGatewayService_ServiceDesc, srv)
}

func _MlGatewayService_Text2Vec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text2VecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MlGatewayServiceServer).Text2Vec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MlGatewayService/Text2Vec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MlGatewayServiceServer).Text2Vec(ctx, req.(*Text2VecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MlGatewayService_ServiceDesc is the grpc.ServiceDesc for MlGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MlGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MlGatewayService",
	HandlerType: (*MlGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Text2Vec",
			Handler:    _MlGatewayService_Text2Vec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mlGateway.proto",
}