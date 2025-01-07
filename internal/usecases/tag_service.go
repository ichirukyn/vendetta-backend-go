package usecases

import (
	"vendetta/internal/app/config"
	"vendetta/internal/domain"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/repositories"
	"vendetta/pkg/utils"
)

type TagService struct {
	c *config.Config
	l *utils.Logger

	TagRepository repositories.TagRepository
}

func (s *TagService) Create(tag *entities.Tag) *domain.AppError {
	if err := s.TagRepository.Create(tag); err != nil {
		return domain.NewBadRequestError("the tag was not create")
	}

	return nil
}

func (s *TagService) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.Tag, *domain.AppError) {
	tags, err := s.TagRepository.GetAll(filter, query)
	if err != nil || len(tags) == 0 {
		return nil, domain.NewNotFoundError("the tags was not found")
	}

	return tags, nil
}

func (s *TagService) GetByID(id string) (*entities.Tag, *domain.AppError) {
	tag, err := s.TagRepository.GetByID(id)
	if err != nil {
		return nil, domain.NewNotFoundError("the tag was not found")
	}

	return tag, nil
}

func (s *TagService) Update(tag *entities.Tag) *domain.AppError {
	_, err := s.TagRepository.GetByID(tag.ID)
	if err != nil {
		return domain.NewNotFoundError("the tag was not found")
	}

	if err := s.TagRepository.Update(tag); err != nil {
		return domain.NewBadRequestError("the tag was not updated")
	}

	return nil
}

func (s *TagService) Delete(id string) *domain.AppError {
	tag, err := s.TagRepository.GetByID(id)
	if err != nil {
		return domain.NewNotFoundError("the tag was not found")
	}

	if err := s.TagRepository.Delete(tag); err != nil {
		return domain.NewBadRequestError("the tag was not deleted")
	}

	return nil
}
func NewTagService(c *config.Config, l *utils.Logger, tagRepository repositories.TagRepository) *TagService {
	return &TagService{
		c: c,
		l: l,

		TagRepository: tagRepository,
	}
}
