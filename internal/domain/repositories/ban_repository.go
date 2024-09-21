package repositories

import (
	"vendetta/internal/domain/entities"
)

type BanRepository interface {
	Create(*entities.Ban) error
	GetAll(*entities.Filter, *entities.WhereQueryFilter) ([]*entities.Ban, error)
	GetByUserID(string, *entities.Filter, *entities.WhereQueryFilter) ([]*entities.Ban, error)
	GetByChatID(string) (*entities.Ban, error)
	Update(*entities.Ban) error
	Delete(*entities.Ban) error
}
