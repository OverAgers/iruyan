package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	JWTSecret       string
	DefaultRoomName string
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .envファイルが見つかりません")
	}

	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		log.Println("⚠️ JWT_SECRET が設定されていません")
	}

	DefaultRoomName = os.Getenv("DEFAULT_ROOM_NAME")
	if JWTSecret == "" {
		log.Println("⚠️ DEFAULT_ROOM_NAME が設定されていません")
	}
}
