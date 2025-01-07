package repositories

import (
	"vendetta/internal/domain/entities"
)

type RaceRepository interface {
	GetAll(*entities.Filter, *entities.WhereQueryFilter) ([]*entities.Race, error)
	GetByID(string) (*entities.Race, error)
}

type RaceEffectRepository interface {
	GetAll(*entities.Filter, *entities.WhereQueryFilter) ([]*entities.RaceEffect, error)
	GetByRaceID(string) ([]*entities.RaceEffect, error)
}
