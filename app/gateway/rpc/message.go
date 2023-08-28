package rpc

import (
	"context"
	pb "douyin/idl/pb/message"
)

// 聊天记录
func MessageChat(ctx context.Context, req *pb.DouyinMessageChatRequest) (resp *pb.DouyinMessageChatResponse, err error) {
	r, err := MessageClient.MessageChat(ctx, req)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// 消息操作
func MessageAction(ctx context.Context, req *pb.DouyinMessageActionRequest) (resp *pb.DouyinMessageActionResponse, err error) {
	r, err := MessageClient.MessageAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return r, nil
}
