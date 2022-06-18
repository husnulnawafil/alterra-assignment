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

func (userRepo *UserRepository) GetListUsers() ([]entities.User, error) {
	var listUser []entities.User
	err := userRepo.Database.Find(&listUser).Error
	return listUser, err
}

func (userRepo *UserRepository) DeleteUser(userID int) error {
	err := userRepo.Database.Delete(&entities.User{}, userID).Error
	return err
}

func (userRepo *UserRepository) UpdateUser(user entities.User) error {
	err := userRepo.Database.Save(&user).Error
	return err
}

func (userRepo *UserRepository) GetUserById(userID int) (entities.User, error) {
	var user entities.User
	err := userRepo.Database.Where("id = ?", userID).Find(&user).Error
	return user, err
}
