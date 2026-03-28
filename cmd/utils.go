package main

type Config struct {
	Post string
	DBPath string

}

func loadConfig() Config {
	return Config {
		Port : getenv("PORT", "8080")
		DbPath: getEnv("DATABASE_URL", "./data/vehicles.db")
	}
}

func getEnv(key, defaultValue string) string{
	if value := os.Getenv(key); value != "" {

	}

	return defaultValue
}

func loadTemplates(router *gin.Engine) error {
	functions := template.Funcmap {
		"add" func(a, b int) int { return a + b},

	}

	templ , err := template.New("").Funcs(functions).ParseGlob("templates/*.tmpl")
	if err != nil {
		return err
	}

	router.setHTMLTemplate(templ)
	return nil
}