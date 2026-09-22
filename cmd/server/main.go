package main

import (
	"log/slog"

	"printing-management-system/config"
	"printing-management-system/internal/logger"
)

func main() {
	appConfig := config.LoadConfig()
	appLogger := logger.NewLogger()

	appLogger.Info(
		"application starting",
		slog.String("appName", appConfig.AppName),
		slog.String("environment", appConfig.AppEnv),
		slog.String("port", appConfig.AppPort),
	)
}
