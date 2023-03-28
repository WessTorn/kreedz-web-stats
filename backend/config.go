package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		BindIP string `yaml:"bind_ip" env-default:"0.0.0.0"`
		Port   string `yaml:"port" env-default:"5000"`
	}
	DataBase struct {
		Host string `yaml:"host" env-default:"127.0.0.1"`
		User string `yaml:"user" env-default:"root"`
		Pass string `yaml:"pass" env-default:"root"`
		DB   string `yaml:"db" env-default:"kreedz"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("Read config")
		instance = &Config{}
		err := cleanenv.ReadConfig("config.yml", instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			panic(err)
		}
	})
	return instance
}
