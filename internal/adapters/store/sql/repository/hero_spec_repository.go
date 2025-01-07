package repositories

import (
	"errors"
	"gorm.io/gorm"
	"vendetta/internal/domain/constants"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/store"
	"vendetta/internal/services/postgres"
)

type HeroSpecRepository struct {
	DB *postgres.Database
}

func (r *HeroSpecRepository) Create(spec *entities.HeroSpec) error {
	result := r.DB.Table(constants.DatabaseHeroesSpecs).Create(&spec).Scan(&spec)
	if result.Error != nil {
		return store.ErrRecordNotCreated
	}

	return nil
}

func (r *HeroSpecRepository) GetByHeroID(heroID string) (*entities.HeroSpec, error) {
	var spec entities.HeroSpec

	result := r.DB.Table(constants.DatabaseHeroesSpecs).Where("hero_id = ?", heroID).First(&spec)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return &spec, result.Error
}

func (r *HeroSpecRepository) Update(hero *entities.HeroSpec) error {
	result := r.DB.Table(constants.DatabaseHeroesSpecs).Save(&hero).Scan(&hero)
	if result.Error != nil {
		return store.ErrRecordNotUpdated
	}

	return nil
}

func (r *HeroSpecRepository) Delete(hero *entities.HeroSpec) error {
	result := r.DB.Table(constants.DatabaseHeroesSpecs).Delete(&hero)
	if result.Error != nil {
		return store.ErrRecordNotDeleted
	}

	return nil
}
