package setup

import (
	"auth-gateway-svc/internal/config"
	"auth-gateway-svc/internal/delivery"
	"auth-gateway-svc/internal/delivery/forwarder"
	"log"
	"net/http"
)

func InitializeAndStart(cfg config.App) {
	fwd := forwarder.NewForwarder(cfg.ForwardHost)
	srv := delivery.NewServer(cfg.HTTPServer, fwd)
	log.Printf("Listening at %s", cfg.HTTPServer.Host)
	log.Print(http.ListenAndServe(cfg.HTTPServer.Host, srv))
}
