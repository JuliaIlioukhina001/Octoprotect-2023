package config

import (
	"encoding/base64"
	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	DBType             string
	SQLitePath         string
	PostgresDSN        string
	HTTPListenAddr     string
	JWTSecretKey       string
	PushoverToken      string
	MaxRetryNotify     int
	AdminTokenBcrypted string
	TelegramToken      string
}

func LoadConfig() {
	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	AppConfig.DBType = viper.GetString("DB_TYPE")
	AppConfig.SQLitePath = viper.GetString("SQLITE_PATH")
	AppConfig.PostgresDSN = viper.GetString("POSTGRES_DSN")
	AppConfig.HTTPListenAddr = viper.GetString("HTTP_LISTEN_ADDR")
	AppConfig.JWTSecretKey = viper.GetString("JWT_SECRET_KEY")
	AppConfig.PushoverToken = viper.GetString("PUSHOVER_TOKEN")
	AppConfig.MaxRetryNotify = viper.GetInt("MAX_RETRY_NOTIFY")
	AppConfig.TelegramToken = viper.GetString("TELEGRAM_TOKEN")

	tokenBytes, err := base64.StdEncoding.DecodeString(viper.GetString("ADMIN_TOKEN_BCRYPT_BASE64"))
	if err != nil {
		panic(err)
	}
	AppConfig.AdminTokenBcrypted = string(tokenBytes)
}
