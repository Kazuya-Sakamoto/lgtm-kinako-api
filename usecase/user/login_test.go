package user

import (
	"errors"
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository/mock"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupLoginUsecase(t *testing.T) (*mock.MockUserRepository, *mock.MockUserHandler, *LoginUsecase, func()) {
	originalSecret := os.Getenv("SECRET")
	os.Setenv("SECRET", "testsecret")

	mockRepository := new(mock.MockUserRepository)
	mockHandler := new(mock.MockUserHandler)
	loginUsecase := NewLoginUsecase(mockRepository, mockHandler)

	return mockRepository, mockHandler, loginUsecase, func() {
		os.Setenv("SECRET", originalSecret)
		mockRepository.AssertExpectations(t)
		mockHandler.AssertExpectations(t)
	}
}

func TestLoginUsecase_Login(t *testing.T) {
	t.Run("正常にLoginが成功すること", func(t *testing.T) {
		mockRepository, mockHandler, loginUsecase, cleanup := setupLoginUsecase(t)
		defer cleanup()

		input := domain.User{
			Email:    "test@test.com",
			Password: "testpassword",
		}
		user := domain.User{
			ID:       1,
			Email:    "test@test.com",
			Password: "$2a$10$98.1xBaze.nkLmN6wRCYGe8/j3kVsOGLICaEHK6zs37AQNCkW8uQq",
		}
		/* 
			モックの期待値の設定
		*/
		mockHandler.On("UserHandler", input).Return(nil)
		mockRepository.On("GetUserByEmail", &domain.User{}, input.Email).Return(user, nil)
		/* 
			ログインの実行
		*/
		token, err := loginUsecase.Login(input)
		require.NoError(t, err)
		require.NotEmpty(t, token)
		/* 
			トークンの解析
		*/
		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})
		require.NoError(t, err)
		require.True(t, parsedToken.Valid)
		/* 
			トークンのクレームの検証
		*/
		now := time.Now().Unix()
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			exp := int64(claims["exp"].(float64))
			assert.Equal(t, float64(user.ID), claims["user_id"])
			require.True(t, exp > now)
		} else {
			t.Fail()
		}
		// モックの検証
		mockHandler.AssertExpectations(t)
		mockRepository.AssertExpectations(t)
	})

	t.Run("ユーザーが存在しない場合Loginが失敗すること", func(t *testing.T) {
		mockRepository, mockHandler, loginUsecase, cleanup := setupLoginUsecase(t)
		defer cleanup()

		input := domain.User{
			Email:    "xxx@test.com",
			Password: "xxxpassword",
		}
		/* 
			ユーザーが見つからない場合のモックの設定
		*/
		mockHandler.On("UserHandler", input).Return(nil)
		mockRepository.On("GetUserByEmail", &domain.User{}, input.Email).Return(domain.User{}, errors.New("user not found"))
		/* 
			ログインの実行
		*/
		_, err := loginUsecase.Login(input)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
		/* 
			モックの検証
		*/
		mockHandler.AssertExpectations(t)
		mockRepository.AssertExpectations(t)
	})
}
