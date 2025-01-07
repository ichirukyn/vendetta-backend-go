package usecases

import (
	"vendetta/internal/app/config"
	"vendetta/internal/domain"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/repositories"
	"vendetta/pkg/utils"
)

type HeroService struct {
	c *config.Config
	l *utils.Logger

	HeroRepository     repositories.HeroRepository
	HeroSpecRepository repositories.HeroSpecRepository
}

func (s *HeroService) Create(hero *entities.Hero) *domain.AppError {
	hero, err := s.HeroRepository.Create(hero)

	if err != nil {
		return domain.NewBadRequestError("the hero was not create")
	}

	return nil
}

func (s *HeroService) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.Hero, *domain.AppError) {
	heroes, err := s.HeroRepository.GetAll(filter, query)
	if err != nil || len(heroes) == 0 {
		return nil, domain.NewNotFoundError("the heroes was not found")
	}

	return heroes, nil
}

func (s *HeroService) GetByID(id string) (*entities.Hero, *domain.AppError) {
	hero, err := s.HeroRepository.GetByID(id)
	if err != nil {
		return nil, domain.NewNotFoundError("the hero was not found")
	}

	return hero, nil
}

func (s *HeroService) GetByUserID(heroID string) ([]*entities.Hero, *domain.AppError) {
	hero, err := s.HeroRepository.GetByUserID(heroID)
	if err != nil {
		return nil, domain.NewNotFoundError("the hero was not found")
	}

	return hero, nil
}

func (s *HeroService) Update(hero *entities.Hero) *domain.AppError {
	_, err := s.HeroRepository.GetByID(hero.ID)
	if err != nil {
		return domain.NewNotFoundError("the hero was not found")
	}

	if err := s.HeroRepository.Update(hero); err != nil {
		return domain.NewBadRequestError("the hero was not updated")
	}

	return nil
}

func (s *HeroService) Delete(id string) *domain.AppError {
	hero, err := s.HeroRepository.GetByID(id)
	if err != nil {
		return domain.NewNotFoundError("the hero was not found")
	}

	if err := s.HeroRepository.Delete(hero); err != nil {
		return domain.NewBadRequestError("the hero was not deleted")
	}

	return nil
}

func NewHeroService(c *config.Config, l *utils.Logger, classRepository repositories.HeroRepository) *HeroService {
	return &HeroService{
		c: c,
		l: l,

		HeroRepository: classRepository,
	}
}
