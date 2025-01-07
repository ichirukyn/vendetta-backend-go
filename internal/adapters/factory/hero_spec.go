package factory

import (
	"vendetta/internal/adapters/http/rest/dto"
	"vendetta/internal/domain/entities"
)

type HeroSpecFactory struct {
}

func (f *HeroSpecFactory) CreateHeroSpecFromCreateDTO(dto dto.CreateHeroSpecDTO) *entities.HeroSpec {
	return &entities.HeroSpec{
		HeroID:         dto.HeroID,
		Accuracy:       dto.Accuracy,
		Strength:       dto.Strength,
		Health:         dto.Health,
		Speed:          dto.Speed,
		Dexterity:      dto.Dexterity,
		Soul:           dto.Soul,
		Intelligence:   dto.Intelligence,
		Submissions:    dto.Submissions,
		CriticalRate:   dto.CriticalRate,
		CriticalDamage: dto.CriticalDamage,
		Resistance:     dto.Resistance,
		TotalSpec:      dto.TotalSpec,
		FreeSpec:       dto.FreeSpec,
	}
}

func (f *HeroSpecFactory) CreateHeroSpecFromUpdateDTO(id string, dto dto.UpdateHeroSpecDTO) *entities.HeroSpec {
	return &entities.HeroSpec{
		HeroID:         id,
		Accuracy:       dto.Accuracy,
		Strength:       dto.Strength,
		Health:         dto.Health,
		Speed:          dto.Speed,
		Dexterity:      dto.Dexterity,
		Soul:           dto.Soul,
		Intelligence:   dto.Intelligence,
		Submissions:    dto.Submissions,
		CriticalRate:   dto.CriticalRate,
		CriticalDamage: dto.CriticalDamage,
		Resistance:     dto.Resistance,
		TotalSpec:      dto.TotalSpec,
		FreeSpec:       dto.FreeSpec,
	}
}

func (f *HeroSpecFactory) FindHeroSpecFromFindDTO(id string, dto dto.FindHeroSpecDTO) *entities.HeroSpec {
	return &entities.HeroSpec{
		ID:     id,
		HeroID: dto.HeroID,
	}
}

// NewHeroSpecFactory инициализатор HeroSpecFactory
func NewHeroSpecFactory() *HeroSpecFactory {
	return &HeroSpecFactory{}
}
