package main

import (
	"log"

	"DewaSRY/sociomile-app/internal/config"
	"DewaSRY/sociomile-app/internal/database"
	"DewaSRY/sociomile-app/internal/database/seeders"
	"DewaSRY/sociomile-app/pkg/lib/logger"
)

func main() {
	config.Load()
	logger.Init()
	database.Connect()

	logger.InfoLog("start initial seeder", map[string]any{} )

	if err := seeders.ClearAllInitTables(); err != nil{
		log.Fatalf("clear table  failed: %v", err)
	}

	if err := seeders.SeedInitialData(); err != nil {
		log.Fatalf("Seeding failed: %v", err)
	}

}
