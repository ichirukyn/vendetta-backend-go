package repositories

import (
	"vendetta/internal/domain/entities"
)

type HeroRepository interface {
	Create(*entities.Hero) error
	GetAll(*entities.Filter, *entities.WhereQueryFilter) ([]*entities.Hero, error)
	GetByID(string) (*entities.Hero, error)
	GetByUserID(string) ([]*entities.Hero, error)
	Update(*entities.Hero) error
	Delete(*entities.Hero) error
}

type HeroSpellRepository interface {
	Create(*entities.HeroSpell) error
	GetAllByHeroID(string, *entities.Filter, *entities.WhereQueryFilter) ([]*entities.HeroSpell, error)
	GetBySpellID(string) ([]*entities.HeroSpell, error)
	Update(*entities.HeroSpell) error
	Delete(*entities.HeroSpell) error
}

type HeroStatisticsRepository interface {
	Create(*entities.HeroStatistics) error
	GetByHeroID(string) (*entities.HeroStatistics, error)
	Update(*entities.HeroStatistics) error
}

type HeroSpecRepository interface {
	Create(*entities.HeroSpec) error
	GetByHeroID(string) (*entities.HeroSpec, error)
	Update(*entities.HeroSpec) error
	Delete(*entities.HeroSpec) error
}

type HeroTeamRepository interface {
	Create(*entities.HeroTeam) error
	GetByHeroID(string) (*entities.HeroTeam, error)
	GetByTeamID(string) (*entities.HeroTeam, error)
	Update(*entities.HeroTeam) error
	Delete(*entities.HeroTeam) error
}

type HeroTechniqueRepository interface {
	Create(*entities.HeroTechnique) error
	GetAllByHeroID(string, *entities.Filter, *entities.WhereQueryFilter) ([]*entities.HeroTechnique, error)
	GetByTechniqueID(string) (*entities.HeroTechnique, error)
	Update(*entities.HeroTechnique) error
	Delete(*entities.HeroTechnique) error
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
