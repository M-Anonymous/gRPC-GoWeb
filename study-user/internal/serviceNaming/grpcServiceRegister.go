package serviceNaming

import (
	"log"
	"study.com/study-user/internal/server"
	_ "study.com/study-user/internal/service"
)

func init() {
	log.Printf("正在服务...")
	server.Serve()
}
