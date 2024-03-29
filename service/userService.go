package service

import (
	"billy/model"
	"billy/repo"
	"billy/tool/codec"
)

type UserService struct {
	repo repo.UserRepo
}

// NewUserService 返回默认的 UserService
func NewUserService() *UserService {
	return &UserService{
		repo: repo.UserRepo{},
	}
}

// CheckLogin 检查用户登录信息
func (service *UserService) CheckLogin(username string, password string) *model.User {
	password = codec.MD5(password)
	return service.repo.CheckAndGet(username, password)
}

// GetById 根据ID获取用户信息
func (service *UserService) GetById(id int64) *model.User {
	return service.repo.GetById(id)
}
