package repositories

import (
	"vendetta/internal/domain/entities"
)

type HeroRepository interface {
	Create(*entities.Hero) (*entities.Hero, error)
	GetAll(*entities.Filter, *entities.WhereQueryFilter) ([]*entities.Hero, error)
	GetByID(string) (*entities.Hero, error)
	GetByUserID(string) ([]*entities.Hero, error)
	Update(*entities.Hero) error
	Delete(*entities.Hero) error
}

type HeroSpecRepository interface {
	Create(*entities.HeroSpec) error
	GetByHeroID(string) (*entities.HeroSpec, error)
	Update(*entities.HeroSpec) error
	Delete(*entities.HeroSpec) error
}

type HeroSkillRepository interface {
	Create(*entities.HeroSkill) error
	GetAllByHeroID(string, *entities.Filter, *entities.WhereQueryFilter) ([]*entities.HeroSkill, error)
	GetBySpellID(string) ([]*entities.HeroSkill, error)
	Update(*entities.HeroSkill) error
	Delete(*entities.HeroSkill) error
}

// PAUSE
type HeroStatisticsRepository interface {
	Create(*entities.HeroStatistics) error
	GetByHeroID(string) (*entities.HeroStatistics, error)
	Update(*entities.HeroStatistics) error
}

type HeroTeamRepository interface {
	Create(*entities.HeroTeam) error
	GetByHeroID(string) (*entities.HeroTeam, error)
	GetByTeamID(string) (*entities.HeroTeam, error)
	Update(*entities.HeroTeam) error
	Delete(*entities.HeroTeam) error
}

type HeroWeaponRepository interface {
	Create(*entities.HeroWeapon) error
	GetAllByHeroID(string, *entities.Filter, *entities.WhereQueryFilter) ([]*entities.HeroWeapon, error)
	GetByWeaponID(string) (*entities.HeroWeapon, error)
	Update(*entities.HeroWeapon) error
	Delete(*entities.HeroWeapon) error
}

type HeroStorageRepository interface {
	Create(*entities.HeroStorage) error
	GetAllByHeroID(string, *entities.Filter, *entities.WhereQueryFilter) ([]*entities.HeroStorage, error)
	GetHeroInventory(string) (*entities.HeroStorage, error)
	GetByStorageID(string) (*entities.HeroStorage, error)
	Update(*entities.HeroStorage) error
	Delete(*entities.HeroStorage) error
}
