package repositories

import (
	"vendetta/internal/domain/entities"
)

type StorageRepository interface {
	Create(*entities.Storage) error
	GetByID(string) (*entities.Storage, error)
	Update(*entities.Storage) error
	Delete(*entities.Storage) error
}

type StorageItemRepository interface {
	Create(*entities.StorageItem) error
	GetAllByStorageID(string, *entities.Filter, *entities.WhereQueryFilter) ([]*entities.StorageItem, error)
	Update(*entities.StorageItem) error
	Delete(*entities.StorageItem) error
}
