package config

import (
	"github.com/zhyea/vibe"
	"log"
)

var config = new(appConf)

//
// init 初始化
func init() {
	vibe.AddConfigFiles("resources/app.yml")
	if err := vibe.ReadConfig(); nil != err {
		log.Fatalf("Load config file app.yml failed: %v\n", err)
	}
	if err := vibe.Unmarshal(config); nil != err {
		log.Fatalf("Render config instance failed: %v\n", err)
	}
}

// IsProd 是否生产模式
func IsProd() bool {
	return modeProd == config.mode
}

// ServerPort 获取服务端口
func ServerPort() int {
	if (server{} == config.server || config.server.port <= 0) {
		return 8080
	}
	return config.server.port
}
