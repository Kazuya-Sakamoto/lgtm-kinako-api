package user

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type SignupUsecase struct {
	re repository.IUserRepository
	ha handler.IUserHandler
}

func NewSignupUsecase(re repository.IUserRepository, ha handler.IUserHandler) *SignupUsecase {
	return &SignupUsecase{re, ha}
}

func (u *SignupUsecase) SignUp(user domain.User) (domain.UserResponse, error) {
	if err := u.ha.UserHandler(user); err != nil {
		return domain.UserResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return domain.UserResponse{}, err
	}
	newUser := domain.User{Email: user.Email, Password: string(hash)}
	if err := u.re.Create(&newUser); err != nil {
		return domain.UserResponse{}, err
	}

	resUser := domain.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}
