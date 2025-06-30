package repository

import (
	"errors"
	"fmt"
	"iruyan-api/infrastructure"
	"iruyan-api/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateTestUser is a helper for inserting test users into DB
func CreateTestUser(iruyanID, password string) error {
	// 事前に既存のテストユーザーを削除
	infrastructure.DB.Where("iruyan_id = ?", iruyanID).Delete(&models.User{})

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		IruyanID: iruyanID,
		Password: string(hashed),
		Name:     "テストユーザー",
		Email:    iruyanID + "@example.com",
	}
	return infrastructure.DB.Create(&user).Error
}

func CreateUser(db *gorm.DB, name, iruyanID, password, email string) (*models.User, error) {
	// 重複チェック
	var existing models.User
	if err := db.Where("iruyan_id = ? OR email = ?", iruyanID, email).First(&existing).Error; err == nil {
		if existing.IruyanID == iruyanID {
			return nil, ErrDuplicateIruyanID
		}
		if existing.Email == email {
			return nil, ErrDuplicateEmail
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	// 残りはモデルレイヤに任せる
	user, err := models.NewUser(name, iruyanID, password, email)
	if err != nil {
		return nil, err
	}

	// 保存
	if err := db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}
