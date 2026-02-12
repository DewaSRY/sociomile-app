package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TODO: update code to use logger
var DB *gorm.DB

func Connect() {
	_ = godotenv.Load()
	database_url := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(database_url), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	log.Println("✅ Database connected")
}