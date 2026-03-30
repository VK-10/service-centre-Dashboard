package main

import (
	"encoding/json"
	"html/template"
	"os"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port             string
	DbPath           string
	SessionSecretKey string
}

func loadConfig() Config {
	return Config{
		Port:             getEnv("PORT", "8080"),
		DbPath:           getEnv("DATABASE_URL", "./data/vehicles.db"),
		SessionSecretKey: getEnv("SESSION_SECRET_KEY", "service-centre-secret-key"),
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
		"json": func(v interface{}) template.JS {
			b, _ := json.Marshal(v)
			return template.JS(b)
		},
	}

	templ, err := template.New("").Funcs(functions).ParseGlob("templates/*.tmpl")
	if err != nil {
		return err
	}

	router.SetHTMLTemplate(templ)
	return nil
}

func setupSessionStore(db *gorm.DB, secretKey []byte) sessions.Store {
	store := gormsessions.NewStore(db, true, secretKey)
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
		Secure:   true,
		SameSite: 3,
	})

	return store
}

func SetSessionValue(c *gin.Context, key string, value interface{}) error {
	session := sessions.Default(c)
	session.Set(key, value)
	return session.Save()
}

// func GetSessionString(c *gin.Context, key string) string {
// 	session := sessions.Default(c)
// 	val := session.Get(key)
// 	if val == nil {
// 		return ""
// 	}
// 	strr, _ := val.string()
// 	return strr
// }

func GetSessionString(c *gin.Context, key string) string {
	session := sessions.Default(c)
	val := session.Get(key)
	if val == nil {
		return ""
	}
	strr, _ := val.(string)
	return strr
}

func ClearSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear()
	return session.Save()
}
