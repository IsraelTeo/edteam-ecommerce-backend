package user

import (
	"edteam-ecommerce-backend/model"
)

// Puerto
type UseCase interface {
	Create(m *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}

type Storage interface {
	Create(m *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}
