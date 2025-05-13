package database

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"micro/repository"
)

var (
	DB    *pgxpool.Pool
	once  sync.Once
	Query *repository.Queries
)

func InitDB() {
	// check .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	once.Do(func() {
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		port := os.Getenv("DB_PORT")
		db_name := os.Getenv("DB_NAME")
		var err error
		var dsn string = "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + db_name
		DB, err = pgxpool.New(context.Background(), dsn)
		if err != nil {
			log.Fatalf("Unable to create connection pool: %v", err)
		}
	})
	Query = repository.New(DB)
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
