package controller

import "github.com/kataras/iris/v12/mvc"

type IndexController struct{}

func (c *IndexController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done // 和你已知的任何标准 API  调用

	// 1-> 方法
	// 2-> 路径
	// 3-> 控制器函数的名称将被解析未一个处理程序 [ handler ]
	// 4-> 任何应该在 MyCustomHandler 之前运行的处理程序[ handlers ]
	b.Handle("GET", "/something/{id:long}", "IndexController")
}

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
