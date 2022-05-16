package model

// User 用户信息
type User struct {

	// Id value
	Id string `json:"id" gorm:"primaryKey"`

	// Username 用户名
	Username string `json:"username"`

	// Password 密码
	Password string `json:"password"`

	// Nickname 昵称
	Nickname string `json:"nickname"`

	// Email 邮箱
	Email string `json:"email"`

	// OpTime 操作时间
	OpTime string `json:"-"`
}
