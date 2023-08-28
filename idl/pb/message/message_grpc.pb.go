// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: message.proto

package message

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
	MessageService_MessageChat_FullMethodName   = "/douyin.idl.chat.MessageService/MessageChat"
	MessageService_MessageAction_FullMethodName = "/douyin.idl.chat.MessageService/MessageAction"
)

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	// 聊天记录
	MessageChat(ctx context.Context, in *DouyinMessageChatRequest, opts ...grpc.CallOption) (*DouyinMessageChatResponse, error)
	// 消息操作
	MessageAction(ctx context.Context, in *DouyinMessageActionRequest, opts ...grpc.CallOption) (*DouyinMessageActionResponse, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) MessageChat(ctx context.Context, in *DouyinMessageChatRequest, opts ...grpc.CallOption) (*DouyinMessageChatResponse, error) {
	out := new(DouyinMessageChatResponse)
	err := c.cc.Invoke(ctx, MessageService_MessageChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) MessageAction(ctx context.Context, in *DouyinMessageActionRequest, opts ...grpc.CallOption) (*DouyinMessageActionResponse, error) {
	out := new(DouyinMessageActionResponse)
	err := c.cc.Invoke(ctx, MessageService_MessageAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	// 聊天记录
	MessageChat(context.Context, *DouyinMessageChatRequest) (*DouyinMessageChatResponse, error)
	// 消息操作
	MessageAction(context.Context, *DouyinMessageActionRequest) (*DouyinMessageActionResponse, error)
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) MessageChat(context.Context, *DouyinMessageChatRequest) (*DouyinMessageChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageChat not implemented")
}
func (UnimplementedMessageServiceServer) MessageAction(context.Context, *DouyinMessageActionRequest) (*DouyinMessageActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageAction not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_MessageChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinMessageChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).MessageChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_MessageChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).MessageChat(ctx, req.(*DouyinMessageChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_MessageAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinMessageActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).MessageAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageService_MessageAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).MessageAction(ctx, req.(*DouyinMessageActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin.idl.chat.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessageChat",
			Handler:    _MessageService_MessageChat_Handler,
		},
		{
			MethodName: "MessageAction",
			Handler:    _MessageService_MessageAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
