package usecase

import (
	"github.com/aprianfirlanda/go-server/internal/domain/model"
	"github.com/aprianfirlanda/go-server/internal/domain/port"
)

type userUsecase struct {
	repo port.UserRepository
}

func NewUserUsecase(repo port.UserRepository) port.UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) CreateUser(user *model.User) error {
	return u.repo.Create(user)
}

func (u *userUsecase) GetUsers() ([]model.User, error) {
	return u.repo.FindAll()
}
