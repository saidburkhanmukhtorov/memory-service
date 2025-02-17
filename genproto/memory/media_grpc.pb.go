// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: submodule-for-timecapsule/memory_service/media.proto

package memory

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MediaService_GetMediaById_FullMethodName = "/memory.MediaService/GetMediaById"
	MediaService_DeleteMedia_FullMethodName  = "/memory.MediaService/DeleteMedia"
	MediaService_GetAllMedia_FullMethodName  = "/memory.MediaService/GetAllMedia"
)

// MediaServiceClient is the client API for MediaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediaServiceClient interface {
	GetMediaById(ctx context.Context, in *GetMediaByIdRequest, opts ...grpc.CallOption) (*Media, error)
	DeleteMedia(ctx context.Context, in *DeleteMediaRequest, opts ...grpc.CallOption) (*DeleteMediaResponse, error)
	GetAllMedia(ctx context.Context, in *GetAllMediaRequest, opts ...grpc.CallOption) (*GetAllMediaResponse, error)
}

type mediaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaServiceClient(cc grpc.ClientConnInterface) MediaServiceClient {
	return &mediaServiceClient{cc}
}

func (c *mediaServiceClient) GetMediaById(ctx context.Context, in *GetMediaByIdRequest, opts ...grpc.CallOption) (*Media, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Media)
	err := c.cc.Invoke(ctx, MediaService_GetMediaById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) DeleteMedia(ctx context.Context, in *DeleteMediaRequest, opts ...grpc.CallOption) (*DeleteMediaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMediaResponse)
	err := c.cc.Invoke(ctx, MediaService_DeleteMedia_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) GetAllMedia(ctx context.Context, in *GetAllMediaRequest, opts ...grpc.CallOption) (*GetAllMediaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllMediaResponse)
	err := c.cc.Invoke(ctx, MediaService_GetAllMedia_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaServiceServer is the server API for MediaService service.
// All implementations must embed UnimplementedMediaServiceServer
// for forward compatibility
type MediaServiceServer interface {
	GetMediaById(context.Context, *GetMediaByIdRequest) (*Media, error)
	DeleteMedia(context.Context, *DeleteMediaRequest) (*DeleteMediaResponse, error)
	GetAllMedia(context.Context, *GetAllMediaRequest) (*GetAllMediaResponse, error)
	mustEmbedUnimplementedMediaServiceServer()
}

// UnimplementedMediaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMediaServiceServer struct {
}

func (UnimplementedMediaServiceServer) GetMediaById(context.Context, *GetMediaByIdRequest) (*Media, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMediaById not implemented")
}
func (UnimplementedMediaServiceServer) DeleteMedia(context.Context, *DeleteMediaRequest) (*DeleteMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMedia not implemented")
}
func (UnimplementedMediaServiceServer) GetAllMedia(context.Context, *GetAllMediaRequest) (*GetAllMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMedia not implemented")
}
func (UnimplementedMediaServiceServer) mustEmbedUnimplementedMediaServiceServer() {}

// UnsafeMediaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediaServiceServer will
// result in compilation errors.
type UnsafeMediaServiceServer interface {
	mustEmbedUnimplementedMediaServiceServer()
}

func RegisterMediaServiceServer(s grpc.ServiceRegistrar, srv MediaServiceServer) {
	s.RegisterService(&MediaService_ServiceDesc, srv)
}

func _MediaService_GetMediaById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMediaByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).GetMediaById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediaService_GetMediaById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).GetMediaById(ctx, req.(*GetMediaByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_DeleteMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).DeleteMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediaService_DeleteMedia_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).DeleteMedia(ctx, req.(*DeleteMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_GetAllMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).GetAllMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MediaService_GetAllMedia_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).GetAllMedia(ctx, req.(*GetAllMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MediaService_ServiceDesc is the grpc.ServiceDesc for MediaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MediaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "memory.MediaService",
	HandlerType: (*MediaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMediaById",
			Handler:    _MediaService_GetMediaById_Handler,
		},
		{
			MethodName: "DeleteMedia",
			Handler:    _MediaService_DeleteMedia_Handler,
		},
		{
			MethodName: "GetAllMedia",
			Handler:    _MediaService_GetAllMedia_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "submodule-for-timecapsule/memory_service/media.proto",
}
