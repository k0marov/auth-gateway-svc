package middleware

import "net/http"

func NewAuthMiddleware() func(forceAdmin bool) func(http.HandlerFunc) http.HandlerFunc {
	return func(forceAdmin bool) func(http.HandlerFunc) http.HandlerFunc {
		return func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				// TODO: implement auth middleware
				next.ServeHTTP(w, r)
			}
		}
	}
}
