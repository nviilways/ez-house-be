package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type dbConfig struct {
	Host string
	User string
	Password string
	DBName string
	Port string
}

func getENV(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

var (
	ENV = getENV("ENV", "testing")
	AppName = "Ez House"
	SecretKey = getENV("SECRET_KEY", "")
	DBConfig = dbConfig {
		Host: getENV("DB_HOST", "localhost"),
		User: getENV("DB_USER", ""),
		Password: getENV("DB_PASSWORD", ""),
		DBName: getENV("DB_NAME", ""),
		Port: getENV("DB_PORT", "5432"),
	}
	AdminKey = getENV("SECRET_KEY_ADMIN", "")
)