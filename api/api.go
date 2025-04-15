package api

import (
	"database/sql"
	"net/http"
)

type Server struct {
	DB *sql.DB
}

func (server *Server) InitApi() {
	http.HandleFunc("/api/user", server.userHandler)
	http.HandleFunc("/api/map", server.mapHandler)
	http.HandleFunc("/api/record", server.recordHandler)
}
