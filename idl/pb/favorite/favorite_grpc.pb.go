// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.1
// source: favorite.proto

package favorite

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
	FavoriteService_FavoriteAction_FullMethodName = "/douyin.idl.favorite.FavoriteService/FavoriteAction"
	FavoriteService_FavoriteList_FullMethodName   = "/douyin.idl.favorite.FavoriteService/FavoriteList"
)

// FavoriteServiceClient is the client API for FavoriteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavoriteServiceClient interface {
	FavoriteAction(ctx context.Context, in *DouyinFavoriteActionRequest, opts ...grpc.CallOption) (*DouyinFavoriteActionResponse, error)
	FavoriteList(ctx context.Context, in *DouyinFavoriteListRequest, opts ...grpc.CallOption) (*DouyinFavoriteListResponse, error)
}

type favoriteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteServiceClient(cc grpc.ClientConnInterface) FavoriteServiceClient {
	return &favoriteServiceClient{cc}
}

func (c *favoriteServiceClient) FavoriteAction(ctx context.Context, in *DouyinFavoriteActionRequest, opts ...grpc.CallOption) (*DouyinFavoriteActionResponse, error) {
	out := new(DouyinFavoriteActionResponse)
	err := c.cc.Invoke(ctx, FavoriteService_FavoriteAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteServiceClient) FavoriteList(ctx context.Context, in *DouyinFavoriteListRequest, opts ...grpc.CallOption) (*DouyinFavoriteListResponse, error) {
	out := new(DouyinFavoriteListResponse)
	err := c.cc.Invoke(ctx, FavoriteService_FavoriteList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteServiceServer is the server API for FavoriteService service.
// All implementations must embed UnimplementedFavoriteServiceServer
// for forward compatibility
type FavoriteServiceServer interface {
	FavoriteAction(context.Context, *DouyinFavoriteActionRequest) (*DouyinFavoriteActionResponse, error)
	FavoriteList(context.Context, *DouyinFavoriteListRequest) (*DouyinFavoriteListResponse, error)
	mustEmbedUnimplementedFavoriteServiceServer()
}

// UnimplementedFavoriteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFavoriteServiceServer struct {
}

func (UnimplementedFavoriteServiceServer) FavoriteAction(context.Context, *DouyinFavoriteActionRequest) (*DouyinFavoriteActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedFavoriteServiceServer) FavoriteList(context.Context, *DouyinFavoriteListRequest) (*DouyinFavoriteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}
func (UnimplementedFavoriteServiceServer) mustEmbedUnimplementedFavoriteServiceServer() {}

// UnsafeFavoriteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavoriteServiceServer will
// result in compilation errors.
type UnsafeFavoriteServiceServer interface {
	mustEmbedUnimplementedFavoriteServiceServer()
}

func RegisterFavoriteServiceServer(s grpc.ServiceRegistrar, srv FavoriteServiceServer) {
	s.RegisterService(&FavoriteService_ServiceDesc, srv)
}

func _FavoriteService_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFavoriteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteService_FavoriteAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).FavoriteAction(ctx, req.(*DouyinFavoriteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteService_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFavoriteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteService_FavoriteList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).FavoriteList(ctx, req.(*DouyinFavoriteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FavoriteService_ServiceDesc is the grpc.ServiceDesc for FavoriteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FavoriteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin.idl.favorite.FavoriteService",
	HandlerType: (*FavoriteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FavoriteAction",
			Handler:    _FavoriteService_FavoriteAction_Handler,
		},
		{
			MethodName: "FavoriteList",
			Handler:    _FavoriteService_FavoriteList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "favorite.proto",
}
