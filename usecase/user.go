package usecase

import (
	"alterra/test/entities"
	"alterra/test/repository"

	"github.com/jinzhu/copier"
)

type UserUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserUseCase(userRepo repository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		UserRepository: userRepo,
	}
}

func (userUseCase *UserUseCase) CreateUser(newUser entities.User) (int, error) {
	err := userUseCase.UserRepository.CreateUser(newUser)
	return 0, err
}

func (userUseCase *UserUseCase) GetListUsers() ([]entities.GetUserResponse, error) {
	var userResponse []entities.GetUserResponse
	listUser, err := userUseCase.UserRepository.GetListUsers()
	copier.Copy(&userResponse, &listUser)
	return userResponse, err
}

func (userUseCase *UserUseCase) DeleteUser(userID int) error {
	err := userUseCase.UserRepository.DeleteUser(userID)
	return err
}

func (userUseCase *UserUseCase) UpdateUser(user entities.User, userID int) error {
	err := userUseCase.UserRepository.UpdateUser(user)
	return err
}

func (userUseCase *UserUseCase) GetUserById(userID int) (entities.GetUserResponse, error) {
	var userResponse entities.GetUserResponse
	user, err := userUseCase.UserRepository.GetUserById(userID)
	copier.Copy(&userResponse, &user)
	return userResponse, err
}
