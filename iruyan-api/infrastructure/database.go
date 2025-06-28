package infrastructure

import (
	"fmt"
	"iruyan-api/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	primaryDSN := os.Getenv("DATABASE_URL")
	if primaryDSN == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// デフォルトのデータベースURL
	defaultDSN := "postgres://defaultuser:defaultpassword@localhost:5432/defaultdb?sslmode=disable"

	var err error
	maxRetries := 5

	// データベース接続をリトライ
	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(postgres.Open(primaryDSN), &gorm.Config{})
		if err == nil {
			fmt.Println("Primary database connection successfully established.")
			break
		}
		log.Printf("Failed to connect to primary database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(2 * time.Second)
	}

	// プライマリ接続が確立できない場合、デフォルトのデータベースに接続を試行
	if err != nil {
		log.Println("Connecting to default database.")
		DB, err = gorm.Open(postgres.Open(defaultDSN), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to both primary and default databases: %v", err)
		}
		fmt.Println("Default database connection successfully established.")
	}

	// モデルをマイグレーション
	if err := DB.AutoMigrate(&models.User{}, &models.Room{}, &models.Seat{}, &models.WorkTime{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
