package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/v12/_examples/mvc/overview/web/controllers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Get("/", func(ctx *context.Context) {})

	app.RegisterView(iris.HTML("./web/view", ".html"))
	//注册控制器
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))
	_ = app.Run(iris.Addr(":8080"))
}
