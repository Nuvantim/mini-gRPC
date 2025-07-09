package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DatabaseConfig struct {
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
	Name     string `envconfig:"DB_NAME"`
}

type ServerConfig struct {
	Port  string `envconfig:"PORT_SERVICE"`
	Rate  uint   `envconfig:"RATE"`
	Burst uint   `envconfig:"BURST"`
	LRU   uint   `envconfig:"LRU"`
}

func CheckEnv() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", errors.New("failed to load configuration file: " + err.Error())
	}
	return "Running Configuration..", nil
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	var cfg DatabaseConfig
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func LoadServerConfig() (*ServerConfig, error) {
	var cfg ServerConfig
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

