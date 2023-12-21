// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: encoder/encoder.proto

package encoder

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

// EncoderClient is the client API for Encoder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EncoderClient interface {
	EncodeBlob(ctx context.Context, in *EncodeBlobRequest, opts ...grpc.CallOption) (*EncodeBlobReply, error)
}

type encoderClient struct {
	cc grpc.ClientConnInterface
}

func NewEncoderClient(cc grpc.ClientConnInterface) EncoderClient {
	return &encoderClient{cc}
}

func (c *encoderClient) EncodeBlob(ctx context.Context, in *EncodeBlobRequest, opts ...grpc.CallOption) (*EncodeBlobReply, error) {
	out := new(EncodeBlobReply)
	err := c.cc.Invoke(ctx, "/encoder.Encoder/EncodeBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EncoderServer is the server API for Encoder service.
// All implementations must embed UnimplementedEncoderServer
// for forward compatibility
type EncoderServer interface {
	EncodeBlob(context.Context, *EncodeBlobRequest) (*EncodeBlobReply, error)
	mustEmbedUnimplementedEncoderServer()
}

// UnimplementedEncoderServer must be embedded to have forward compatible implementations.
type UnimplementedEncoderServer struct {
}

func (UnimplementedEncoderServer) EncodeBlob(context.Context, *EncodeBlobRequest) (*EncodeBlobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncodeBlob not implemented")
}
func (UnimplementedEncoderServer) mustEmbedUnimplementedEncoderServer() {}

// UnsafeEncoderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EncoderServer will
// result in compilation errors.
type UnsafeEncoderServer interface {
	mustEmbedUnimplementedEncoderServer()
}

func RegisterEncoderServer(s grpc.ServiceRegistrar, srv EncoderServer) {
	s.RegisterService(&Encoder_ServiceDesc, srv)
}

func _Encoder_EncodeBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncodeBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncoderServer).EncodeBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/encoder.Encoder/EncodeBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncoderServer).EncodeBlob(ctx, req.(*EncodeBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Encoder_ServiceDesc is the grpc.ServiceDesc for Encoder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Encoder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "encoder.Encoder",
	HandlerType: (*EncoderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EncodeBlob",
			Handler:    _Encoder_EncodeBlob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "encoder/encoder.proto",
}
