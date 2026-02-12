package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"DewaSRY/sociomile-app/pkg/lib/logger"
)

var DB *gorm.DB

func Connect() {
	_ = godotenv.Load()
	database_url := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(database_url), &gorm.Config{})
	if err != nil {
		logger.ErrorLog("Failed to connect to database", map[string]any{
			"errors": err.Error(),
		})

	}

	DB = db
	log.Println("✅ Database connected")
	logger.InfoLog("Database connected", map[string]any{
		"message": "connection success",
	})
}

func Close() {
	if DB == nil {
		return
	}

	sqlDB, err := DB.DB()
	if err != nil {
		logger.ErrorLog("Failed to get sql.DB for closing", map[string]any{
			"errors": err.Error(),
		})
		return
	}

	if err := sqlDB.Close(); err != nil {
		logger.ErrorLog("Error closing database connection", map[string]any{
			"errors": err.Error(),
		})
	} else {
		logger.InfoLog("Database connection closed", map[string]any{
			"message": "connection close success",
		})
	}
}
