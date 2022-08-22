package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("app started")
	envInit()
	connectToDB()
	fmt.Println("app ready")
}

func envInit() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Подключение к СУБД
var dbConnect *gorm.DB

// Функция подключения к СУБД
func connectToDB() {
	dbUser, dbPassword, dbName, dbPort, dbHost :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_HOST")
	intPort, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatal("Error parse port from env: ", err)
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, intPort, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connect to DB: ", err)
	}

	dbConnect = db
}
