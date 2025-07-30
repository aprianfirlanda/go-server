package repository

import (
	"github.com/aprianfirlanda/go-server/internal/domain/model"
)

type InMemoryUserRepo struct{}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{}
}

func (r *InMemoryUserRepo) GetAll() ([]model.User, error) {
	users := []model.User{
		{ID: 1, Name: "Aprian", Email: "aprian@example.com"},
		{ID: 2, Name: "Firlanda", Email: "firlanda@example.com"},
	}
	return users, nil
}
