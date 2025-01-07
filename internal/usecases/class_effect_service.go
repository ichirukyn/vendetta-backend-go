package usecases

import (
	"vendetta/internal/app/config"
	"vendetta/internal/domain"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/repositories"
	"vendetta/pkg/utils"
)

type ClassEffectService struct {
	c *config.Config
	l *utils.Logger

	ClassEffectRepository repositories.ClassEffectRepository
}

func (s *ClassEffectService) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.ClassEffect, *domain.AppError) {
	effects, err := s.ClassEffectRepository.GetAll(filter, query)
	if err != nil || len(effects) == 0 {
		return nil, domain.NewNotFoundError("the effects was not found")
	}

	return effects, nil
}

func (s *ClassEffectService) GetByID(id string) ([]*entities.ClassEffect, *domain.AppError) {
	classEffects, err := s.ClassEffectRepository.GetByClassID(id)
	if err != nil {
		return nil, domain.NewNotFoundError("the class effects was not found")
	}

	return classEffects, nil
}

func NewClassEffectService(c *config.Config, l *utils.Logger, classRepository repositories.ClassEffectRepository) *ClassEffectService {
	return &ClassEffectService{
		c: c,
		l: l,

		ClassEffectRepository: classRepository,
	}
}
