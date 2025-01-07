package factory

import (
	"vendetta/internal/adapters/http/rest/dto"
	"vendetta/internal/domain/entities"
)

// ClassFactory фабрика для создания сущностей пользователя (entities.Class)
type ClassFactory struct {
}

// CreateClassFromCreateDTO создание пользователя из dto.CreateClassDTO
func (f *ClassFactory) CreateClassFromCreateDTO(dto dto.CreateClassDTO) *entities.Class {
	return &entities.Class{
		ID:        dto.ID,
		Name:      dto.Name,
		Desc:      dto.Desc,
		DescShort: dto.DescShort,
		RaceID:    dto.RaceID,
		MainAttr:  dto.MainAttr,
		Type:      dto.Type,
		Hidden:    dto.Hidden,
		TagID:     dto.TagID,
	}
}

// CreateClassFromUpdateDTO создание пользователя из dto.UpdateClassDTO
func (f *ClassFactory) CreateClassFromUpdateDTO(id string, dto dto.UpdateClassDTO) *entities.Class {
	return &entities.Class{
		ID:        id,
		Name:      dto.Name,
		Desc:      dto.Desc,
		DescShort: dto.DescShort,
		RaceID:    dto.RaceID,
		MainAttr:  dto.MainAttr,
		Type:      dto.Type,
		Hidden:    dto.Hidden,
		TagID:     dto.TagID,
	}
}

// NewClassFactory инициализатор ClassFactory
func NewClassFactory() *ClassFactory {
	return &ClassFactory{}
}
