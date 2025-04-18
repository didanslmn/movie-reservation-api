package main

import (
	"fmt"
	"log"
	"os"

	"github.com/didanslmn/movie-reservation-api/config"
	"github.com/didanslmn/movie-reservation-api/internal/user/model"
	"github.com/didanslmn/movie-reservation-api/router"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	config.LoadEnv()
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	```
	newUser := model.User{
		Name:     "admin123",
		Email:    "admin123@yahoo.com",
		Password: "11111111",
		Role:     model.RoleAdmin,
	};
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashedPassword)
	// Melakukan insert data ke database
	result := db.Create(&newUser)
	if result.Error != nil {
		fmt.Println("Gagal menambahkan user:", result.Error)
		return
	}

	// Data user yang berhasil di-insert akan terisi ID dan CreatedAt
	fmt.Println("Berhasil menambahkan user dengan ID:", newUser.ID)
	```
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatalf("JWT_SECRET environment variable is not set")
	}
	r := router.SetupRouter(db, jwtSecret)
	if r == nil {
		log.Fatalf("Failed to setup router")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	log.Printf("Server running at http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
