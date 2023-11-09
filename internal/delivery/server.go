package delivery

import (
	"auth-gateway-svc/internal/config"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	cfg config.HTTPServer
	r   chi.Router
}

func NewServer(cfg config.HTTPServer) http.Handler {
	srv := &Server{cfg, chi.NewRouter()}
	srv.defineEndpoints()
	return srv
}

func (s *Server) defineEndpoints() {
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
