package main

import (
	"billy/web/routes"
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"time"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(iris.Cache(1 * time.Minute))

	// 配置服务器
	app.ConfigureHost(func(h *iris.Supervisor) {
		h.RegisterOnShutdown(onShutDown)
	})

	// Gracefully Shutdown
	iris.RegisterOnInterrupt(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = app.Shutdown(ctx)
	})

	// 初始化路由
	routes.InitRouter(app)
	// 配置html引擎
	htmlEngine := iris.HTML("./web/template", ".html")
	// 开发模式下开启重载
	htmlEngine.Reload(true)

	app.RegisterView(htmlEngine)

	_ = app.Run(iris.Addr(":8080"), iris.WithoutInterruptHandler, iris.WithOptimizations)
}

// onShutDown
func onShutDown() {
	println("Server terminated!")
}
