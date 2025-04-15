package main

import (
	"log"
	"os"

	"github.com/didanslmn/movie-reservation-api/config"
	"github.com/didanslmn/movie-reservation-api/router"
)

func main() {
	config.LoadEnv()
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	r := router.SetupRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	log.Printf("Server running at http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
