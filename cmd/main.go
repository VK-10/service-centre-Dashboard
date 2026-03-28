package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config := loadConfig()

	logger := slog.New(slog.newTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dbModel, err := models.InitDB(cfg.DBPath)
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
	slog.Info("Starting server on port", "port", cfg.Port)

	router.Run(":" + cfg.Port)
}
