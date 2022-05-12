package config

// appConf 应用配置
type appConf struct {

	// mode 应用模式
	mode appMode

	// server 服务相关信息
	server server

	// database 数据库连接信息
	database Database
}

// appMode 应用模式
type appMode string

const (
	// modeDev 开发模式
	modeDev appMode = "dev"

	// modeProd 生产模式
	modeProd appMode = "prod"
)

// server 服务器配置
type server struct {
	// port 应用端口
	port int
}

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
