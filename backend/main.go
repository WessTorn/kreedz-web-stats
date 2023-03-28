package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/index.gohtml", "frontend/head.gohtml", "frontend/header.gohtml")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		panic(err)
	}
}

func handleFunc(db *sql.DB) {
	http.Handle("/frontend/static/", http.StripPrefix("/frontend/static/", http.FileServer(http.Dir("./frontend/static/"))))

	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/lastpro/", GetLastPro(db)).Methods("GET")
	rtr.HandleFunc("/players/", GetPlayersTop(db)).Methods("GET")
	rtr.HandleFunc("/maps/", GetMapsStatic(db)).Methods("GET")
	rtr.HandleFunc("/map/pro/{id:[0-9]+}", GetMapRecordPro(db)).Methods("GET")
	rtr.HandleFunc("/map/nub/{id:[0-9]+}", GetMapRecordNub(db)).Methods("GET")
	rtr.HandleFunc("/player/{id:[0-9]+}", GetStatusPlayer(db)).Methods("GET")

	http.Handle("/", rtr)
}

func main() {
	cfg := GetConfig()
	var dataSourceName = fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.DataBase.User, cfg.DataBase.Pass, cfg.DataBase.Host, cfg.DataBase.DB)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Println("open", dataSourceName)
	handleFunc(db)
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port), nil)
	if err != nil {
		panic(err)
	}
}
