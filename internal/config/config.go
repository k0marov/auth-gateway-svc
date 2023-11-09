package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type HTTPServer struct {
	Host string `default:":8080"`
}

type App struct {
	HTTPServer  HTTPServer
	ForwardHost string
}

func ReadFromEnv() App {
	var cfg App
	err := envconfig.Process("profiles", &cfg)
	if err != nil {
		log.Panicf("while parsing app config from env: %w", err)
	}
	return cfg
}
