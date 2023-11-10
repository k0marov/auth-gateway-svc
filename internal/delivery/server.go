package delivery

import (
	"auth-gateway-svc/internal/config"
	"auth-gateway-svc/internal/core"
	"auth-gateway-svc/internal/service"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type IAuthService interface {
	Register(login, password string) (*service.Tokens, error)
	Login(login, password string) (*service.Tokens, error)
}

type Server struct {
	cfg       config.HTTPServer
	forwarder http.Handler
	r         chi.Router
	svc       IAuthService
}

func NewServer(cfg config.HTTPServer, forwarder http.Handler, svc IAuthService) http.Handler {
	srv := &Server{cfg, forwarder, chi.NewRouter(), svc}
	srv.defineEndpoints()
	return srv
}

func (s *Server) defineEndpoints() {
	s.r.Post("/api/v1/auth/register", s.Register)
	s.r.Post("/api/v1/auth/sign-in", s.SignIn)
	s.r.Handle("/*", s.forwarder)
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	var authReq AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&authReq); err != nil {
		core.WriteErrorResponse(w, &core.ClientError{
			DisplayMessage: fmt.Sprintf("while decoding request: %v", err),
			HTTPCode:       0,
		})
		return
	}
	tokens, err := s.svc.Register(authReq.Login, authReq.Password)
	if err != nil {
		core.WriteErrorResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(tokens)
}

func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SignIn"))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
