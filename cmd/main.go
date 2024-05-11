package main

import (
	"Bakers_backend/internal/delivery"
	"Bakers_backend/pkg/config"
	"Bakers_backend/pkg/database"
	"Bakers_backend/pkg/logger"
)

func main() {
	log, loggerInfoFile, loggerErrorFile := logger.InitLogger()

	defer loggerInfoFile.Close()
	defer loggerErrorFile.Close()

	config.InitConfig()
	log.Info("Config initialized")

	db := database.GetDB()
	log.Info("Database initialized")

	delivery.Start(db, log)
	//db.Close() // !! УДАЛИТЬ ПРИ ЗАПУСКЕ
}
