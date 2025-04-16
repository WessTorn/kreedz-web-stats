package site

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func (server *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("frontend", "head.gohtml"),
		filepath.Join("frontend", "navbar.gohtml"),
		filepath.Join("frontend", "home.gohtml"),
	)

	if err != nil {
		http.Error(w, "Ошибка шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "home", nil)
	if err != nil {
		http.Error(w, "Ошибка рендера: "+err.Error(), http.StatusInternalServerError)
	}
}
