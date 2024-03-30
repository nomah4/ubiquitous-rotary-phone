package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"CLI_MSG_Service"`

	AppVersion string // `env:"APP_VER"`
	URLPing        string `env:"URL"`
	CheckPing  bool
	URLFile    string `env:"URL_FILE"`
	Sha256     string `env:"SHA256"`
	CheckFile  bool
	URLLogin   string `env:"URL_LOGIN"`
	CheckLogin bool
}

func GetBoolVar(name string) bool {
	return os.Getenv(name) == "true"
}

func GetStringVar(name string) string {
	variable := os.Getenv(name)
	if len(variable ) == 0 {
		log.Fatalf(fmt.Sprintf("Variable %s is empty", name))
	}
	return variable
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	cfg.AppName = os.Getenv("APP_NAME")
	if len(cfg.AppName) == 0 {
		return nil, errors.New("no APP_NAME env variable")
	}
	cfg.AppVersion = "0.0.1" // os.Getenv("APP_VER")
	cfg.CheckPing = GetBoolVar("CHECK_PING")
	if cfg.CheckPing {
		cfg.URLPing = GetStringVar("URL_PING")
	}

	cfg.CheckFile = GetBoolVar("CHECK_FILE")
	if cfg.CheckFile {
		cfg.URLFile = GetStringVar("URL_FILE")
		cfg.Sha256 = GetStringVar("SHA256")
	}

	cfg.CheckLogin = GetBoolVar("CHECK_LOGIN")
	if cfg.CheckLogin {
		cfg.URLLogin = GetStringVar("URL_LOGIN")
	}

	return &cfg, nil
}
