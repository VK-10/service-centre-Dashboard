package main

import (
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port   string
	DbPath string
}

func loadConfig() Config {
	return Config{
		Port:   getEnv("PORT", "8080"),
		DbPath: getEnv("DATABASE_URL", "./data/vehicles.db"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func Loadtemplates(router *gin.Engine) error {
	functions := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}

	templ, err := template.New("").Funcs(functions).ParseGlob("templates/*.tmpl")
	if err != nil {
		return err
	}

	router.SetHTMLTemplate(templ)
	return nil
}
