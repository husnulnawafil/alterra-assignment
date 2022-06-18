package usecase

import "alterra/test/entities"

type UserUseCaseInterface interface {
	CreateUser(newUser entities.User) (int, error)
	GetListUsers() ([]entities.User, error)
	DeleteUser(userID int) error
	UpdateUser(user entities.User, userID int) error
	GetUserById(userID int) (entities.User, error)
}

type AuthUseCaseInterface interface {
	Login(email string, password string) (string, error)
}
