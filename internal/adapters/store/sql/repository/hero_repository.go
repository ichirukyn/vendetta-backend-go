package repositories

import (
	"errors"
	"gorm.io/gorm"
	"vendetta/internal/domain/constants"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/store"
	"vendetta/internal/services/postgres"
)

type HeroRepository struct {
	DB *postgres.Database
}

func (r *HeroRepository) Create(hero *entities.Hero) (*entities.Hero, error) {
	h := &entities.Hero{}

	result := r.DB.Table(constants.DatabaseHeroes).Create(&hero).Scan(&h)

	if result.Error != nil {
		return nil, store.ErrRecordNotCreated
	}

	return h, result.Error
}

func (r *HeroRepository) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.Hero, error) {
	var heroes []*entities.Hero

	result := r.DB.
		Table(constants.DatabaseHeroes).
		Find(&heroes).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Where(*query).
		Scan(&heroes)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return heroes, result.Error
}

func (r *HeroRepository) GetByID(id string) (*entities.Hero, error) {
	hero := &entities.Hero{}

	result := r.DB.Table(constants.DatabaseHeroes).Where("id = ?", id).First(hero).Scan(&hero)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return hero, result.Error
}

func (r *HeroRepository) GetByUserID(userID string) ([]*entities.Hero, error) {
	var heroes []*entities.Hero
	result := r.DB.Table(constants.DatabaseHeroes).
		Where("user_id = ?", userID).
		Find(&heroes)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return heroes, result.Error
}

func (r *HeroRepository) Update(hero *entities.Hero) error {
	result := r.DB.Table(constants.DatabaseHeroes).Save(&hero).Scan(&hero)
	if result.Error != nil {
		return store.ErrRecordNotUpdated
	}

	return nil
}

func (r *HeroRepository) Delete(hero *entities.Hero) error {
	result := r.DB.Table(constants.DatabaseHeroes).Delete(&hero)
	if result.Error != nil {
		return store.ErrRecordNotDeleted
	}

	return nil
}
