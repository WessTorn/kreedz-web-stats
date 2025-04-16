package api

import (
	"database/sql"
	"kreedz-web-stats/config"
	log "kreedz-web-stats/logger"
	"net/http"
	"strings"
)

type Server struct {
	DB *sql.DB
}

func (server *Server) InitApi() {
	http.HandleFunc("/api/user", Protect(server.userHandler))
	http.HandleFunc("/api/map", Protect(server.mapHandler))
	http.HandleFunc("/api/record", Protect(server.recordHandler))
}

func Protect(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		ipOnly := strings.Split(ip, ":")[0]

		log.InfoLogger.Println("Request:", ipOnly, r.Header.Get("X-API-Key"))

		if ipOnly != config.RemoveIp() {
			http.Error(w, "Forbidden - IP not allowed", http.StatusForbidden)
			return
		}

		key := r.Header.Get("X-API-Key")
		if key != config.ApiKey() {
			http.Error(w, "Unauthorized - Invalid API Key", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
