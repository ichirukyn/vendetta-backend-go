package repositories

import "vendetta/internal/domain/entities"

type TagRepository interface {
	Create(*entities.Tag) error
	GetAll(*entities.Filter, *entities.WhereQueryFilter) ([]*entities.Tag, error)
	GetByID(string) (*entities.Tag, error)
	Update(*entities.Tag) error
	Delete(*entities.Tag) error
}
