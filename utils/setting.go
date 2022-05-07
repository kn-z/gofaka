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

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	cfg, err := ini.Load("/home/gofaka/config/config.ini")
	if err != nil {
		log.Println("Failed to read config.ini", err)
		os.Exit(1)
	}
	LoadServer(cfg)
	LoadData(cfg)
}

func LoadServer(cfg *ini.File) {
	AppMode = cfg.Section("server").Key("AppMode").MustString("debug")
	HttpPort = cfg.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = cfg.Section("server").Key("JwtKey").MustString("")
}

func LoadData(cfg *ini.File) {
	Db = cfg.Section("database").Key("Db").MustString("mysql")
	DbHost = cfg.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = cfg.Section("database").Key("DbPort").MustString("3306")
	DbUser = cfg.Section("database").Key("DbUser").MustString("root")
	DbPassword = cfg.Section("database").Key("DbPassword").MustString("root")
	DbName = cfg.Section("database").Key("DbName").MustString("KNcloud")

}
