package repo

import (
	"billy/model"
)

// UserRepo 用户信息载体
type UserRepo struct {
}

func NewUserRepo() UserRepo {
	return UserRepo{}
}

// Insert 新增用户记录
func (r *UserRepo) Insert(user *model.User) {
	db.Create(user)
}

// Update 更新用户记录
func (r *UserRepo) Update(user *model.User) {
	db.Model(user).Updates(user)
}

// GetById 根据ID查询用户信息
func (r *UserRepo) GetById(id int64) *model.User {
	var user model.User
	db.Last(&user, id)
	return &user
}

// GetByUsername 根据用户名查询用户信息
func (r *UserRepo) GetByUsername(username string) *model.User {
	var user model.User
	db.Last(&user, "username=?", username)
	return &user
}

// CheckAndGet 检查用户密码并获取登录信息
func (r *UserRepo) CheckAndGet(username string, password string) *model.User {
	var user model.User
	db.Last(&user, "username=? and password=?", username, password)
	return &user
}
