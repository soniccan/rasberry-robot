package env

import (
	"log"

	"github.com/joho/godotenv"
)

// SetEnv は環境変数を設定する
func SetEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
