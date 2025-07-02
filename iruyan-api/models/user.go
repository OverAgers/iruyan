package models

import (
	"iruyan-api/pkg/errdefs"
	"regexp"

	"golang.org/x/crypto/bcrypt"
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
		return errdefs.ErrPasswordTooShort
	}

	// 英字と数字が少なくとも1つずつ含まれているかを正規表現で確認
	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString

	if !hasLetter(password) || !hasNumber(password) {
		return errdefs.ErrPasswordMissingChars
	}

	return nil
}

// validateEmail: 簡易的なメールアドレスバリデーション
func validateEmail(email string) error {
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	if !match {
		return errdefs.ErrInvalidEmail
	}
	return nil
}

// CheckPassword 受け取ったプレーンテキストのパスワードをハッシュ化されたパスワードと比較するメソッド
func (u *User) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return errdefs.ErrInvalidPassword
	}
	return nil
}
