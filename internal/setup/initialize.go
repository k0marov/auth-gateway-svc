package setup

import (
	"auth-gateway-svc/internal/config"
	"auth-gateway-svc/internal/delivery"
	"log"
	"net/http"
)

func InitializeAndStart(cfg config.App) {
	srv := delivery.NewServer(cfg.HTTPServer)
	log.Printf("Listening at %s", cfg.HTTPServer.Host)
	log.Print(http.ListenAndServe(cfg.HTTPServer.Host, srv))
}
