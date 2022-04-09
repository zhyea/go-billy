package main

import (
	"bison/web/controller"
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"time"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(iris.Cache(1 * time.Minute))

	mvc.New(app).Handle(new(controller.IndexController))

	// 配置服务器
	app.ConfigureHost(func(h *iris.Supervisor) {
		h.RegisterOnShutdown(onShutDown)
	})

	// Gracefully Shutdown
	iris.RegisterOnInterrupt(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// 关闭所有主机
		_ = app.Shutdown(ctx)
	})

	app.RegisterView(iris.HTML("./web/view", ".html"))

	_ = app.Run(iris.Addr(":8080"), iris.WithoutInterruptHandler, iris.WithConfiguration(iris.YAML("./configs/app.yml")))
}

// onShutDown
func onShutDown() {
	println("Server terminated!")
}
