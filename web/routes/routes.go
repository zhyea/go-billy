package routes

import (
	"billy/web/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(app *iris.Application) {
	mvc.New(app).Handle(new(controller.IndexController))
}
