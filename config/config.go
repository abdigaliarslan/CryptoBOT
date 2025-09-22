package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramAPI string
	CryptoAPI   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TelegramAPI := os.Getenv("TELEGRAM_API")
	if TelegramAPI == "" {
		log.Fatal("TELEGRAM_API environment variable is required")
	}

	CryptoAPI := os.Getenv("CRYPTO_API")
	if CryptoAPI == "" {
		log.Fatal("CRYPTO_API environment variable is required")
	}

	return &Config{
		TelegramAPI: TelegramAPI,
		CryptoAPI:   CryptoAPI,
	}
}
