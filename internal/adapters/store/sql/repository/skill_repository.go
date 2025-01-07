package repositories

import (
	"errors"
	"gorm.io/gorm"
	"vendetta/internal/domain/constants"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/store"
	"vendetta/internal/services/postgres"
)

type SkillRepository struct {
	DB *postgres.Database
}

func (r *SkillRepository) Create(skill *entities.Skill) error {
	result := r.DB.Table(constants.DatabaseSkills).Create(&skill).Scan(&skill)
	if result.Error != nil {
		return store.ErrRecordNotCreated
	}

	return nil
}

func (r *SkillRepository) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.Skill, error) {
	var skills []*entities.Skill

	result := r.DB.
		Table(constants.DatabaseSkills).
		Find(&skills).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Where(*query).
		Scan(&skills)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return skills, result.Error
}

func (r *SkillRepository) GetByID(id string) (*entities.Skill, error) {
	skill := &entities.Skill{}

	result := r.DB.Table(constants.DatabaseSkills).Where("id = ?", id).First(skill).Scan(&skill)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return skill, result.Error
}

func (r *SkillRepository) Update(skill *entities.Skill) error {
	result := r.DB.Table(constants.DatabaseSkills).Save(&skill).Scan(&skill)
	if result.Error != nil {
		return store.ErrRecordNotUpdated
	}

	return nil
}

func (r *SkillRepository) Delete(skill *entities.Skill) error {
	result := r.DB.Table(constants.DatabaseSkills).Delete(&skill)
	if result.Error != nil {
		return store.ErrRecordNotDeleted
	}

	return nil
}
