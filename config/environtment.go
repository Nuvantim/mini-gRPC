package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AppName string `envconfig:"APP_NAME"`
}

type DatabaseConfig struct {
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
	Name     string `envconfig:"DB_NAME"`
}

type ServerConfig struct {
	Port  string `envconfig:"PORT_SERVICE"`
	Rate  int    `envconfig:"RATE"`
	Burst int    `envconfig:"BURST"`
	LRU   int    `envconfig:"LRU"`
}

func GetAppConfig() (*AppConfig, error) {
	var cfg AppConfig
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func CheckEnv() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", errors.New("Failed to load configuration file: " + err.Error())
	}
	return "Running Configuration..", nil
}

func GetDatabaseConfig() (*DatabaseConfig, error) {
	var cfg DatabaseConfig
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func GetServerConfig() (*ServerConfig, error) {
	var cfg ServerConfig
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
