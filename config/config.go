package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	Port      string
	SecretKey string
)

func LoadConfig() {
	godotenv.Load()
	Port = os.Getenv("PORT")
	SecretKey = os.Getenv("SECRET_KEY")
	if Port == ""{
		Port = "8080"
	}
}