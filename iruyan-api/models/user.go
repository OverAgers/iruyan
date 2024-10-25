package models

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User 構造体
type User struct {
	ID       uuid.UUID
	Name     string
	Username string
	password string // 非公開フィールド（ハッシュ化されたパスワード）
	Icon     string
	Task     string
	Status   string
	Email    string
}

// NewUser: User構造体のコンストラクタ関数
func NewUser(name, username, password, icon, task, status, email string) (*User, error) {
	// パスワードのバリデーション
	if err := validatePassword(password); err != nil {
		return nil, err
	}

	// メールアドレスのバリデーション
	if err := validateEmail(email); err != nil {
		return nil, err
	}

	// パスワードをハッシュ化
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	// 新しいUserインスタンスを生成し、ハッシュ化されたパスワードを設定
	return &User{
		ID:       uuid.New(),
		Name:     name,
		Username: username,
		password: hashedPassword, // ハッシュ化済みのパスワード
		Icon:     icon,
		Task:     task,
		Status:   status,
		Email:    email,
	}, nil
}

// HashPassword: パスワードをハッシュ化する関数
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// validatePassword: パスワードが4桁の数字かを確認する内部関数
func validatePassword(password string) error {
	match, _ := regexp.MatchString(`^\d{4}$`, password)
	if !match {
		return fmt.Errorf("Password must be exactly 4 digits")
	}
	return nil
}

// validateEmail: 簡易的なメールアドレスバリデーション
func validateEmail(email string) error {
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	if !match {
		return fmt.Errorf("Invalid email format")
	}
	return nil
}

// CheckPasswordHash: パスワードがハッシュと一致するかをチェックする関数
func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}

// 仮想的なデータベースからユーザーを取得する関数
func GetUserByUsername(username string) (*User, error) {
	// ここでデータベース接続し、指定されたusernameを持つユーザーを検索する
	// 以下は仮の例です：

	// 仮のユーザー（データベースから取得したものと想定）
	if username == "john_doe" {
		// ハッシュ化されたパスワード（"1234"のハッシュ化例）
		hashedPassword := "$2a$10$7QxHj7rE3hVtZMQTZWnTEOvGTSNOjZxBoPlODhO.Oz9ypEYbfG8O6"
		return &User{
			Username: "john_doe",
			password: hashedPassword, // ハッシュ化されたパスワード
		}, nil
	}

	return nil, fmt.Errorf("user not found")
}
