package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBString struct {
	User string
}

var (
	Port      string
	SecretKey string
)

func LoadConfig() {
	godotenv.Load()
	Port = os.Getenv("PORT")
	SecretKey = os.Getenv("SECRET_KEY")
	if Port == "" {
		Port = "8080"
	}
}

type DBcnf struct {
	Host      string
	DBPort    int
	DBname    string
	DBUser    string
	Password  string
	EnableSSL bool
}

func DBstringLoad() *DBcnf {
	godotenv.Load()
	DBhost := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")

	p, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Println("Invalid DBPORT:", err)
		return nil
	}

	DBuser := os.Getenv("DBUSER")
	DBName := os.Getenv("DBNAME")
	DBpassword := os.Getenv("PASSWORD")

	EnbSSL := os.Getenv("ENABLE_SSL_MODE")
	EnableSSLMode, _ := strconv.ParseBool(EnbSSL)

	if DBhost == "" || DBuser == "" || DBName == "" {
		log.Println("Missing required DB env variables")
		return nil
	}

	return &DBcnf{
		Host:      DBhost,
		DBPort:    p,
		DBUser:    DBuser,
		DBname:    DBName,
		Password:  DBpassword,
		EnableSSL: EnableSSLMode,
	}
}
