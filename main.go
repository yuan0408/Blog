package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/yuan0408/go-gin-example/pkg/setting"
	"github.com/yuan0408/go-gin-example/routers"
	"log"
	"syscall"
)

func main() {
	//给进程发送SIGTERM信号实现优雅的重启
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endpoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endpoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
