package main

import (
	"blog-server-go/models"
	"blog-server-go/pkg/gredis"
	"blog-server-go/pkg/logging"
	"blog-server-go/pkg/setting"
	"blog-server-go/pkg/util"
	"blog-server-go/routers"
	"fmt"

	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

func main() {

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout

	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	//endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	server := endless.NewServer(endPoint, routers.InitRouter())

	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

	//router := routers.InitRouter()
	//
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//s.ListenAndServe()

}
