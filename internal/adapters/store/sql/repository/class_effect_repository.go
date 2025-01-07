package repositories

import (
	"errors"
	"gorm.io/gorm"
	"vendetta/internal/domain/constants"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/store"
	"vendetta/internal/services/postgres"
)

type ClassEffectRepository struct {
	DB *postgres.Database
}

func (r *ClassEffectRepository) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.ClassEffect, error) {
	var classes []*entities.ClassEffect

	result := r.DB.
		Table(constants.DatabaseClassesEffects).
		Find(&classes).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Where(*query).
		Scan(&classes)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return classes, result.Error
}

func (r *ClassEffectRepository) GetByClassID(id string) ([]*entities.ClassEffect, error) {
	effects := []*entities.ClassEffect{}

	result := r.DB.Table(constants.DatabaseClassesEffects).Where("class_id = ?", id).Find(&effects).Scan(&effects)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return effects, result.Error
}
