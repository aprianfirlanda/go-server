package port

import "github.com/aprianfirlanda/go-server/internal/domain/model"

type UserRepository interface {
	Create(user *model.User) error
	FindAll() ([]model.User, error)
}
