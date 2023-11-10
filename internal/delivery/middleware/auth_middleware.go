package middleware

import (
	"auth-gateway-svc/internal/core"
	"auth-gateway-svc/internal/service"
	"net/http"
	"strings"
)

type TokenVerifier interface {
	Verify(token string) (u *service.UserClaims, ok bool)
}

func NewAuthMiddleware(verifier TokenVerifier) func(forceAdmin bool) func(http.HandlerFunc) http.HandlerFunc {
	return func(forceAdmin bool) func(http.HandlerFunc) http.HandlerFunc {
		return func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				token := strings.TrimPrefix(w.Header().Get("Authorization"), "Bearer ")
				user, ok := verifier.Verify(token)
				if !ok {
					r.Header.Del("Authorization")
				}
				if forceAdmin && (!ok || user.Login != "admin") {
					core.WriteErrorResponse(w, core.CEUnauthorized)
					return
				}
				next.ServeHTTP(w, r)
			}
		}
	}
}
