package user

import (
	"os"
	"time"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	ur repository.IUserRepository
	uh handler.IUserHandler
}

func NewLoginUsecase(ur repository.IUserRepository, uh handler.IUserHandler) *LoginUsecase {
	return &LoginUsecase{ur, uh}
}

func (lu *LoginUsecase) Login(user domain.User) (string, error) {
	if err := lu.uh.UserHandler(user); err != nil {
		return "", err
	}
	storedUser := domain.User{}
	if err := lu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
