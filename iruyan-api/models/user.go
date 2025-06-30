package models

import (
	"errors"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 構造体
type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"size:255;not null"`
	IruyanID string `gorm:"uniqueIndex;size:255;not null"`
	Password string `gorm:"size:255;not null"` // ハッシュ化されたパスワード
	Email    string `gorm:"uniqueIndex;size:255;not null"`
}

// NewUser: User構造体のコンストラクタ関数
func NewUser(name, iruyanID, password, email string) (*User, error) {
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	if err := validatePassword(password); err != nil {
		return nil, err
	}
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:     name,
		IruyanID: iruyanID,
		Password: hashedPassword,
		Email:    email,
	}, nil
}

// HashPassword: パスワードをハッシュ化する関数
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// validatePassword: パスワードが英数字を含む6文字以上かを確認する内部関数
func validatePassword(password string) error {
	// パスワードが少なくとも6文字以上であり、英字と数字がそれぞれ少なくとも1つ含まれているかをチェック
	if len(password) < 6 {
		return fmt.Errorf("password must be at least 6 characters long")
	}

	// 英字と数字が少なくとも1つずつ含まれているかを正規表現で確認
	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString

	if !hasLetter(password) || !hasNumber(password) {
		return fmt.Errorf("password must contain at least one letter and one number")
	}

	return nil
}

// validateEmail: 簡易的なメールアドレスバリデーション
func validateEmail(email string) error {
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	if !match {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

// CheckPassword 受け取ったプレーンテキストのパスワードをハッシュ化されたパスワードと比較するメソッド
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// GetAllUsers retrieves all users from the database.
func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// FindByID - IDを元にUserが存在するかを検索するメソッド
func (u *User) FindByID(db *gorm.DB, userID uint) error {
	if err := db.Where("id = ?", userID).First(u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}

// FindByIruyanID - IruyanID を元に User を検索するメソッド
func (u *User) FindByIruyanID(db *gorm.DB, iruyanID string) error {
	err := db.Where("iruyan_id = ?", iruyanID).First(u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user not found with iruyan_id: %s", iruyanID)
		}
		return fmt.Errorf("failed to find user by iruyan_id (%s): %w", iruyanID, err)
	}
	return nil
}

func (u *User) DeleteByIruyanID(db *gorm.DB, iruyanID string) error {
	// まず削除対象のユーザーを取得
	if err := u.FindByIruyanID(db, iruyanID); err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}

	// 見つかったユーザーを削除
	if err := db.Delete(u).Error; err != nil {
		return fmt.Errorf("failed to delete user with iruyan_id (%s): %w", iruyanID, err)
	}

	return nil
}
