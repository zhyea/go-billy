package routes

import (
	"billy/service"
	"billy/web/controller"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(app *mvc.Application) {
	// Index
	app.Handle(new(controller.IndexController))
	// User
	app.Party("/user").Register(service.NewUserService()).Handle(new(controller.UserController))
	// User
	app.Party("/admin").Handle(new(controller.AdminController))
}
