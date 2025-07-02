package usecases

import (
	"iruyan-api/models"
	"iruyan-api/repositories"
)

type AuthUsecase interface {
	RegisterUser(name, iruyanID, password, email string) (*models.User, error)
	LoginUser(iruyanID, password string) (*models.User, error)
	LogoutUser(iruyanID string) (*models.User, error)
}

type authUsecase struct {
	Repo repositories.UserRepositoryInterface
}

func (u *authUsecase) RegisterUser(name, iruyanID, password, email string) (*models.User, error) {
	return u.Repo.CreateUser(name, iruyanID, password, email)
}

func (u *authUsecase) LoginUser(iruyanID, password string) (*models.User, error) {
	user, err := u.Repo.FindByIruyanID(iruyanID)
	if err != nil {
		return nil, err
	}

	if err := user.CheckPassword(password); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *authUsecase) LogoutUser(iruyanID string) (*models.User, error) {
	user, err := u.Repo.FindByIruyanID(iruyanID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
