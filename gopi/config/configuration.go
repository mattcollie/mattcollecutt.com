package config

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
	Metrics struct {
		GTM string `toml:"gtm" env:"GTM" env-default:""`
	} `toml:"metrics"`
}

var Cfg Config

func ReadConfig() {
	err := cleanenv.ReadConfig("config.toml", &Cfg)
	if err != nil {
		log.Fatalf("Fatal error config file: %s \\n", err)
	}
}
