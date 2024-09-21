package factory

import (
	"vendetta/internal/adapters/http/rest/dto"
	"vendetta/internal/domain/entities"
)

// UserFactory фабрика для создания сущностей пользователя (entities.User)
type UserFactory struct {
}

// CreateUserFromCreateDTO создание пользователя из dto.CreateUserDTO
func (f *UserFactory) CreateUserFromCreateDTO(dto dto.CreateUserDTO) *entities.User {
	return &entities.User{
		ChatID: dto.ChatID,
		Role:   dto.Role,
		Login:  dto.Login,
	}
}

// CreateUserFromUpdateDTO создание пользователя из dto.UpdateUserDTO
func (f *UserFactory) CreateUserFromUpdateDTO(id string, dto dto.UpdateUserDTO) *entities.User {
	return &entities.User{
		ID:     id,
		ChatID: dto.ChatID,
		Role:   dto.Role,
		Login:  dto.Login,
	}
}

// NewUserFactory инициализатор UserFactory
func NewUserFactory() *UserFactory {
	return &UserFactory{}
}
