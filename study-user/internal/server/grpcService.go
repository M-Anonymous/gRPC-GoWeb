package server

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

var GrpcServer *grpc.Server

var lis net.Listener

func init() {
	GrpcServer = grpc.NewServer()
	//监听服务
	log.Printf("开始监听服务! \n")
	listen, err := net.Listen("tcp", viper.GetString("server.addr"))
	lis = listen
	if err != nil {
		log.Printf("监听服务失败,失败原因是: %v", err)
		panic(err)
	}
	log.Printf("监听服务中... \n")
}

func Serve() {
	GrpcServer.Serve(lis)
}
