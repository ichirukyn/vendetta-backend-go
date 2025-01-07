package repositories

import (
	"errors"
	"gorm.io/gorm"
	"vendetta/internal/domain/constants"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/store"
	"vendetta/internal/services/postgres"
)

type ClassRepository struct {
	DB *postgres.Database
}

func (r *ClassRepository) Create(class *entities.Class) error {
	result := r.DB.Table(constants.DatabaseClasses).Create(&class).Scan(&class)
	if result.Error != nil {
		return store.ErrRecordNotCreated
	}

	return nil
}

func (r *ClassRepository) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.Class, error) {
	var classs []*entities.Class

	result := r.DB.
		Table(constants.DatabaseClasses).
		Find(&classs).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Where(*query).
		Scan(&classs)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return classs, result.Error
}

func (r *ClassRepository) GetByID(id string) (*entities.Class, error) {
	class := &entities.Class{}

	result := r.DB.Table(constants.DatabaseClasses).Where("id = ?", id).First(class).Scan(&class)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return class, result.Error
}

func (r *ClassRepository) Update(class *entities.Class) error {
	result := r.DB.Table(constants.DatabaseClasses).Save(&class).Scan(&class)
	if result.Error != nil {
		return store.ErrRecordNotUpdated
	}

	return nil
}

func (r *ClassRepository) Delete(class *entities.Class) error {
	result := r.DB.Table(constants.DatabaseClasses).Delete(&class)
	if result.Error != nil {
		return store.ErrRecordNotDeleted
	}

	return nil
}
