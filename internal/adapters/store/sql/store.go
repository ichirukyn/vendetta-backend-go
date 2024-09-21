package sql

import (
	sql "vendetta/internal/adapters/store/sql/repository"
	"vendetta/internal/domain/repositories"
	"vendetta/internal/services/postgres"
)

type Store struct {
	DB *postgres.Database

	userRepository *sql.UserRepository
}

func (s *Store) Hero() repositories.HeroRepository {
	return nil
}

func (s *Store) HeroSpell() repositories.HeroSpellRepository {
	return nil
}

func (s *Store) HeroSpec() repositories.HeroSpecRepository {
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

func (s *Store) HeroTechnique() repositories.HeroTechniqueRepository {
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

func New(db *postgres.Database) *Store {
	return &Store{
		DB: db,
	}
}
