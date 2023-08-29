package rpc

import (
	"context"
	"douyin/idl/pb/comment"
	"douyin/idl/pb/message"
	"douyin/idl/pb/relation"
	"douyin/idl/pb/video"
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"douyin/config"
	"douyin/idl/pb/favorite"
	"douyin/idl/pb/user"

	"douyin/discovery"
)

var (
	Register       *discovery.Resolver
	ctx            context.Context
	CancelFunc     context.CancelFunc
	MessageClient  message.MessageServiceClient
	CommentClient  comment.CommentServiceClient
	FavoriteClient favorite.FavoriteServiceClient
	RelationClient relation.RelationServiceClient
	UserClient     user.UserServiceClient
	VideoClient    video.VideoServiceClient
)

func Init() {
	Register = discovery.NewResolver([]string{config.Conf.Etcd.Address}, logrus.New())
	resolver.Register(Register)
	ctx, CancelFunc = context.WithTimeout(context.Background(), 3*time.Second)

	defer Register.Close()
	initClient(config.Conf.Domain["favorite"].Name, &FavoriteClient)
	initClient(config.Conf.Domain["user"].Name, &UserClient)
	initClient(config.Conf.Domain["video"].Name, &VideoClient)
	//initClient(config.Conf.Domain["comment"].Name, &CommentClient)
	//initClient(config.Conf.Domain["message"].Name, &MessageClient)
	//initClient(config.Conf.Domain["relation"].Name, &RelationClient)
}

func initClient(serviceName string, client interface{}) {
	conn, err := connectServer(serviceName)

	if err != nil {
		panic(err)
	}

	switch c := client.(type) {
	case *favorite.FavoriteServiceClient:
		*c = favorite.NewFavoriteServiceClient(conn)
	case *user.UserServiceClient:
		*c = user.NewUserServiceClient(conn)
	case *comment.CommentServiceClient:
		*c = comment.NewCommentServiceClient(conn)
	//case *message.MessageServiceServer:
	//	*c = message.NewMessageServiceClient(conn)
	case *relation.RelationServiceClient:
		*c = relation.NewRelationServiceClient(conn)
	case *video.VideoServiceClient:
		*c = video.NewVideoServiceClient(conn)
	default:
		panic("unsupported client type")
	}
}

func connectServer(serviceName string) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	addr := fmt.Sprintf("%s:///%s", Register.Scheme(), serviceName)

	// Load balance
	if config.Conf.Services[serviceName].LoadBalance {
		log.Printf("load balance enabled for %s\n", serviceName)
		opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, "round_robin")))
	}

	conn, err = grpc.DialContext(ctx, addr, opts...)
	return
}
