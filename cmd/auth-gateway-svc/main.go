package main

import (
	"auth-gateway-svc/internal/config"
	"auth-gateway-svc/internal/setup"
)

func main() {
	cfg := config.ReadFromEnv()
	setup.InitializeAndStart(cfg)
}
