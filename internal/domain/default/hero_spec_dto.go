package _default

import (
	"vendetta/internal/adapters/http/rest/dto"
)

func NewCreateHeroSpecDTO(heroID string) dto.CreateHeroSpecDTO {
	return dto.CreateHeroSpecDTO{
		HeroID:         heroID,
		Accuracy:       1,
		Strength:       1,
		Health:         1,
		Speed:          1,
		Dexterity:      1,
		Soul:           1,
		Intelligence:   1,
		Submissions:    1,
		CriticalRate:   0.05,
		CriticalDamage: 0.5,
		Resistance:     0.01,
		TotalSpec:      8,
		FreeSpec:       20,
	}
}
