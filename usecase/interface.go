package usecase

import "alterra/test/entities"

type UserUseCaseInterface interface {
	CreateUser(newUser entities.User) (int, error)
	GetListUsers() ([]entities.User, error)
	DeleteUser(userID int) error
}
