package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (server *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type request struct {
		SteamID string `json:"steam_id"`
		Name    string `json:"name"`
	}

	type response struct {
		ID int64 `json:"id"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.SteamID == "" {
		http.Error(w, "SteamID is required", http.StatusBadRequest)
		return
	}

	var id int64
	err := server.DB.QueryRow("SELECT id FROM users WHERE steam_id = ?", req.SteamID).Scan(&id)
	if err == sql.ErrNoRows {
		res, err := server.DB.Exec("INSERT INTO users (steam_id, name) VALUES (?, ?)", req.SteamID, req.Name)
		if err != nil {
			http.Error(w, "Database insert error", http.StatusInternalServerError)
			return
		}
		id, _ = res.LastInsertId()
	} else if err != nil {
		http.Error(w, "Database select error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{ID: id})
}
