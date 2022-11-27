package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	globalC   *GlobalC
	logC      *LogC
	cookieC   *CookieC
	telegramC *TelegramC
)

type GlobalC struct {
	AppPath   string     `json:"app_path"`
	CookieC   *CookieC   `json:"cookie_c"`
	LogC      *LogC      `json:"log_c"`
	TelegramC *TelegramC `json:"telegram_c"`
}

type LogC struct {
	MaxSize    int `json:"max_size" mapstructure:"max_size"`
	MaxAge     int `json:"max_age" mapstructure:"max_age"`
	MaxBackups int `json:"max_backups" mapstructure:"max_backups"`
}

type CookieC struct {
	KuaishouCookie string `json:"kuaishou_cookie" mapstructure:"kuaishou_cookie"`
	XiguaCookie    string `json:"xigua_cookie" mapstructure:"xigua_cookie"`
}

type TelegramC struct {
	ApiToken string `json:"api_token" mapstructure:"api_token"`
}

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	viper.SetConfigFile(path + "/config.toml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("load config [file] err:", err)
	}
	err = viper.UnmarshalKey("cookie", &cookieC)
	if err != nil {
		log.Fatal("load config [cookie] err:", err)
	}
	err = viper.UnmarshalKey("log", &logC)
	if err != nil {
		log.Fatal("load config [log] err:", err)
	}
	err = viper.UnmarshalKey("telegram", &telegramC)
	if err != nil {
		log.Fatal("load config [telegram] err:", err)
	}
	globalC = &GlobalC{
		AppPath:   path,
		CookieC:   cookieC,
		LogC:      logC,
		TelegramC: telegramC,
	}
}

func GetAppVersion() string {
	return "v0.0.1"
}

func GetGlobalInstance() *GlobalC {
	return globalC
}

func GetCookieInstance() *CookieC {
	return cookieC
}

func GetLogInstance() *LogC {
	return logC
}

func GetTelegramInstance() *TelegramC {
	return telegramC
}
