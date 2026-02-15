package tests

import (
	"DewaSRY/sociomile-app/internal/config"
	"DewaSRY/sociomile-app/internal/database"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	godotenv.Load("../../../.env")
	config.Load()
	logger.Init();
	db := database.Connect()

	tx := db.Begin()
	if tx.Error != nil {
		t.Fatalf("failed to start transaction: %v", tx.Error)
	}

	t.Cleanup(func() {
		tx.Rollback()
	})

	return tx
}
