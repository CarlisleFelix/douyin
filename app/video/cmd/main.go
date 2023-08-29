package main

import (
	"douyin/app/gateway/rpc"
	"douyin/app/video/internal/controller"
	"douyin/app/video/internal/dal/dao"
	"douyin/config"
	"douyin/discovery"
	pb "douyin/idl/pb/video"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	config.InitConfig()
	rpc.Init()
	dao.InitDB()
	// etcd 地址
	etcdAddress := []string{config.Conf.Etcd.Address}
	// 服务注册
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	grpcAddress := config.Conf.Services["video"].Addr[0]
	// 程序结束时注销
	defer etcdRegister.Stop()
	userNode := discovery.Server{
		Name: config.Conf.Domain["video"].Name,
		Addr: grpcAddress,
	}
	server := grpc.NewServer()
	defer server.Stop()
	// 绑定service
	pb.RegisterVideoServiceServer(server, controller.GetVideoSrv())
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err := etcdRegister.Register(userNode, 10); err != nil {
		panic(fmt.Sprintf("start global failed, err: %v", err))
	}
	logrus.Info("global started listen on ", grpcAddress)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
