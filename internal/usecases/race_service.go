package usecases

import (
	"vendetta/internal/app/config"
	"vendetta/internal/domain"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/repositories"
	"vendetta/pkg/utils"
)

type RaceService struct {
	c *config.Config
	l *utils.Logger

	RaceRepository repositories.RaceRepository
}

func (s *RaceService) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.Race, *domain.AppError) {
	races, err := s.RaceRepository.GetAll(filter, query)
	if err != nil || len(races) == 0 {
		return nil, domain.NewNotFoundError("the race was not found")
	}

	return races, nil
}

func (s *RaceService) GetByID(id string) (*entities.Race, *domain.AppError) {
	race, err := s.RaceRepository.GetByID(id)
	if err != nil {
		return nil, domain.NewNotFoundError("the race was not found")
	}

	return race, nil
}

func NewRaceService(c *config.Config, l *utils.Logger, classRepository repositories.RaceRepository) *RaceService {
	return &RaceService{
		c: c,
		l: l,

		RaceRepository: classRepository,
	}
}
