package config

//
// AppConf 应用配置
type AppConf struct {

	// Mode 应用模式
	Mode AppMode

	// Server 服务相关信息
	Server Server

	// Database 数据库连接
	Database Database
}

//
// AppMode 应用模式
type AppMode string

const (

	// ModeDev 开发模式
	ModeDev AppMode = "dev"

	// ModeProd 生产模式
	ModeProd AppMode = "prod"
)

//
// Server 服务器配置
type Server struct {

	// Port 应用端口
	Port int
}

//
// Database 数据库连接配置
type Database struct {

	// Dsn 数据库连接
	Dsn string

	// Timeout	连接超时时间
	Timeout string

	// MaxOpen 最大连接数
	MaxOpen int

	// MaxIdle 最大空闲连接数
	MaxIdle int
}
