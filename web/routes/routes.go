package routes

import (
	"billy/web/controller"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(context *mvc.Application) {
	// Index
	context.Handle(new(controller.IndexController))
}
