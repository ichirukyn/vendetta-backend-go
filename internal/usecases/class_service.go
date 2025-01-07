package usecases

import (
	"vendetta/internal/app/config"
	"vendetta/internal/domain"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/repositories"
	"vendetta/pkg/utils"
)

type ClassService struct {
	c *config.Config
	l *utils.Logger

	ClassRepository repositories.ClassRepository
}

func (s *ClassService) Create(class *entities.Class) *domain.AppError {
	if err := s.ClassRepository.Create(class); err != nil {
		return domain.NewBadRequestError("the class was not create")
	}

	return nil
}

func (s *ClassService) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.Class, *domain.AppError) {
	classes, err := s.ClassRepository.GetAll(filter, query)
	if err != nil || len(classes) == 0 {
		return nil, domain.NewNotFoundError("the classes was not found")
	}

	return classes, nil
}

func (s *ClassService) GetByID(id string) (*entities.Class, *domain.AppError) {
	class, err := s.ClassRepository.GetByID(id)
	if err != nil {
		return nil, domain.NewNotFoundError("the class was not found")
	}

	return class, nil
}

func (s *ClassService) Update(class *entities.Class) *domain.AppError {
	_, err := s.ClassRepository.GetByID(class.ID)
	if err != nil {
		return domain.NewNotFoundError("the class was not found")
	}

	if err := s.ClassRepository.Update(class); err != nil {
		return domain.NewBadRequestError("the class was not updated")
	}

	return nil
}

func (s *ClassService) Delete(id string) *domain.AppError {
	class, err := s.ClassRepository.GetByID(id)
	if err != nil {
		return domain.NewNotFoundError("the class was not found")
	}

	if err := s.ClassRepository.Delete(class); err != nil {
		return domain.NewBadRequestError("the class was not deleted")
	}

	return nil
}

func NewClassService(c *config.Config, l *utils.Logger, classRepository repositories.ClassRepository) *ClassService {
	return &ClassService{
		c: c,
		l: l,

		ClassRepository: classRepository,
	}
}
