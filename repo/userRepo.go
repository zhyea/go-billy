package repo

import (
	"billy/model"
	"sync"
)

// UserRepo 用户方法
type UserRepo interface {

	// GetByUsername 根据用户名查询用户信息
	GetByUsername(username string) model.User

	// CheckAndGet 检查用户密码并获取登录信息
	CheckAndGet(username string, password string) model.User
}

// userQueryRepo 用户信息载体
type userQueryRepo struct {
	source map[int64]model.User
	mu     sync.RWMutex
}

func (r *userQueryRepo) GetByUsername(username string) model.User {

	return model.User{}
}

func (r *userQueryRepo) CheckAndGet(username string, password string) model.User {
	return model.User{}
}
