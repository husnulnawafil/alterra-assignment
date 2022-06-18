package usecase

import (
	"alterra/test/entities"
	"alterra/test/repository"
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

func (userUseCase *UserUseCase) GetListUsers() ([]entities.User, error) {
	listUser, err := userUseCase.UserRepository.GetListUsers()
	return listUser, err
}

func (userUseCase *UserUseCase) DeleteUser(userID int) error {
	err := userUseCase.UserRepository.DeleteUser(userID)
	return err
}

func (userUseCase *UserUseCase) UpdateUser(user entities.User, userID int) error {
	err := userUseCase.UserRepository.UpdateUser(user)
	return err
}
