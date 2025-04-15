package site

import (
	"fmt"
	"net/http"
)

func (server *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to the site!</h1>")
}
