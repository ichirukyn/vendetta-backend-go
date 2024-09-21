package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vendetta/internal/domain/entities"
)

type StorageRepository interface {
	Create(*entities.Storage) error
	GetByID(primitive.ObjectID) (*entities.Storage, error)
	Update(*entities.Storage) error
	Delete(*entities.Storage) error
}

type StorageItemRepository interface {
	Create(*entities.StorageItem) error
	GetAllByStorageID(primitive.ObjectID, *entities.Filter, *entities.WhereQueryFilter) ([]*entities.StorageItem, error)
	Update(*entities.StorageItem) error
	Delete(*entities.StorageItem) error
}
