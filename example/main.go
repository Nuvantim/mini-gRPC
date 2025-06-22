package main

import (
	"example/database"
	"example/internal/server"
	"example/internal/service"
	"log"
)

func main() {
	db := database.InitDB()
	defer db.Close()
	port := os.Getenv("PORT")
	srv := server.New(
		":"+port,
		category.New(db.Queries).Register, // Langsung pass method
		product.New(db.Queries).Register,  // Tanpa interface
	)

	log.Println("Server started on :8080")
	log.Fatal(srv.ListenAndServe())
}
