package repository

import "alterra/test/entities"

type UserRepositoryInterface interface {
	CreateUser(newUser entities.User) error
	GetListUsers() ([]entities.User, error)
	DeleteUser(userID int) error
	UpdateUser(user entities.User) error
	GetUserById(userID int) (entities.User, error)
}

type AuthRepositoryInterface interface {
	Login(email string, password string) (string, error)
}
