package repositories

import (
	"vendetta/internal/domain/entities"
)

type ItemRepository interface {
	Create(*entities.User) error
	GetAll(*entities.Filter, *entities.WhereQueryFilter) ([]*entities.User, error)
	GetByID(string) (*entities.User, error)
	GetByChatID(string) (*entities.User, error)
	Update(*entities.User) error
	Delete(*entities.User) error
}
