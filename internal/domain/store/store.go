package store

import (
	"vendetta/internal/domain/repositories"
)

type Store interface {
	Hero() repositories.HeroRepository
	HeroSkill() repositories.HeroSkillRepository
	HeroSpec() repositories.HeroSpecRepository
	HeroStatistics() repositories.HeroStatisticsRepository
	HeroStorage() repositories.HeroStorageRepository
	HeroTeam() repositories.HeroTeamRepository
	HeroWeapon() repositories.HeroWeaponRepository

	Storage() repositories.StorageRepository
	StorageItem() repositories.StorageItemRepository

	Item() repositories.ItemRepository

	User() repositories.UserRepository
	Ban() repositories.BanRepository

	Race() repositories.RaceRepository
	RaceEffect() repositories.RaceEffectRepository
	Class() repositories.ClassRepository
	ClassEffect() repositories.ClassEffectRepository

	Skill() repositories.SkillRepository
}
