package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	mongo_db_url string
	db_name      string
	auth_token   string
}

var Cfg Config

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("env file not found")
	}

	Cfg = Config{
		mongo_db_url: os.Getenv("MONGO_DB_URL"),
		db_name:      os.Getenv("DB_NAME"),
		auth_token:   os.Getenv("auth_token"),
	}
}

func (c Config) Get(key string) string {

	switch key {
	case "mongo_db_url":
		return c.mongo_db_url
	case "db_name":
		return c.db_name
	case "auth_token":
		return c.auth_token
	default:
		return ""
	}
}
