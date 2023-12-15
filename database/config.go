package database

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type DBConfig struct {
	Name     string `env:"DB_NAME" default:"db_1"`
	Host     string `env:"DB_HOST" default:"localhost"` //DIHAPUS
	Port     string `env:"DB_PORT" default:"3306"`      //DIHAPUS
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

type ServerConfig struct {
	ServiceHost string `env:"SERVICE_HOST"`
	ServicePort string `env:"SERVICE_PORT"`
	DB          DBConfig
}

var Config ServerConfig

func init() {
	err := loadConfig()
	if err != nil {
		panic(err)
	}
}

func loadConfig() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Warn().Msg("Cannot find .env file. OS Environments will be used")
	}
	err = env.Parse(&Config)

	return err
}
