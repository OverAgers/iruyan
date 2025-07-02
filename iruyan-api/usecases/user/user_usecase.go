package usecases

import (
	"errors"
	"iruyan-api/models"
	"iruyan-api/pkg/errdefs"
	"iruyan-api/repositories"
)

type UserUsecase interface {
	CheckUserExists(iruyanID string) error
	DeleteUser(iruyanID string) error
	GetAllUsers() ([]models.User, error)
	GetUserByIruyanID(iruyanId string) (*models.User, error)
}

type userUsecase struct {
	Repo repositories.UserRepositoryInterface
}

func (u *userUsecase) CheckUserExists(iruyanID string) error {
	_, err := u.Repo.FindByIruyanID(iruyanID)
	return err
}

func (u *userUsecase) DeleteUser(iruyanID string) error {
	user, err := u.Repo.FindByIruyanID(iruyanID)
	if err != nil {
		return err
	}
	return u.Repo.Delete(user)
}

func (u *userUsecase) GetAllUsers() ([]models.User, error) {
	return u.Repo.GetAllUsers()
}

func (u *userUsecase) GetUserByIruyanID(iruyanId string) (*models.User, error) {
	user, err := u.Repo.FindByIruyanID(iruyanId)
	if err != nil {
		if errors.Is(err, errdefs.ErrUserNotFound) {
			return nil, errdefs.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}
