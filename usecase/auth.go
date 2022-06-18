package usecase

import "alterra/test/repository"

type AuthUseCase struct {
	AuthRepository repository.AuthRepositoryInterface
}

func NewAuthUseCase(authRepo repository.AuthRepositoryInterface) AuthUseCaseInterface {
	return &AuthUseCase{
		AuthRepository: authRepo,
	}
}

func (authUseCase *AuthUseCase) Login(email string, password string) (string, error) {
	token, err := authUseCase.AuthRepository.Login(email, password)
	return token, err
}
