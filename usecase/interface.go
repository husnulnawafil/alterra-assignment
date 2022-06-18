package usecase

import "alterra/test/entities"

type UserUseCaseInterface interface {
	CreateUser(newUser entities.User) (int, error)
	GetListUsers(search, role string) ([]entities.GetUserResponse, error)
	DeleteUser(userID int) error
	UpdateUser(user entities.User, userID int) error
	GetUserById(userID int) (entities.GetUserResponse, error)
}

type AuthUseCaseInterface interface {
	Login(email string, password string) (string, error)
}
