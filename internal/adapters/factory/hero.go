package factory

import (
	"vendetta/internal/adapters/http/rest/dto"
	"vendetta/internal/domain/entities"
)

type HeroFactory struct {
}

func (f *HeroFactory) CreateHeroFromCreateDTO(dto dto.CreateHeroDTO) *entities.Hero {
	return &entities.Hero{
		UserID:   dto.UserID,
		RaceID:   dto.RaceID,
		ClassID:  dto.ClassID,
		Name:     dto.Name,
		Rank:     dto.Rank,
		IsHidden: dto.IsHidden,
		Level:    dto.Level,
		Exp:      dto.Exp,
	}
}

func (f *HeroFactory) CreateHeroFromUpdateDTO(id string, dto dto.UpdateHeroDTO) *entities.Hero {
	return &entities.Hero{
		ID:       id,
		Name:     dto.Name,
		Rank:     dto.Rank,
		IsHidden: dto.IsHidden,
		Level:    dto.Level,
		Exp:      dto.Exp,
	}
}

func (f *HeroFactory) FindHeroFromFindDTO(id string, dto dto.FindHeroesDTO) *entities.Hero {
	return &entities.Hero{
		ID:     id,
		UserID: dto.UserID,
		Name:   dto.Name,
	}
}

// NewHeroFactory инициализатор HeroFactory
func NewHeroFactory() *HeroFactory {
	return &HeroFactory{}
}
