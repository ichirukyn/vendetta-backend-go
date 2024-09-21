package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vendetta/internal/domain/entities"
)

type HeroRepository interface {
	Create(*entities.Hero) error
	GetAll(*entities.Filter, *entities.WhereQueryFilter) ([]*entities.Hero, error)
	GetByID(primitive.ObjectID) (*entities.Hero, error)
	GetByUserID(primitive.ObjectID) ([]*entities.Hero, error)
	Update(*entities.Hero) error
	Delete(*entities.Hero) error
}

type HeroSpellRepository interface {
	Create(*entities.HeroSpell) error
	GetAllByHeroID(primitive.ObjectID, *entities.Filter, *entities.WhereQueryFilter) ([]*entities.HeroSpell, error)
	GetBySpellID(primitive.ObjectID) ([]*entities.HeroSpell, error)
	Update(*entities.HeroSpell) error
	Delete(*entities.HeroSpell) error
}

type HeroStatisticsRepository interface {
	Create(*entities.HeroStatistics) error
	GetByHeroID(primitive.ObjectID) (*entities.HeroStatistics, error)
	Update(*entities.HeroStatistics) error
}

type HeroSpecRepository interface {
	Create(*entities.HeroSpec) error
	GetByHeroID(primitive.ObjectID) (*entities.HeroSpec, error)
	Update(*entities.HeroSpec) error
	Delete(*entities.HeroSpec) error
}

type HeroTeamRepository interface {
	Create(*entities.HeroTeam) error
	GetByHeroID(primitive.ObjectID) (*entities.HeroTeam, error)
	GetByTeamID(primitive.ObjectID) (*entities.HeroTeam, error)
	Update(*entities.HeroTeam) error
	Delete(*entities.HeroTeam) error
}

type HeroTechniqueRepository interface {
	Create(*entities.HeroTechnique) error
	GetAllByHeroID(primitive.ObjectID, *entities.Filter, *entities.WhereQueryFilter) ([]*entities.HeroTechnique, error)
	GetByTechniqueID(primitive.ObjectID) (*entities.HeroTechnique, error)
	Update(*entities.HeroTechnique) error
	Delete(*entities.HeroTechnique) error
}

type HeroWeaponRepository interface {
	Create(*entities.HeroWeapon) error
	GetAllByHeroID(primitive.ObjectID, *entities.Filter, *entities.WhereQueryFilter) ([]*entities.HeroWeapon, error)
	GetByWeaponID(primitive.ObjectID) (*entities.HeroWeapon, error)
	Update(*entities.HeroWeapon) error
	Delete(*entities.HeroWeapon) error
}

type HeroStorageRepository interface {
	Create(*entities.HeroStorage) error
	GetAllByHeroID(primitive.ObjectID, *entities.Filter, *entities.WhereQueryFilter) ([]*entities.HeroStorage, error)
	GetHeroInventory(primitive.ObjectID) (*entities.HeroStorage, error)
	GetByStorageID(primitive.ObjectID) (*entities.HeroStorage, error)
	Update(*entities.HeroStorage) error
	Delete(*entities.HeroStorage) error
}
