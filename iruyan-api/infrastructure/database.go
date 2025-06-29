package infrastructure

import (
	"fmt"
	"iruyan-api/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	primaryDSN := os.Getenv("DATABASE_URL")
	if primaryDSN == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	defaultDSN := "postgres://defaultuser:defaultpassword@localhost:5432/defaultdb?sslmode=disable"

	var err error
	var db *gorm.DB
	maxRetries := 5

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(primaryDSN), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // SQLログを出力
		})
		if err == nil {
			DB = db // グローバル変数に代入
			fmt.Println("✅ Primary database connection established.")
			break
		}
		log.Printf("⚠️ Failed to connect to primary database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Println("🔁 Connecting to default database.")
		db, err = gorm.Open(postgres.Open(defaultDSN), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("❌ Failed to connect to both primary and default databases: %v", err)
		}
		DB = db
		fmt.Println("✅ Default database connection established.")
	}

	// 🔽 モデルのマイグレーション（依存関係順に注意）
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.Seat{},
		&models.WorkTime{},
	); err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}

	fmt.Println("✅ Database migrated successfully.")
}
