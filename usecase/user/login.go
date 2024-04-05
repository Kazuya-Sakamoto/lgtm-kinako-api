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
	re repository.IUserRepository
	ha handler.IUserHandler
}

func NewLoginUsecase(re repository.IUserRepository, ha handler.IUserHandler) *LoginUsecase {
	return &LoginUsecase{re, ha}
}

func (u *LoginUsecase) Login(user domain.User) (string, error) {
	if err := u.ha.UserHandler(user); err != nil {
		return "", err
	}
	storedUser := domain.User{}
	if err := u.re.FindByEmail(&storedUser, user.Email); err != nil {
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
