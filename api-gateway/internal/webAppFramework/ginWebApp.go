package webAppFramework

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var GinWebApp *ginWebApp

type ginWebApp struct {
	Router *gin.Engine
	Srv    *http.Server
}

func init() {
	GinWebApp = &ginWebApp{
		Router: gin.Default(),
	}
}

func (app *ginWebApp) WebServe() {
	srv := &http.Server{
		Addr:    viper.GetString("server.addr"),
		Handler: app.Router,
	}
	app.Srv = srv
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

}

func (app *ginWebApp) WebClose() {
	// 等待中断信号以优雅地关闭服务器（设置 3 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := app.Srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
