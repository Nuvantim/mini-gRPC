package database

import (
	"context"
	"log"
	"sync"

	"example/config"
	"example/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB      *pgxpool.Pool
	once    sync.Once
	Queries *repository.Queries
)

func InitDB() {
	once.Do(func() {
		// Load Database Configuration
		var dbconfig, errs = config.GetDatabaseConfig()
		if errs != nil {
			log.Fatal(errs)
		}
		// Start Connection Database
		var err error
		var dsn string = "postgres://" + dbconfig.User + ":" + dbconfig.Password + "@" + dbconfig.Host + ":" + dbconfig.Port + "/" + dbconfig.Name
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
