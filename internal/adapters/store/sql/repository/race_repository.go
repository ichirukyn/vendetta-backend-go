package repositories

import (
	"errors"
	"gorm.io/gorm"
	"vendetta/internal/domain/constants"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/store"
	"vendetta/internal/services/postgres"
)

type RaceRepository struct {
	DB *postgres.Database
}

func (r *RaceRepository) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.Race, error) {
	var races []*entities.Race

	result := r.DB.
		Table(constants.DatabaseRaces).
		Find(&races).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Where(*query).
		Scan(&races)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return races, result.Error
}

func (r *RaceRepository) GetByID(id string) (*entities.Race, error) {
	user := &entities.Race{}

	result := r.DB.Table(constants.DatabaseRaces).Where("id = ?", id).First(user).Scan(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return user, result.Error
}
