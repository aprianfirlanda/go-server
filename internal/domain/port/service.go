package port

import "github.com/aprianfirlanda/go-server/internal/domain/model"

type UserUsecase interface {
	CreateUser(user *model.User) error
	GetUsers() ([]model.User, error)
}
