package usecases

import (
	"vendetta/internal/app/config"
	"vendetta/internal/domain"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/repositories"
	"vendetta/pkg/utils"
)

type SkillService struct {
	c *config.Config
	l *utils.Logger

	SkillRepository repositories.SkillRepository
}

func (s *SkillService) Create(skill *entities.Skill) *domain.AppError {
	if err := s.SkillRepository.Create(skill); err != nil {
		return domain.NewBadRequestError("the skill was not create")
	}

	return nil
}

func (s *SkillService) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.Skill, *domain.AppError) {
	skills, err := s.SkillRepository.GetAll(filter, query)
	if err != nil || len(skills) == 0 {
		return nil, domain.NewNotFoundError("the skills was not found")
	}

	return skills, nil
}

func (s *SkillService) GetByID(id string) (*entities.Skill, *domain.AppError) {
	skill, err := s.SkillRepository.GetByID(id)
	if err != nil {
		return nil, domain.NewNotFoundError("the skill was not found")
	}

	return skill, nil
}

func (s *SkillService) Update(skill *entities.Skill) *domain.AppError {
	_, err := s.SkillRepository.GetByID(skill.ID)
	if err != nil {
		return domain.NewNotFoundError("the skill was not found")
	}

	if err := s.SkillRepository.Update(skill); err != nil {
		return domain.NewBadRequestError("the skill was not updated")
	}

	return nil
}

func (s *SkillService) Delete(id string) *domain.AppError {
	skill, err := s.SkillRepository.GetByID(id)
	if err != nil {
		return domain.NewNotFoundError("the skill was not found")
	}

	if err := s.SkillRepository.Delete(skill); err != nil {
		return domain.NewBadRequestError("the skill was not deleted")
	}

	return nil
}

func NewSkillService(c *config.Config, l *utils.Logger, skillRepository repositories.SkillRepository) *SkillService {
	return &SkillService{
		c: c,
		l: l,

		SkillRepository: skillRepository,
	}
}
