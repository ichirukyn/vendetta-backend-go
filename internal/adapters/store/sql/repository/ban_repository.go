package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vendetta/internal/domain/entities"
)

type BanRepository interface {
	Create(*entities.Ban) error
	GetAll(*entities.Filter, *entities.WhereQueryFilter) ([]*entities.Ban, error)
	GetByUserID(primitive.ObjectID, *entities.Filter, *entities.WhereQueryFilter) ([]*entities.Ban, error)
	GetByChatID(string) (*entities.Ban, error)
	Update(*entities.Ban) error
	Delete(*entities.Ban) error
}
