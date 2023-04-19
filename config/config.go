package config

import (
	"log"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigData struct {
	PortNum string
	Static  string
	LogFile string
}

var Config ConfigData

func init() {
	LoadConfig()
	// utils.LoggingSettings(Config.LogFile)

}

// Configファイルを読み取り情報を取得する
func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigData{
		PortNum: cfg.Section("web").Key("port").MustString("8080"),
		Static:  cfg.Section("web").Key("static").String(),
		LogFile: cfg.Section("web").Key("logfile").String(),
	}

}
