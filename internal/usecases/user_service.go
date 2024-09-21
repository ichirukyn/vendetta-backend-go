package usecases

import (
	"vendetta/internal/app/config"
	"vendetta/internal/domain"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/repositories"
	"vendetta/pkg/utils"
)

type UserService struct {
	c *config.Config
	l *utils.Logger

	UserRepository repositories.UserRepository
}

func (s *UserService) Create(user *entities.User) *domain.AppError {
	if err := s.UserRepository.Create(user); err != nil {
		return domain.NewBadRequestError("the user was not create")
	}

	return nil
}

func (s *UserService) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.User, *domain.AppError) {
	users, err := s.UserRepository.GetAll(filter, query)
	if err != nil || len(users) == 0 {
		return nil, domain.NewNotFoundError("the users was not found")
	}

	return users, nil
}

func (s *UserService) GetByID(id string) (*entities.User, *domain.AppError) {
	user, err := s.UserRepository.GetByID(id)
	if err != nil {
		return nil, domain.NewNotFoundError("the user was not found")
	}

	return user, nil
}

func (s *UserService) GetByChatID(chatID string) (*entities.User, *domain.AppError) {
	user, err := s.UserRepository.GetByChatID(chatID)
	if err != nil {
		return nil, domain.NewNotFoundError("the user was not found")
	}

	return user, nil
}

func (s *UserService) Update(user *entities.User) *domain.AppError {
	_, err := s.UserRepository.GetByID(user.ID)
	if err != nil {
		return domain.NewNotFoundError("the user was not found")
	}

	if err := s.UserRepository.Update(user); err != nil {
		return domain.NewBadRequestError("the user was not updated")
	}

	return nil
}

func (s *UserService) Delete(id string) *domain.AppError {
	user, err := s.UserRepository.GetByID(id)
	if err != nil {
		return domain.NewNotFoundError("the user was not found")
	}

	if err := s.UserRepository.Delete(user); err != nil {
		return domain.NewBadRequestError("the user was not deleted")
	}

	return nil
}

func NewUserService(c *config.Config, l *utils.Logger, userRepository repositories.UserRepository) *UserService {
	return &UserService{
		c: c,
		l: l,

		UserRepository: userRepository,
	}
}
