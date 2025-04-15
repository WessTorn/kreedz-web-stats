package site

import (
	"database/sql"
	"net/http"
)

type Server struct {
	DB *sql.DB
}

func (server *Server) InitSite() {
	http.HandleFunc("/", server.homeHandler) // Главная страница
}
