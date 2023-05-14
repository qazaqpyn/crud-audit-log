package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB     Mongo
	Server Server
}

type Mongo struct {
	URI      string
	Database string
}

type Server struct {
	Port int
}

func New() (*Config, error) {
	cfg := new(Config)

	if err := godotenv.Load(); err != nil {
		return nil, errors.New("No .env file found\n")
	}

	uri := os.Getenv("DB_URI")
	if uri == "" {
		return nil, errors.New("You must set your 'MONGODB_URI' environmental variable\n")
	}

	if err := envconfig.Process("db", &cfg.DB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("server", &cfg.Server); err != nil {
		return nil, err
	}

	cfg.DB.URI = uri

	return cfg, nil
}
