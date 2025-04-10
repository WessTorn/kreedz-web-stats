package config

import (
	log "kreedz-web-stats/logger"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BindIP string
	Port   string
}

var data Config

func InitConfig() error {
	err := godotenv.Load("config.env")
	if err != nil {
		return err
	}

	data = Config{
		BindIP: os.Getenv("BIND_IP"),
		Port:   os.Getenv("PORT"),
	}

	log.InfoLogger.Printf("Конфигурация загружена: %v", data)

	return nil
}

func BindIP() string {
	return data.BindIP
}

func BindPort() string {
	return data.BindIP
}
