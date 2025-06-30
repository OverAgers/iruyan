package infrastructure

import (
	"fmt"
	"iruyan-api/models"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

// InitTestDB sets up an in-memory SQLite DB for testing.
func InitTestDB(t *testing.T) {
	var err error
	TestDB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("❌ Failed to connect to in-memory test database: %v", err)
	}

	// テスト中に使うモデルだけを migrate
	err = TestDB.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.Seat{},
		&models.WorkTime{},
	)
	if err != nil {
		t.Fatalf("❌ Failed to migrate test schema: %v", err)
	}

	// 通常の DB も上書きして使えるようにしておく
	DB = TestDB
	fmt.Fprintln(os.Stderr, "✅ TestDB initialized and migrated")
}

func SetNilDB() {
	DB = nil
}
