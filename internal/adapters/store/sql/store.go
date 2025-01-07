package sql

import (
	sql "vendetta/internal/adapters/store/sql/repository"
	"vendetta/internal/domain/repositories"
	"vendetta/internal/services/postgres"
)

type Store struct {
	DB *postgres.Database

	heroRepository     *sql.HeroRepository
	heroSpecRepository *sql.HeroSpecRepository

	userRepository        *sql.UserRepository
	raceRepository        *sql.RaceRepository
	raceEffectRepository  *sql.RaceEffectRepository
	classRepository       *sql.ClassRepository
	classEffectRepository *sql.ClassEffectRepository
	skillRepository       *sql.SkillRepository
}

func (s *Store) Hero() repositories.HeroRepository {
	if s.heroRepository != nil {
		return s.heroRepository
	}

	s.heroRepository = &sql.HeroRepository{
		DB: s.DB,
	}

	return s.heroRepository
}

func (s *Store) HeroSpec() repositories.HeroSpecRepository {
	if s.heroSpecRepository != nil {
		return s.heroSpecRepository
	}

	s.heroSpecRepository = &sql.HeroSpecRepository{
		DB: s.DB,
	}

	return s.heroSpecRepository
}

func (s *Store) HeroSkill() repositories.HeroSkillRepository {
	return nil
}

func (s *Store) HeroStatistics() repositories.HeroStatisticsRepository {
	return nil
}

func (s *Store) HeroStorage() repositories.HeroStorageRepository {
	return nil
}

func (s *Store) HeroTeam() repositories.HeroTeamRepository {
	return nil
}

func (s *Store) HeroWeapon() repositories.HeroWeaponRepository {
	return nil
}

func (s *Store) Storage() repositories.StorageRepository {
	return nil
}

func (s *Store) StorageItem() repositories.StorageItemRepository {
	return nil
}

func (s *Store) Item() repositories.ItemRepository {
	return nil
}

func (s *Store) User() repositories.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &sql.UserRepository{
		DB: s.DB,
	}

	return s.userRepository
}

func (s *Store) Ban() repositories.BanRepository {
	return nil
}

func (s *Store) Race() repositories.RaceRepository {
	if s.raceRepository != nil {
		return s.raceRepository
	}

	s.raceRepository = &sql.RaceRepository{
		DB: s.DB,
	}

	return s.raceRepository
}

func (s *Store) RaceEffect() repositories.RaceEffectRepository {
	if s.raceEffectRepository != nil {
		return s.raceEffectRepository
	}

	s.raceEffectRepository = &sql.RaceEffectRepository{
		DB: s.DB,
	}

	return s.raceEffectRepository
}

func (s *Store) Class() repositories.ClassRepository {
	if s.classRepository != nil {
		return s.classRepository
	}

	s.classRepository = &sql.ClassRepository{
		DB: s.DB,
	}

	return s.classRepository
}

func (s *Store) ClassEffect() repositories.ClassEffectRepository {
	if s.classEffectRepository != nil {
		return s.classEffectRepository
	}

	s.classEffectRepository = &sql.ClassEffectRepository{
		DB: s.DB,
	}

	return s.classEffectRepository
}

func (s *Store) Skill() repositories.SkillRepository {
	if s.skillRepository != nil {
		return s.skillRepository
	}

	s.skillRepository = &sql.SkillRepository{
		DB: s.DB,
	}

	return s.skillRepository
}

func New(db *postgres.Database) *Store {
	return &Store{
		DB: db,
	}
}
