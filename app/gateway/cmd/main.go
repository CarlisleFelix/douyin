package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"douyin/app/gateway/routes"
	"douyin/app/gateway/rpc"
	"douyin/config"
	"douyin/utils/shutdown"
)

func main() {
	config.InitConfig()
	rpc.Init()

	go startListen() // 转载路由
	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-osSignals
		fmt.Println("exit! ", s)
	}
	fmt.Println("gateway listen on :8001")
}

func startListen() {

	r := routes.NewRouter()
	server := &http.Server{
		Addr:           config.Conf.Server.Addr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("gateway启动失败, err: ", err)
	}
	go func() {
		// 优雅关闭
		shutdown.GracefullyShutdown(server)
	}()
}
