package ginHandler

import (
	"github.com/gin-gonic/gin"
	"log"
	"study.com/api-gateway/internal/serviceDiscovery"
	pb "study.com/study-user/pkg/pb.service"
)

func UserRegister(ginCtx *gin.Context) {
	etcdDiscovery := serviceDiscovery.EtcdCenter
	//获取用户服务
	conn, _ := etcdDiscovery.Discovery("study-user-service")
	userService := pb.NewUserServiceClient(conn)
	data, err := userService.SayHello(ginCtx, &pb.Request{})
	if err != nil {
		log.Printf("调用用户服务出错了,原因是: %v \n", err)
	}
	ginCtx.JSON(200, gin.H{
		"message": data,
	})
}
