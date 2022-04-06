package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Get("/", func(ctx context.Context) {})

	app.RegisterView(iris.HTML("./web/view", ".html"))
	_ = app.Run(iris.Addr(":8080"))
}
