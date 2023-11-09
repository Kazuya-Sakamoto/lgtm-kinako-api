package user

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func setupSignupUsecase(t *testing.T) (*mock.MockUserRepository, *mock.MockUserHandler, *SignupUsecase, func()) {
	mockRepository := new(mock.MockUserRepository)
	mockHandler := new(mock.MockUserHandler)
	signupUsecase := NewSignupUsecase(mockRepository, mockHandler)

	return mockRepository, mockHandler, signupUsecase, func() {
		mockRepository.AssertExpectations(t)
		mockHandler.AssertExpectations(t)
	}
}

func TestSignupUsecase_SignUp(t *testing.T) {
	t.Run("正常にSignupが成功すること", func(t *testing.T) {
		mockRepository, mockHandler, signupUsecase, cleanup := setupSignupUsecase(t)
		defer cleanup()

		input := domain.User{
			Email:    "test@test.com",
			Password: "testpassword",
		}

		mockHandler.On("UserHandler", input).Return(nil)
		mockRepository.On("CreateUser", testifyMock.AnythingOfType("*domain.User")).Run(func(args testifyMock.Arguments) {
			arg := args.Get(0).(*domain.User)
			arg.ID = 1
		}).Return(nil)
		resUser, err := signupUsecase.SignUp(input)

		require.NoError(t, err)
		assert.NotEmpty(t, resUser.ID)
		assert.Equal(t, input.Email, resUser.Email)
	})

	t.Run("Handlerがエラーを返した場合にSignupが失敗すること", func(t *testing.T) {
		_, mockUserHandler, mockHandler, cleanup := setupSignupUsecase(t)
		defer cleanup()

		input := domain.User{
			Email:    "test@test.com",
			Password: "testpassword",
		}

		mockUserHandler.On("UserHandler", input).Return(assert.AnError)
		_, err := mockHandler.SignUp(input)

		require.Error(t, err)
	})

	t.Run("CreateUserがエラーを返した場合にSignupが失敗すること", func(t *testing.T) {
		mockRepository, mockHandler, signupUsecase, cleanup := setupSignupUsecase(t)
		defer cleanup()

		input := domain.User{
			Email:    "test@test.com",
			Password: "testpassword",
		}

		mockHandler.On("UserHandler", input).Return(nil)
		mockRepository.On("CreateUser", testifyMock.AnythingOfType("*domain.User")).Return(assert.AnError)
		_, err := signupUsecase.SignUp(input)

		require.Error(t, err)
	})
}
