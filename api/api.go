package api

import (
	"database/sql"
	"fmt"
	"kreedz-web-stats/config"
	log "kreedz-web-stats/logger"
	"net/http"
)

func InitApi(db *sql.DB) {
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		userHandler(w, r, db)
	})

	http.HandleFunc("/api/map", func(w http.ResponseWriter, r *http.Request) {
		mapHandler(w, r, db)
	})

	http.HandleFunc("/api/record", func(w http.ResponseWriter, r *http.Request) {
		recordHandler(w, r, db)
	})

	log.InfoLogger.Printf("Starting server on %s:%s ...", config.BindIP(), config.BindPort())

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", config.BindIP(), config.BindPort()), nil)
	if err != nil {
		log.ErrorLogger.Fatal(err)
	}
}
