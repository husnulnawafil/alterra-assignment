package repository

import "alterra/test/entities"

type UserRepositoryInterface interface {
	CreateUser(newUser entities.User) error
	GetListUsers() ([]entities.User, error)
	DeleteUser(userID int) error
	UpdateUser(user entities.User) error
}
