package ginRouter

import (
	"github.com/gin-gonic/gin"
	"log"
	"study.com/api-gateway/internal/webAppFramework"
	"study.com/api-gateway/internal/webAppFramework/ginHandler"
)

func init() {
	log.Printf("正在注册用户路由...")
	ginRouter := webAppFramework.GinWebApp.Router
	RegisterUserRouter(ginRouter)
	log.Printf("用户路由注册完成! \n")
}

func RegisterUserRouter(router *gin.Engine) {
	router.GET("/hello", ginHandler.UserRegister)
}
