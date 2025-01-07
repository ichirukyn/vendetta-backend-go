package repositories

import (
	"errors"
	"gorm.io/gorm"
	"vendetta/internal/domain/constants"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/store"
	"vendetta/internal/services/postgres"
)

type RaceEffectRepository struct {
	DB *postgres.Database
}

func (r *RaceEffectRepository) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.RaceEffect, error) {
	var races []*entities.RaceEffect

	result := r.DB.
		Table(constants.DatabaseRacesEffects).
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

func (r *RaceEffectRepository) GetByRaceID(id string) ([]*entities.RaceEffect, error) {
	effects := []*entities.RaceEffect{}

	result := r.DB.Table(constants.DatabaseRacesEffects).Where("race_id = ?", id).Find(&effects).Scan(&effects)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return effects, result.Error
}
