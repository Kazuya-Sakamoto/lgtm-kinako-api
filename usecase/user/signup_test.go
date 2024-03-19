package user

import (
	"testing"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func setupSignupUsecase(t *testing.T) (*mock.MockUserRepository, *mock.MockUserHandler, *SignupUsecase, func()) {
	mr := new(mock.MockUserRepository)
	mh := new(mock.MockUserHandler)
	usecase := NewSignupUsecase(mr, mh)

	return mr, mh, usecase, func() {
		mr.AssertExpectations(t)
		mh.AssertExpectations(t)
	}
}

func Test_UserUsecase_SignUp(t *testing.T) {
	t.Run("正常にSignupが成功すること", func(t *testing.T) {
		mr, mh, usecase, cleanup := setupSignupUsecase(t)
		defer cleanup()

		input := domain.User{
			Email:    "test@test.com",
			Password: "testpassword",
		}

		mh.On("UserHandler", input).Return(nil)
		mr.On("Create", testifyMock.AnythingOfType("*domain.User")).Run(func(args testifyMock.Arguments) {
			arg := args.Get(0).(*domain.User)
			arg.ID = 1
		}).Return(nil)
		res, err := usecase.SignUp(input)

		require.NoError(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, input.Email, res.Email)
	})

	t.Run("Handlerがエラーを返した場合にSignupが失敗すること", func(t *testing.T) {
		_, mockUserHandler, mh, cleanup := setupSignupUsecase(t)
		defer cleanup()

		input := domain.User{
			Email:    "test@test.com",
			Password: "testpassword",
		}

		mockUserHandler.On("UserHandler", input).Return(assert.AnError)
		_, err := mh.SignUp(input)

		require.Error(t, err)
	})

	t.Run("Createがエラーを返した場合にSignupが失敗すること", func(t *testing.T) {
		mr, mh, usecase, cleanup := setupSignupUsecase(t)
		defer cleanup()

		input := domain.User{
			Email:    "test@test.com",
			Password: "testpassword",
		}

		mh.On("UserHandler", input).Return(nil)
		mr.On("Create", testifyMock.AnythingOfType("*domain.User")).Return(assert.AnError)
		_, err := usecase.SignUp(input)

		require.Error(t, err)
	})
}
