package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
