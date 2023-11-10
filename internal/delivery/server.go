package delivery

import (
	_ "auth-gateway-svc/docs"
	"auth-gateway-svc/internal/config"
	"auth-gateway-svc/internal/core"
	"auth-gateway-svc/internal/service"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
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

//	@title			auth-gateway-svc
//	@version		1.0
//	@description	Auth gateway which provides JWT creation, validation, user creation and login.
//	@description	It handles endpoints for login and registration, and all other requests it proxies to the next gateway,
//	@description	but before that it checks that the JWT provided in Authorization header is valid.
//  @description	If JWT is invalid, the Authorization header is deleted before proxying.
//  @description	Authorization header must start with "Bearer " prefix.

//	@contact.name	Sam Komarov
//	@contact.url	github.com/k0marov
//	@contact.email	sam@skomarov.com

// @host		localhost:8080
// @schemes     https http
func (s *Server) defineEndpoints() {
	s.r.Get("/swagger/*", httpSwagger.WrapHandler)
	s.r.Post("/api/v1/auth/register", s.authMW(true)(s.Register))
	s.r.Post("/api/v1/auth/login", s.authMW(false)(s.Login))
	s.r.Handle("/*", s.authMW(false)(s.forwarder.ServeHTTP))
}

// Register godoc
//
//		@Summary		Register a user and return auth tokens for the created user.
//		@Summary		Only admins can register users, so if the caller's login is not 'admin', 403 is returned.
//		@Tags			auth
//		@Accept 		json
//	    @Param			auth	body		AuthRequest	true	"auth data"
//		@Produce		json
//		@Success		200	{object}	TokensResponse
//		@Failure		400	{object}	core.ClientError
//		@Failure		403	{object}	core.ClientError
//		@Router			/api/v1/auth/register [post]
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

// Login godoc
//
//	@Summary		Login using an email and password. Returns auth tokens.
//	@Tags			auth
//	@Accept 		json
//	@Param			auth	body		AuthRequest	true	"auth data"
//	@Produce		json
//	@Success		200	{object}	TokensResponse
//	@Failure		400	{object}	core.ClientError
//	@Router			/api/v1/auth/login [post]
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
