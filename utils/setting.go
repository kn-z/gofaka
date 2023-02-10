package utils

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DbType   string
	DbHost   string
	DbPort   string
	DbUser   string
	DbPasswd string
	DbName   string

	EmHost   string
	EmPort   int
	EmUser   string
	EmPasswd string

	WebName    string
	WebUrl     string
	WebBpColor string
	AdminPath  string
)

func Init() {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Println("Failed to read config.ini", err)
		os.Exit(1)
	}
	LoadServer(cfg)
	LoadData(cfg)
	LoadEmail(cfg)
	LoadWebInfo(cfg)
}

func LoadServer(cfg *ini.File) {
	AppMode = cfg.Section("server").Key("AppMode").MustString("debug")
	HttpPort = cfg.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = cfg.Section("server").Key("JwtKey").MustString("")
}

func LoadData(cfg *ini.File) {
	DbType = cfg.Section("database").Key("DbType").MustString("mysql")
	DbHost = cfg.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = cfg.Section("database").Key("DbPort").MustString("3306")
	DbUser = cfg.Section("database").Key("DbUser").MustString("root")
	DbPasswd = cfg.Section("database").Key("DbPasswd").MustString("root")
	DbName = cfg.Section("database").Key("DbName").MustString("KNcloud")
}

func LoadEmail(cfg *ini.File) {
	EmHost = cfg.Section("email").Key("EmHost").MustString("smtp.example.com")
	EmPort = cfg.Section("email").Key("EmPort").MustInt(587)
	EmUser = cfg.Section("email").Key("EmUser").MustString("emuser")
	EmPasswd = cfg.Section("email").Key("EmPasswd").MustString("empasswd")
}

func LoadWebInfo(cfg *ini.File) {
	WebName = cfg.Section("web").Key("WebName").MustString("KNcloud")
	WebUrl = cfg.Section("web").Key("WebUrl").MustString("www.kncloud.app")
	WebBpColor = cfg.Section("web").Key("WebBpColor").MustString("#35393e")
	AdminPath = cfg.Section("web").Key("AdminPath").MustString("backend")
}
