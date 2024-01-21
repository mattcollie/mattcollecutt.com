package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Server struct {
		Host string `toml:"host" env:"HOST" env-default:"127.0.0.1"`
		Port string `toml:"port" env:"PORT" env-default:"3000"`
	} `toml:"server"`
	Instagram struct {
		UserId      string `toml:"userid"`
		AccessToken string `toml:"access_token"`
	} `toml:"instagram"`
}

var cfg Config

func readConfig() {
	err := cleanenv.ReadConfig("config.toml", &cfg)
	if err != nil {
		log.Fatalf("Fatal error config file: %s \\n", err)
	}
}
