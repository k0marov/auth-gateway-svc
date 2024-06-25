package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type HTTPServer struct {
	Host string `default:":8080"`
}

type JWT struct {
	PrivateKeyFilePath string `required:"true"`
	AccessExpireInSec  int    `default:"100000"`
	AdminPassword      string `required:"true"`
}

type UsersDB struct {
	LevelDBPath string `required:"true"`
}

type App struct {
	HTTPServer  HTTPServer
	ForwardHost string `required:"true"`
	JWT         JWT
	UsersDB     UsersDB
}

func ReadFromEnv() App {
	var cfg App
	err := envconfig.Process("auth_gateway", &cfg)
	if err != nil {
		log.Fatalf("while parsing app config from env: %v", err)
	}
	return cfg
}
