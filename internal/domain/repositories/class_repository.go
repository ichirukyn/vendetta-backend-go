package repositories

import "vendetta/internal/domain/entities"

type ClassRepository interface {
	Create(*entities.Class) error
	GetAll(*entities.Filter, *entities.WhereQueryFilter) ([]*entities.Class, error)
	GetByID(string) (*entities.Class, error)
	Update(*entities.Class) error
	Delete(*entities.Class) error
}

type ClassEffectRepository interface {
	GetAll(*entities.Filter, *entities.WhereQueryFilter) ([]*entities.ClassEffect, error)
	GetByClassID(string) ([]*entities.ClassEffect, error)
}
