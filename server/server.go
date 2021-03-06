package server

import (
	"crypto/tls"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hectane/go-asyncserver"
	"github.com/nathan-osman/dcui/docker"
)

// Server provides the web interface for DCUI.
type Server struct {
	config *Config
	docker *docker.Docker
	server *server.AsyncServer
}

// New creates a new server with the specified configuration.
func New(config *Config, docker *docker.Docker) (*Server, error) {
	var (
		s = &Server{
			config: config,
			docker: docker,
			server: server.New(config.Addr),
		}
		r = mux.NewRouter()
	)
	if len(config.TLSCert) != 0 && len(config.TLSKey) != 0 {
		cert, err := tls.LoadX509KeyPair(
			config.TLSCert,
			config.TLSKey,
		)
		if err != nil {
			return nil, err
		}
		s.server.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	}
	s.server.Handler = r
	r.HandleFunc("/", s.index)
	r.HandleFunc("/action", s.action)
	r.PathPrefix("/static").Handler(http.FileServer(HTTP))
	if err := s.server.Start(); err != nil {
		return nil, err
	}
	return s, nil
}

// Close stops the server.
func (s *Server) Close() {
	s.server.Stop()
}
