package setup

import (
	"auth-gateway-svc/internal/config"
	"auth-gateway-svc/internal/delivery"
	"auth-gateway-svc/internal/delivery/forwarder"
	"auth-gateway-svc/internal/repository"
	"auth-gateway-svc/internal/service"
	"auth-gateway-svc/internal/service/bcrypt_hasher"
	"log"
	"net/http"
)

func InitializeAndStart(cfg config.App) {
	usersRepo := repository.NewUsersLevelDB(cfg.UsersDB.LevelDBPath)
	tokens := service.NewTokensService(cfg.JWT)
	hasher := bcrypt_hasher.NewBcryptHasher(8)
	svc := service.NewAuth(tokens, usersRepo, hasher)
	fwd := forwarder.NewForwarder(cfg.ForwardHost)
	srv := delivery.NewServer(cfg.HTTPServer, fwd, svc)
	log.Printf("Listening at %s", cfg.HTTPServer.Host)
	log.Print(http.ListenAndServe(cfg.HTTPServer.Host, srv))
}
