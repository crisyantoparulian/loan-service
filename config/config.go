package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	App      App
	Database Database
	Mail     SmtpMail
}

type App struct {
	Port int
}
type Database struct {
	PostgreDSN   string
	MaxOpenConn  int
	MaxIddleConn int
	DebugMode    bool
}

type SmtpMail struct {
	SMTPHost   string
	SMTPPort   int
	SMTPUser   string
	SMTPPass   string
	SMTPSender string
}

var (
	once sync.Once
	cfg  *Config
)

// LoadConfig from .env once
func LoadConfig() *Config {
	once.Do(func() {
		fmt.Println("APP ENV", os.Getenv("APP_ENV"))
		// Load .env file only in local development
		if os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == "" {
			if err := godotenv.Load(); err != nil {
				fmt.Println("Warning: No .env file found, using system environment variables")
			}
		}

		cfg = &Config{
			App: App{
				Port: getIntFromEnvWithDefaultVal("APP_PORT", 8080),
			},
			Database: Database{
				PostgreDSN:   os.Getenv("DATABASE_URL"),
				MaxOpenConn:  getIntFromEnvWithDefaultVal("MAX_OPEN_CONN", 20),
				MaxIddleConn: getIntFromEnvWithDefaultVal("MAX_IDDLE_CONN", 5),
				DebugMode:    true,
			},
			Mail: SmtpMail{
				SMTPHost:   os.Getenv("SMTP_HOST"),
				SMTPPort:   getIntFromEnvWithDefaultVal(os.Getenv("SMTP_PORT"), 587),
				SMTPUser:   os.Getenv("SMTP_USER"),
				SMTPPass:   os.Getenv("SMTP_PASS"),
				SMTPSender: os.Getenv("SMTP_SENDER"),
			},
		}
	})

	return cfg
}

func getIntFromEnvWithDefaultVal(key string, defaultVal int) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return defaultVal
	}
	return val
}
