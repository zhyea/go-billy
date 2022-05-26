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

//
// TemplateDir 模板路径
func TemplateDir() string {
	if (Front{}) == Config.Front || (Template{}) == Config.Front.Template || "" == Config.Front.Template.Path {
		return "./web/template"
	}
	return Config.Front.Template.Path
}

//
// TemplateExtension 模板扩展名
func TemplateExtension() string {
	if (Front{}) == Config.Front || (Template{}) == Config.Front.Template || "" == Config.Front.Template.Extension {
		return ".html"
	}
	return Config.Front.Template.Extension
}

//
// ThemeStatic 静态文件路径
func ThemeStatic() string {
	if "" == Config.Front.Theme {
		return TemplateDir() + "/static"
	}
	return TemplateDir() + "/" + Config.Front.Theme + "/static"
}

//
// AdminStatic 管理后台静态文件目录
func AdminStatic() string {
	return TemplateDir() + "/admin/static"
}

//
// Favicon 网站图标路径
func Favicon() string {
	if "" == Config.Front.Favicon {
		return TemplateDir() + "/" + Config.Front.Theme + "/static/imgs/favicon.ico"
	}
	return TemplateDir() + "/static/imgs/favicon.ico"
}
