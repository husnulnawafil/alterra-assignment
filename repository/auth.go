package repository

import (
	"alterra/test/delivery/middlewares"
	"alterra/test/entities"

	"gorm.io/gorm"
)

type AuthRepository struct {
	Database *gorm.DB
}

func NewAuthRepository(database *gorm.DB) *AuthRepository {
	return &AuthRepository{
		Database: database,
	}
}

func (authRepo *AuthRepository) Login(email string, password string) (string, error) {
	var user entities.User
	if err := authRepo.Database.Where("email = ?", email).Find(&user).Error; err != nil {
		return "", err
	}

	token, err := middlewares.CreateToken(int(user.ID), user.Name, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}
