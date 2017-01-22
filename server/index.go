package server

import (
	"net/http"

	"github.com/flosch/pongo2"
)

// index displays the home page.
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	services, _ := s.docker.Status()
	s.render(w, r, "index.html", pongo2.Context{
		"test":     "test",
		"services": services,
	})
}
