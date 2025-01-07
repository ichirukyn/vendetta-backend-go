package usecases

import (
	"vendetta/internal/app/config"
	"vendetta/internal/domain"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/repositories"
	"vendetta/pkg/utils"
)

type RaceEffectService struct {
	c *config.Config
	l *utils.Logger

	RaceEffectRepository repositories.RaceEffectRepository
}

func (s *RaceEffectService) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.RaceEffect, *domain.AppError) {
	effects, err := s.RaceEffectRepository.GetAll(filter, query)
	if err != nil || len(effects) == 0 {
		return nil, domain.NewNotFoundError("the effects was not found")
	}

	return effects, nil
}

func (s *RaceEffectService) GetByID(id string) ([]*entities.RaceEffect, *domain.AppError) {
	raceEffects, err := s.RaceEffectRepository.GetByRaceID(id)
	if err != nil {
		return nil, domain.NewNotFoundError("the race effects was not found")
	}

	return raceEffects, nil
}

func NewRaceEffectService(c *config.Config, l *utils.Logger, classRepository repositories.RaceEffectRepository) *RaceEffectService {
	return &RaceEffectService{
		c: c,
		l: l,

		RaceEffectRepository: classRepository,
	}
}
