package site

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func (server *Server) mapsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		filepath.Join("frontend", "maps.gohtml"),
		filepath.Join("frontend", "head.gohtml"),
		filepath.Join("frontend", "navbar.gohtml"),
	)

	type exampleMaps struct {
		Num  int
		Name string
		ID   int
	}

	var Maps = []exampleMaps{
		{Num: 1, Name: "de_dust2", ID: 0},
		{Num: 2, Name: "de_inferno", ID: 1},
		{Num: 3, Name: "de_nuke", ID: 2},
	}

	if err != nil {
		http.Error(w, "Ошибка шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "maps", Maps)
	if err != nil {
		http.Error(w, "Ошибка рендера: "+err.Error(), http.StatusInternalServerError)
	}
}
