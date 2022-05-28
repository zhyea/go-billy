package controller

import (
	"github.com/kataras/iris/v12/mvc"
)

type AdminController struct {
}

//
// Get hello
func (c *AdminController) Get() interface{} {
	view := mvc.View{
		Name: "admin/index.html",
		Data: map[string]interface{}{
			"Title":     "Hello Page",
			"MyMessage": "Welcome to my awesome website",
		},
	}
	return view
}
