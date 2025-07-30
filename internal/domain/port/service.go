package port

import "github.com/aprianfirlanda/go-server/internal/domain/model"

type UserService interface {
	FetchUsers() ([]model.User, error)
}
