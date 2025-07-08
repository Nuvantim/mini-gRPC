package database

import (
	"context"
	"log"
	"sync"

	"example/config"
	"example/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
	DB      *pgxpool.Pool
	once    sync.Once
	Queries *repository.Queries
)

func InitDB() {
	once.Do(func() {
		dbconfig := config.DatabaseEnvirontment()
		var err error
		var dsn string = "postgres://" + dbconfig.user + ":" + dbconfig.password + "@" + dbconfig.host + ":" + dbconfig.port + "/" + dbconfig.name
		DB, err = pgxpool.New(context.Background(), dsn)
		if err != nil {
			log.Fatalf("Unable to create connection pool: %v", err)
		}

		// Inisialisasi Queries
		Queries = repository.New(DB)
	})
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
