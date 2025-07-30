package port

import "github.com/aprianfirlanda/go-server/internal/domain/model"

type UserRepository interface {
	GetAll() ([]model.User, error)
}
