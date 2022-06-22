package usecase

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Run("TestLoginSuccess", func(t *testing.T) {
		authUseCase := NewAuthUseCase(mockAuthRepository{})
		data, err := authUseCase.Login("husnul@mail.com", "12345")
		assert.Nil(t, err)
		assert.Equal(t, "husnul@mail.com", data)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		authUseCase := NewAuthUseCase(mockAuthRepositoryError{})
		data, err := authUseCase.Login("", "")
		assert.NotNil(t, err)
		assert.Nil(t, nil, data)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("TestDeleteUserSucsess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})

	})
}

// === mock success ===
type mockAuthRepository struct{}

type mockUserRepository struct{}

func (ma mockAuthRepository) Login(email string, password string) (string, error) {
	return "husnul@mail.com", nil
}

func (mu mockUserRepository) DeleteUser(userID int) error {
	return nil
}

// === mock error ===

type mockAuthRepositoryError struct{}

func (ma mockAuthRepositoryError) Login(email string, password string) (string, error) {
	return "", fmt.Errorf("error")
}
