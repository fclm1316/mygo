package main

import (
	"allinone/controller"
	// "allinone/middleware"
	"allinone/tool"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tool.ParseConfig()
	if err != nil {
		log.Fatalf("解析失败 ", err)
	}
	// 初始化数据库
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		log.Fatalln(err)
		return
	}

	tool.InitGlobalVariable()

	gin.SetMode(cfg.AppInfo.AppMode)
	gin.DisableConsoleColor()

	app := gin.Default()
	// 中间件
	// app.Use(middleware.CostTime())

	registerRouter(app)

	// app.Run(cfg.AppInfo.AppHost + ":" + cfg.AppInfo.AppPort)
	AddrPort := fmt.Sprintf("%s:%s", cfg.AppInfo.AppHost, cfg.AppInfo.AppPort)
	Srv := &http.Server{
		Addr:    AddrPort,
		Handler: app,
	}

	go func() {
		if err := Srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	tool.DbEngine.Engine.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := Srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)

	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 1 seconds.")
	}
	log.Println("Server exiting")
}

// 注册路由
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.SysInfoController).Router(router)
	new(controller.SysEnvController).Router(router)
	new(controller.SysNameController).Router(router)
	new(controller.FileListController).Router(router)
	new(controller.SysUserPasswdController).Router(router)
	new(controller.FileInfoController).Router(router)
	new(controller.SysBatchController).Router(router)
	new(controller.FileResultController).Router(router)
}
