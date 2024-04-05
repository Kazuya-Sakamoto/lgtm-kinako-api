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
	re := new(mock.MockUserRepository)
	ha := new(mock.MockUserHandler)
	usecase := NewSignupUsecase(re, ha)

	return re, ha, usecase, func() {
		re.AssertExpectations(t)
		ha.AssertExpectations(t)
	}
}

func Test_UserUsecase_SignUp(t *testing.T) {
	t.Run("正常にSignupが成功すること", func(t *testing.T) {
		re, ha, usecase, cleanup := setupSignupUsecase(t)
		defer cleanup()

		input := domain.User{
			Email:    "test@test.com",
			Password: "testpassword",
		}

		ha.On("UserHandler", input).Return(nil)
		re.On("Create", testifyMock.AnythingOfType("*domain.User")).Run(func(args testifyMock.Arguments) {
			arg := args.Get(0).(*domain.User)
			arg.ID = 1
		}).Return(nil)
		res, err := usecase.SignUp(input)

		require.NoError(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, input.Email, res.Email)
	})

	t.Run("Handlerがエラーを返した場合にSignupが失敗すること", func(t *testing.T) {
		_, ha, mh, cleanup := setupSignupUsecase(t)
		defer cleanup()

		input := domain.User{
			Email:    "test@test.com",
			Password: "testpassword",
		}

		ha.On("UserHandler", input).Return(assert.AnError)
		_, err := mh.SignUp(input)

		require.Error(t, err)
	})

	t.Run("Createがエラーを返した場合にSignupが失敗すること", func(t *testing.T) {
		re, ha, usecase, cleanup := setupSignupUsecase(t)
		defer cleanup()

		input := domain.User{
			Email:    "test@test.com",
			Password: "testpassword",
		}

		ha.On("UserHandler", input).Return(nil)
		re.On("Create", testifyMock.AnythingOfType("*domain.User")).Return(assert.AnError)
		_, err := usecase.SignUp(input)

		require.Error(t, err)
	})
}
