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

	// Front 前端相关配置
	Front Front
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

//
// Front 前端配置
type Front struct {

	// Theme 主题
	Theme string

	// Favicon  网站图标路径
	Favicon string

	// Template 模板信息
	Template Template
}

//
// Template 模板相关信息
type Template struct {

	// Path 路径
	Path string

	// Extension 模板文件扩展名
	Extension string
}
