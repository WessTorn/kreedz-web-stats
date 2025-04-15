package main

import (
	"kreedz-web-stats/config"
	"kreedz-web-stats/database"
	log "kreedz-web-stats/logger"
)

func main() {
	err := log.InitLogger("./log/")
	if err != nil {
		log.ErrorLogger.Fatal(err) // TODO:
	}

	err = config.InitConfig()
	if err != nil {
		log.ErrorLogger.Fatal(err)
	}

	db, err := database.InitDB()
	if err != nil {
		log.ErrorLogger.Printf("Ошибка при инициализации бд: %v", err)
		return
	}
	defer db.Close()

	err = database.MigrateDB(db)
	if err != nil {
		log.ErrorLogger.Printf("Ошибка при миграции бд: %v", err)
		return
	}

}
