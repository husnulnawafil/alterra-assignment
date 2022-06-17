package repository

import (
	"alterra/test/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (userRepo *UserRepository) CreateUser(newUser entities.User) error {
	err := userRepo.Database.Create(&newUser).Error
	return err
}
