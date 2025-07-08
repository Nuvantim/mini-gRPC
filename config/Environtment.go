package config

import (
	"github.com/joho/godotenv"
	"os"
)

type DatabaseConfig struct {
	host     string
	port     string
	user     string
	password string
	name     string
}

type ServerConfig struct {
	port  string
	// rate  uint
	// burst uint
	// lru   uint
}

func CheckEnv() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}
	return "Environtment Added", nil
}

func DatabaseEnvirontment() DatabaseConfig {
	return DatabaseConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		name:     os.Getenv("DB_NAME"),
	}
}

func ServerEnvirontment() ServerConfig {
	return ServerConfig{
		port:  os.Getenv("PORT_SERVICE"),
		// rate:  os.Getenv("RATE"),
		// burst: os.Getenv("BURST"),
		// lru:   os.Getenv("LRU"),
	}
}
