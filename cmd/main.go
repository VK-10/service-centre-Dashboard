package main

import (
	"log/slog"
	"os"
	"service-tracker-go/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config := loadConfig()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dbModel, err := models.InitDB(config.DbPath)
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)

	}

	slog.Info("Databse initialised successfully")

	RegisterCustomValidator()

	h := NewHandler(dbModel)

	router := gin.Default()

	if err := Loadtemplates(router); err != nil {
		slog.Error("Failed to load templates", "error", err)
		os.Exit(1)
	}

	setUpRoutes(router, h)
	slog.Info("Starting server on port", "port", config.Port)

	router.Run(":" + config.Port)
}
