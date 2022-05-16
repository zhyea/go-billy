package controller

import (
	"billy/model"
	"billy/service"
)

type UserController struct {
	userService service.UserService
}

//
// GetBy 返回用户信息
func (c *UserController) GetBy(id int64) *model.User {
	return c.userService.GetById(id)
}
