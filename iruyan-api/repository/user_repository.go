package repository

import (
	"iruyan-api/infrastructure"
	"iruyan-api/models"

	"golang.org/x/crypto/bcrypt"
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
