package config

// appConf 应用配置
type appConf struct {
	// mode 应用模式
	mode appMode
	// server 服务相关信息
	server server
}

// server 服务器配置
type server struct {
	port int
}

// appMode 应用模式
type appMode string

const (
	// modeDev 开发模式
	modeDev appMode = "dev"
	// modeProd 生产模式
	modeProd appMode = "prod"
)
