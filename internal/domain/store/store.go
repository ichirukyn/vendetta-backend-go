package store

import (
	"vendetta/internal/domain/repositories"
)

type Store interface {
	Hero() repositories.HeroRepository
	HeroSpell() repositories.HeroSpellRepository
	HeroSpec() repositories.HeroSpecRepository
	HeroStatistics() repositories.HeroStatisticsRepository
	HeroStorage() repositories.HeroStorageRepository
	HeroTeam() repositories.HeroTeamRepository
	HeroWeapon() repositories.HeroWeaponRepository
	HeroTechnique() repositories.HeroTechniqueRepository

	Storage() repositories.StorageRepository
	StorageItem() repositories.StorageItemRepository

	Item() repositories.ItemRepository

	User() repositories.UserRepository
	Ban() repositories.BanRepository
}
