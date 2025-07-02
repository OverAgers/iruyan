package repositories

import (
	"errors"
	"fmt"
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/pkg/errdefs"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	NewInstance() *models.User
	CreateUser(name, iruyanID, password, email string) (*models.User, error)
	Delete(user *models.User) error
	FindByID(userID string) (*models.User, error)
	FindByIruyanID(iruyanID string) (*models.User, error)
	DeleteByIruyanID(iruyanID string) error
	GetAllUsers() ([]models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func (r *userRepository) NewInstance() *models.User {
	return &models.User{}
}

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

func (r *userRepository) CreateUser(name, iruyanID, password, email string) (*models.User, error) {
	// 重複チェック
	var existing models.User
	if err := r.DB.Where("iruyan_id = ? OR email = ?", iruyanID, email).First(&existing).Error; err == nil {
		if existing.IruyanID == iruyanID {
			return nil, errdefs.ErrDuplicateIruyanID
		}
		if existing.Email == email {
			return nil, errdefs.ErrDuplicateEmail
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
	if err := r.DB.Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}

func (r *userRepository) Delete(user *models.User) error {
	if err := r.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

// FindByID - IDを元にUserが存在するかを検索するメソッド
func (r *userRepository) FindByID(userID uint) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// FindByIruyanID - IruyanID を元に User を検索するメソッド
func (r *userRepository) FindByIruyanID(iruyanID string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("iruyan_id = ?", iruyanID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found with iruyan_id: %s", iruyanID)
		}
		return nil, fmt.Errorf("failed to find user by iruyan_id (%s): %w", iruyanID, err)
	}
	return &user, nil
}

func (r *userRepository) DeleteByIruyanID(iruyanID string) error {
	user, err := r.FindByIruyanID(iruyanID)
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}

	if err := r.DB.Delete(user).Error; err != nil {
		return fmt.Errorf("failed to delete user with iruyan_id (%s): %w", iruyanID, err)
	}

	return nil
}

// GetAllUsers retrieves all users from the database.
func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
