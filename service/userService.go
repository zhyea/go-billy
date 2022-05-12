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
func NewUserService(repo repo.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

// CheckLogin 检查用户登录信息
func (service *UserService) CheckLogin(username string, password string) *model.User {
	password = codec.MD5(password)
	return service.repo.CheckAndGet(username, password)
}
