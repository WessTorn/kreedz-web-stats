package main

import (
	"fmt"
	"kreedz-web-stats/api"
	"kreedz-web-stats/config"
	"kreedz-web-stats/database"
	log "kreedz-web-stats/logger"
	"kreedz-web-stats/site"
	"net/http"
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

	api := &api.Server{DB: db}
	api.InitApi()

	site := &site.Server{DB: db}
	site.InitSite()

	addr := fmt.Sprintf("%s:%s", config.BindIP(), config.BindPort())
	log.InfoLogger.Printf("Starting server on %s", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.ErrorLogger.Fatal(err)
	}

}
