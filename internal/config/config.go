package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port     string   `envconfig:"PORT" default:"8080"`
	Host     string   `envconfig:"HOST" default:"localhost"`
	DBConfig DBConfig `envconfig:"DBConfig"`
}

type DBConfig struct {
	Host string `envconfig:"DB_HOST" default:"localhost"`
	Port string `envconfig:"DB_PORT" default:"5432"`
	User string `envconfig:"DB_USER" default:"postgres"`
	Pass string `envconfig:"DB_PASS" default:"DB_PASS"`
}

func NewConfig() *Config {
	httpPort := os.Getenv("PORT")
	httpHost := os.Getenv("HOST")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	return &Config{
		Port: httpPort,
		Host: httpHost,
		DBConfig: DBConfig{
			Host: dbHost,
			Port: dbPort,
			User: dbUser,
			Pass: dbPass,
		},
	}
}

func LoadConfig(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
