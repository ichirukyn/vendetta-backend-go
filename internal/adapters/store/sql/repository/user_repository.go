package repositories

import (
	"errors"
	"gorm.io/gorm"
	"vendetta/internal/domain/constants"
	"vendetta/internal/domain/entities"
	"vendetta/internal/domain/store"
	"vendetta/internal/services/postgres"
)

type UserRepository struct {
	DB *postgres.Database
}

func (r *UserRepository) Create(user *entities.User) error {
	user.GenerateID()

	result := r.DB.Table(constants.DatabaseUsers).Create(&user).Scan(&user)
	if result.Error != nil {
		return store.ErrRecordNotCreated
	}

	return nil
}

func (r *UserRepository) GetAll(filter *entities.Filter, query *entities.WhereQueryFilter) ([]*entities.User, error) {
	var users []*entities.User

	result := r.DB.
		Table(constants.DatabaseUsers).
		Find(&users).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Where(*query).
		Scan(&users)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return users, result.Error
}

func (r *UserRepository) GetByID(id string) (*entities.User, error) {
	user := &entities.User{}

	result := r.DB.Table(constants.DatabaseUsers).Where("id = ?", id).First(user).Scan(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return user, result.Error
}

func (r *UserRepository) GetByChatID(chatID string) (*entities.User, error) {
	user := &entities.User{}

	result := r.DB.Table(constants.DatabaseUsers).Where("chat_id = ?", chatID).First(user).Scan(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return user, result.Error
}

func (r *UserRepository) Update(user *entities.User) error {
	result := r.DB.Table(constants.DatabaseUsers).Save(&user).Scan(&user)
	if result.Error != nil {
		return store.ErrRecordNotUpdated
	}

	return nil
}

func (r *UserRepository) Delete(user *entities.User) error {
	result := r.DB.Table(constants.DatabaseUsers).Delete(&user)
	if result.Error != nil {
		return store.ErrRecordNotDeleted
	}

	return nil
}
