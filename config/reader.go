package config

import (
	"github.com/zhyea/vibe"
	"log"
)

var Config = new(AppConf)

//
// init 初始化
func init() {
	vibe.AddConfigFiles("resources/app.yml")
	if err := vibe.ReadConfig(); nil != err {
		log.Fatalf("Load Config file app.yml failed: %v\n", err)
	}
	if err := vibe.Unmarshal(Config); nil != err {
		log.Fatalf("Render Config instance failed: %v\n", err)
	}
}

//
// IsDevMode 是否生产模式
func IsDevMode() bool {
	return ModeDev == Config.Mode
}

//
// Port 服务端口
func Port() int {
	if (Server{} == Config.Server || Config.Server.Port <= 0) {
		return 8080
	}
	return Config.Server.Port
}

//
// GetDbConfig 获取数据库连接信息
func GetDbConfig() Database {
	return Config.Database
}
