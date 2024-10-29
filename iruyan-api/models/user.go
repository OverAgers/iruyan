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
	Username string `gorm:"uniqueIndex;size:255;not null"`
	Password string `gorm:"size:255;not null"` // ハッシュ化されたパスワード
	Task     string `gorm:"size:255;default:''"`
	Email    string `gorm:"uniqueIndex;size:255;not null"`
}

// NewUser: User構造体のコンストラクタ関数
func NewUser(db *gorm.DB, name, username, password, email string) (*User, error) {
	// ユーザーネームの重複チェック
	var existingUser User
	if err := db.Where("username = ?", username).Or("email = ?", email).First(&existingUser).Error; err == nil {
		if existingUser.Username == username {
			return nil, fmt.Errorf("username '%s' is already taken", username)
		}
		if existingUser.Email == email {
			return nil, fmt.Errorf("email '%s' is already registered", email)
		}
	}

	// メールアドレスのバリデーション
	if err := validateEmail(email); err != nil {
		return nil, err
	}

	// パスワードのバリデーション
	if err := validatePassword(password); err != nil {
		return nil, err
	}

	// パスワードをハッシュ化
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	// 新しいUserインスタンスを生成し、ハッシュ化されたパスワードを設定
	return &User{
		Name:     name,
		Username: username,
		Password: hashedPassword,
		Task:     "",
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

// FindByID - IDを元にUserが存在するかを検索するメソッド
func (u *User) FindByID(db *gorm.DB, userID uint) error {
	if err := db.Where("id = ?", userID).First(u).Error; err != nil {
		return errors.New("user not found")
	}
	return nil
}