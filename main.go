package main

import (
	"billy/config"
	"billy/web/routes"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"time"
)

func main() {
	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	// 初始化路由，需要置于缓存配置之前
	mvc.Configure(app, routes.InitRouter)

	// 配置html引擎
	htmlEngine := iris.HTML(config.TemplateDir(), config.TemplateExtension())
	app.RegisterView(htmlEngine)
	// 开发模式相关设置
	if config.IsDevMode() {
		// 开发模式下开启重载
		htmlEngine.Reload(true)
		// 设置日志级别为debug
		app.Logger().SetLevel("debug")
		// 禁用缓存
		app.Use(iris.NoCache)
	}

	// 静态文件处理
	app.HandleDir("static", config.ThemeStatic())
	app.HandleDir("admin/static", config.AdminStatic())
	app.Favicon(config.Favicon())

	// 置于开发模式相关配置之后，不然开发模式NoCache不生效
	app.Use(iris.StaticCache(24 * time.Hour))
	app.Use(iris.Cache(1 * time.Minute))

	_ = app.Run(iris.Addr(fmt.Sprintf(":%d", config.Port())), iris.WithoutInterruptHandler, iris.WithOptimizations)
}
