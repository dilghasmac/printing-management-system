package main

import (
	"log/slog"
	"os"

	"printing-management-system/config"
	"printing-management-system/internal/database"
	"printing-management-system/internal/handler"
	"printing-management-system/internal/logger"
)

func main() {
	appConfig := config.LoadConfig()
	appLogger := logger.NewLogger()

	db, err := database.ConnectDatabase(appConfig)
	if err != nil {
		appLogger.Error(
			"database connection failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	appLogger.Info("database connected successfully")

	err = database.RunMigrations(db)
	if err != nil {
		appLogger.Error(
			"database migration failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	appLogger.Info("database migrations completed")

	router := handler.SetupRouter()

	appLogger.Info(
		"server starting",
		slog.String("port", appConfig.AppPort),
	)

	err = router.Run(":" + appConfig.AppPort)
	if err != nil {
		appLogger.Error(
			"server failed",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}
}
