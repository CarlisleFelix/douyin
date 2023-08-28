// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: user.proto

package user

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
	UserService_UserRegister_FullMethodName        = "/douyin.idl.user.UserService/UserRegister"
	UserService_UserLogin_FullMethodName           = "/douyin.idl.user.UserService/UserLogin"
	UserService_GetUserInfo_FullMethodName         = "/douyin.idl.user.UserService/GetUserInfo"
	UserService_UpdateTotalFavorite_FullMethodName = "/douyin.idl.user.UserService/UpdateTotalFavorite"
	UserService_UpdateFavoriteCount_FullMethodName = "/douyin.idl.user.UserService/UpdateFavoriteCount"
	UserService_UpdateFollowCount_FullMethodName   = "/douyin.idl.user.UserService/UpdateFollowCount"
	UserService_UpdateFollowerCount_FullMethodName = "/douyin.idl.user.UserService/UpdateFollowerCount"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// 注册
	UserRegister(ctx context.Context, in *DouyinUserRegisterRequest, opts ...grpc.CallOption) (*DouyinUserRegisterResponse, error)
	// 登录
	UserLogin(ctx context.Context, in *DouyinUserLoginRequest, opts ...grpc.CallOption) (*DouyinUserLoginResponse, error)
	// 获取用户信息
	GetUserInfo(ctx context.Context, in *DouyinUserRequest, opts ...grpc.CallOption) (*DouyinUserResponse, error)
	// 更新用户获赞数
	UpdateTotalFavorite(ctx context.Context, in *DouyinTotalFavoriteRequest, opts ...grpc.CallOption) (*DouyinTotalFavoriteResponse, error)
	// 更新用户点赞数
	UpdateFavoriteCount(ctx context.Context, in *DouyinFavoriteCountRequest, opts ...grpc.CallOption) (*DouyinFavoriteCountResponse, error)
	// 更新用户关注数
	UpdateFollowCount(ctx context.Context, in *DouyinFollowCountRequest, opts ...grpc.CallOption) (*DouyinFollowCountResponse, error)
	// 更新用户粉丝数
	UpdateFollowerCount(ctx context.Context, in *DouyinFollowerCountRequest, opts ...grpc.CallOption) (*DouyinFollowerCountResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserRegister(ctx context.Context, in *DouyinUserRegisterRequest, opts ...grpc.CallOption) (*DouyinUserRegisterResponse, error) {
	out := new(DouyinUserRegisterResponse)
	err := c.cc.Invoke(ctx, UserService_UserRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserLogin(ctx context.Context, in *DouyinUserLoginRequest, opts ...grpc.CallOption) (*DouyinUserLoginResponse, error) {
	out := new(DouyinUserLoginResponse)
	err := c.cc.Invoke(ctx, UserService_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfo(ctx context.Context, in *DouyinUserRequest, opts ...grpc.CallOption) (*DouyinUserResponse, error) {
	out := new(DouyinUserResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateTotalFavorite(ctx context.Context, in *DouyinTotalFavoriteRequest, opts ...grpc.CallOption) (*DouyinTotalFavoriteResponse, error) {
	out := new(DouyinTotalFavoriteResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateTotalFavorite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateFavoriteCount(ctx context.Context, in *DouyinFavoriteCountRequest, opts ...grpc.CallOption) (*DouyinFavoriteCountResponse, error) {
	out := new(DouyinFavoriteCountResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateFavoriteCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateFollowCount(ctx context.Context, in *DouyinFollowCountRequest, opts ...grpc.CallOption) (*DouyinFollowCountResponse, error) {
	out := new(DouyinFollowCountResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateFollowCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateFollowerCount(ctx context.Context, in *DouyinFollowerCountRequest, opts ...grpc.CallOption) (*DouyinFollowerCountResponse, error) {
	out := new(DouyinFollowerCountResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateFollowerCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the global API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// 注册
	UserRegister(context.Context, *DouyinUserRegisterRequest) (*DouyinUserRegisterResponse, error)
	// 登录
	UserLogin(context.Context, *DouyinUserLoginRequest) (*DouyinUserLoginResponse, error)
	// 获取用户信息
	GetUserInfo(context.Context, *DouyinUserRequest) (*DouyinUserResponse, error)
	// 更新用户获赞数
	UpdateTotalFavorite(context.Context, *DouyinTotalFavoriteRequest) (*DouyinTotalFavoriteResponse, error)
	// 更新用户点赞数
	UpdateFavoriteCount(context.Context, *DouyinFavoriteCountRequest) (*DouyinFavoriteCountResponse, error)
	// 更新用户关注数
	UpdateFollowCount(context.Context, *DouyinFollowCountRequest) (*DouyinFollowCountResponse, error)
	// 更新用户粉丝数
	UpdateFollowerCount(context.Context, *DouyinFollowerCountRequest) (*DouyinFollowerCountResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserRegister(context.Context, *DouyinUserRegisterRequest) (*DouyinUserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedUserServiceServer) UserLogin(context.Context, *DouyinUserLoginRequest) (*DouyinUserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfo(context.Context, *DouyinUserRequest) (*DouyinUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserServiceServer) UpdateTotalFavorite(context.Context, *DouyinTotalFavoriteRequest) (*DouyinTotalFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTotalFavorite not implemented")
}
func (UnimplementedUserServiceServer) UpdateFavoriteCount(context.Context, *DouyinFavoriteCountRequest) (*DouyinFavoriteCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFavoriteCount not implemented")
}
func (UnimplementedUserServiceServer) UpdateFollowCount(context.Context, *DouyinFollowCountRequest) (*DouyinFollowCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFollowCount not implemented")
}
func (UnimplementedUserServiceServer) UpdateFollowerCount(context.Context, *DouyinFollowerCountRequest) (*DouyinFollowerCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFollowerCount not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRegister(ctx, req.(*DouyinUserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserLogin(ctx, req.(*DouyinUserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfo(ctx, req.(*DouyinUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateTotalFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinTotalFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateTotalFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateTotalFavorite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateTotalFavorite(ctx, req.(*DouyinTotalFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateFavoriteCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFavoriteCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateFavoriteCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateFavoriteCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateFavoriteCount(ctx, req.(*DouyinFavoriteCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateFollowCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFollowCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateFollowCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateFollowCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateFollowCount(ctx, req.(*DouyinFollowCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateFollowerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFollowerCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateFollowerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateFollowerCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateFollowerCount(ctx, req.(*DouyinFollowerCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin.idl.user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _UserService_UserRegister_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _UserService_UserLogin_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserService_GetUserInfo_Handler,
		},
		{
			MethodName: "UpdateTotalFavorite",
			Handler:    _UserService_UpdateTotalFavorite_Handler,
		},
		{
			MethodName: "UpdateFavoriteCount",
			Handler:    _UserService_UpdateFavoriteCount_Handler,
		},
		{
			MethodName: "UpdateFollowCount",
			Handler:    _UserService_UpdateFollowCount_Handler,
		},
		{
			MethodName: "UpdateFollowerCount",
			Handler:    _UserService_UpdateFollowerCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
