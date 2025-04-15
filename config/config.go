package config

import (
	log "kreedz-web-stats/logger"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BindIP    string
	Port      string
	ApiKey    string
	RemoveIps string
}

var data Config

func InitConfig() error {
	err := godotenv.Load("config.env")
	if err != nil {
		return err
	}

	data = Config{
		BindIP:    os.Getenv("BIND_IP"),
		Port:      os.Getenv("BIND_PORT"),
		ApiKey:    os.Getenv("API_KEY"),
		RemoveIps: os.Getenv("REMOVE_IP"),
	}

	log.InfoLogger.Printf("Конфигурация загружена: %v", data)

	return nil
}

func BindIP() string {
	return data.BindIP
}

func BindPort() string {
	return data.Port
}

func ApiKey() string {
	return data.ApiKey
}

func RemoveIp() string {
	return data.RemoveIps
}
