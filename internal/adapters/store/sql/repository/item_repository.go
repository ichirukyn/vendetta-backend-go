package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vendetta/internal/domain/entities"
)

type ItemRepository interface {
	Create(*entities.User) error
	GetAll(*entities.Filter, *entities.WhereQueryFilter) ([]*entities.User, error)
	GetByID(primitive.ObjectID) (*entities.User, error)
	GetByChatID(string) (*entities.User, error)
	Update(*entities.User) error
	Delete(*entities.User) error
}
