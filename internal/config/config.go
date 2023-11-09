package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type HTTPServer struct {
	Host string `default:":8080"`
}

type JWT struct {
	PrivateKey        string `required:"true"`
	AccessExpireInSec int    `default:"100000"`
	AdminSecret       string `required:"true"`
}

type UsersDB struct {
	LevelDBPath string `required:"true"`
}

type App struct {
	HTTPServer  HTTPServer
	ForwardHost string
	JWT         JWT
	UsersDB     UsersDB
}

func ReadFromEnv() App {
	var cfg App
	err := envconfig.Process("profiles", &cfg)
	if err != nil {
		log.Panicf("while parsing app config from env: %w", err)
	}
	return cfg
}
