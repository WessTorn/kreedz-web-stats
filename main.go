package main

import (
	"kreedz-web-stats/config"
	log "kreedz-web-stats/logger"
)

func main() {
	err := log.InitLogger("./log/")
	if err != nil {
		log.ErrorLogger.Fatal(err)
	}

	err = config.InitConfig()
	if err != nil {
		log.ErrorLogger.Fatal(err)
	}
}
