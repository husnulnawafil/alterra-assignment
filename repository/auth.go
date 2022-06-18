package repository

import (
	"alterra/test/delivery/middlewares"
	"alterra/test/entities"
	"errors"

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
	auth := authRepo.Database.Where("email = ?", email).Find(&user)
	if auth.Error != nil {
		return "", auth.Error
	}
	if auth.RowsAffected == 0 {
		return "", errors.New("user not found")
	}

	token, err := middlewares.CreateToken(int(user.ID), user.Name, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}
