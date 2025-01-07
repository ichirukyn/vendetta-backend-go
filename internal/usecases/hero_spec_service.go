package usecases

import (
	"vendetta/internal/app/config"
	"vendetta/internal/domain"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/repositories"
	"vendetta/pkg/utils"
)

type HeroSpecService struct {
	c *config.Config
	l *utils.Logger

	HeroSpecRepository repositories.HeroSpecRepository
}

func (s *HeroSpecService) Create(spec *entities.HeroSpec) *domain.AppError {
	if err := s.HeroSpecRepository.Create(spec); err != nil {
		return domain.NewBadRequestError("the hero_spec was not create")
	}

	return nil
}

func (s *HeroSpecService) GetByHeroID(specID string) (*entities.HeroSpec, *domain.AppError) {
	spec, err := s.HeroSpecRepository.GetByHeroID(specID)
	if err != nil {
		return nil, domain.NewNotFoundError("the hero_spec was not found")
	}

	return spec, nil
}

func (s *HeroSpecService) Update(spec *entities.HeroSpec) *domain.AppError {
	_, err := s.HeroSpecRepository.GetByHeroID(spec.HeroID)
	if err != nil {
		return domain.NewNotFoundError("the hero_spec was not found")
	}

	if err := s.HeroSpecRepository.Update(spec); err != nil {
		return domain.NewBadRequestError("the hero_spec was not updated")
	}

	return nil
}

func (s *HeroSpecService) Delete(id string) *domain.AppError {
	spec, err := s.HeroSpecRepository.GetByHeroID(id)
	if err != nil {
		return domain.NewNotFoundError("the hero_spec was not found")
	}

	if err := s.HeroSpecRepository.Delete(spec); err != nil {
		return domain.NewBadRequestError("the hero_spec was not deleted")
	}

	return nil
}

func NewHeroSpecService(c *config.Config, l *utils.Logger, specRepository repositories.HeroSpecRepository) *HeroSpecService {
	return &HeroSpecService{
		c: c,
		l: l,

		HeroSpecRepository: specRepository,
	}
}
