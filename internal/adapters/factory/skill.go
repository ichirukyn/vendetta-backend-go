package factory

import (
	"vendetta/internal/adapters/http/rest/dto"
	"vendetta/internal/domain/entities"
)

// SkillFactory фабрика для создания сущностей пользователя (entities.Skill)
type SkillFactory struct {
}

// CreateSkillFromCreateDTO создание пользователя из dto.CreateSkillDTO
func (f *SkillFactory) CreateSkillFromCreateDTO(dto dto.CreateSkillDTO) *entities.Skill {
	return &entities.Skill{
		Name:      dto.Name,
		Desc:      dto.Desc,
		DescShort: dto.DescShort,
		Rank:      dto.Rank,
		Damage:    dto.Damage,
		Element:   dto.Element,
		Distance:  dto.Distance,
		Type:      dto.Type,
		MainStat:  dto.MainStat,
		Cooldown:  dto.Cooldown,
		UserID:    dto.UserID,
		ClassID:   dto.ClassID,
		RaceID:    dto.RaceID,
		IsStack:   dto.IsStack,
		Hidden:    dto.Hidden,
	}
}

// CreateSkillFromUpdateDTO создание пользователя из dto.UpdateSkillDTO
func (f *SkillFactory) CreateSkillFromUpdateDTO(id string, dto dto.UpdateSkillDTO) *entities.Skill {
	return &entities.Skill{
		ID:        id,
		Name:      dto.Name,
		Desc:      dto.Desc,
		DescShort: dto.DescShort,
		Rank:      dto.Rank,
		Damage:    dto.Damage,
		Element:   dto.Element,
		Distance:  dto.Distance,
		Type:      dto.Type,
		MainStat:  dto.MainStat,
		Cooldown:  dto.Cooldown,
		ClassID:   dto.ClassID,
		RaceID:    dto.RaceID,
		IsStack:   dto.IsStack,
		Hidden:    dto.Hidden,
	}
}

// CreateSkillFromFindDTO создание пользователя из dto.FindSkillDTO
func (f *SkillFactory) CreateSkillFromFindDTO(id string, dto dto.FindSkillsDTO) *entities.Skill {
	return &entities.Skill{
		ID:       id,
		Name:     dto.Name,
		Rank:     dto.Rank,
		Element:  dto.Element,
		Distance: dto.Distance,
		Type:     dto.Type,
		MainStat: dto.MainStat,
		ClassID:  dto.ClassID,
		RaceID:   dto.RaceID,
		IsStack:  dto.IsStack,
		Hidden:   dto.Hidden,
	}
}

// NewSkillFactory инициализатор SkillFactory
func NewSkillFactory() *SkillFactory {
	return &SkillFactory{}
}
