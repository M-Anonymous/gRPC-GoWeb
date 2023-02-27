package ginRouterRegister

import (
	"log"
	app "study.com/api-gateway/internal/webAppFramework"
	_ "study.com/api-gateway/internal/webAppFramework/ginRouter"
)

func init() {
	log.Printf("Web服务正在启动... \n")
	app.GinWebApp.WebServe()
	log.Printf("Web服务启动完成! \n")
	defer app.GinWebApp.WebClose()
}
