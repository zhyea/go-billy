package controller

import "github.com/kataras/iris/v12/mvc"

type IndexController struct{}

// Get index
func (c *IndexController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Hello World!</h1>",
	}
}

// GetPing ping
func (c *IndexController) GetPing() string {
	return "pong!"
}

// GetHello hello
func (c *IndexController) GetHello() interface{} {
	return map[string]string{"message": "Hello World!"}
}
