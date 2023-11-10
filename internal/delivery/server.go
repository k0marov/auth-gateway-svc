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

type AuthMiddleware = func(forceAdmin bool) func(http.HandlerFunc) http.HandlerFunc

type Server struct {
	cfg       config.HTTPServer
	forwarder http.Handler
	r         chi.Router
	svc       IAuthService
	authMW    AuthMiddleware
}

func NewServer(cfg config.HTTPServer, forwarder http.Handler, svc IAuthService, authMW AuthMiddleware) http.Handler {
	srv := &Server{cfg, forwarder, chi.NewRouter(), svc, authMW}
	srv.defineEndpoints()
	return srv
}

func (s *Server) defineEndpoints() {
	s.r.Post("/api/v1/auth/register", s.authMW(true)(s.Register))
	s.r.Post("/api/v1/auth/login", s.authMW(false)(s.Login))
	s.r.Handle("/*", s.authMW(false)(s.forwarder.ServeHTTP))
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	var authReq AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&authReq); err != nil {
		core.WriteErrorResponse(w, &core.ClientError{
			DisplayMessage: fmt.Sprintf("while decoding request: %v", err),
			HTTPCode:       http.StatusBadRequest,
		})
		return
	}
	tokens, err := s.svc.Register(authReq.Login, authReq.Password)
	if err != nil {
		core.WriteErrorResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(TokensResponse{AccessToken: tokens.Access})
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var authReq AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&authReq); err != nil {
		core.WriteErrorResponse(w, &core.ClientError{
			DisplayMessage: fmt.Sprintf("while decoding request: %v", err),
			HTTPCode:       http.StatusBadRequest,
		})
		return
	}
	tokens, err := s.svc.Login(authReq.Login, authReq.Password)
	if err != nil {
		core.WriteErrorResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(TokensResponse{AccessToken: tokens.Access})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
