package delivery

import (
	"auth-gateway-svc/internal/config"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	cfg       config.HTTPServer
	forwarder http.Handler
	r         chi.Router
}

func NewServer(cfg config.HTTPServer, forwarder http.Handler) http.Handler {
	srv := &Server{cfg, forwarder, chi.NewRouter()}
	srv.defineEndpoints()
	return srv
}

func (s *Server) defineEndpoints() {
	s.r.Post("/api/v1/register", s.Register)
	s.r.Post("/api/v1/sign-in", s.SignIn)
	s.r.Handle("/*", s.forwarder)
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register"))
}

func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SignIn"))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
