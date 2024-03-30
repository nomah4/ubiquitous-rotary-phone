package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"CLI_MSG_Service"`

	AppVersion string // `env:"APP_VER"`
	URL        string `env:"URL"`
	CheckPing  bool
	URLFile    string `env:"URL_FILE"`
	Sha256     string `env:"SHA256"`
	CheckFile  bool
	URLLogin   string `env:"URL_LOGIN"`
	CheckLogin bool
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Ошибка при загрузке файла .env: - %v ", err)
	}

	cfg.AppName = os.Getenv("APP_NAME")
	cfg.AppVersion = "0.0.1" // os.Getenv("APP_VER")
	cfg.URL = os.Getenv("URL")
	checkPing := os.Getenv("CHECK_PING")
	if checkPing == "true" {
		cfg.CheckPing = true
	}

	cfg.URLFile = os.Getenv("URL_FILE")
	cfg.Sha256 = os.Getenv("SHA256")
	checkFile := os.Getenv("CHECK_FILE")
	if checkFile == "true" {
		cfg.CheckFile = true
	}

	cfg.URLLogin = os.Getenv("URL_LOGIN")
	checkLogin := os.Getenv("CHECK_LOGIN")
	if checkLogin == "true" {
		cfg.CheckLogin = true
	}

	return &cfg, nil
}
