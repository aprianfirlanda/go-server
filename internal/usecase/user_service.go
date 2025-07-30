package usecase

import (
	"github.com/aprianfirlanda/go-server/internal/domain/model"
	"github.com/aprianfirlanda/go-server/internal/domain/port"
)

type UserServiceImpl struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) port.UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) FetchUsers() ([]model.User, error) {
	return s.repo.GetAll()
}
