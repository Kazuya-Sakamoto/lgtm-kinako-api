package user

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type UserUsecase struct {
	SignUpUsecase *SignupUsecase
	LoginUsecase  *LoginUsecase
}

func NewUserUsecase(ur repository.IUserRepository, uh handler.IUserHandler) *UserUsecase {
	return &UserUsecase{
		SignUpUsecase: NewSignupUsecase(ur, uh),
		LoginUsecase:  NewLoginUsecase(ur, uh),
	}
}

func (u *UserUsecase) SignUp(user domain.User) (domain.UserResponse, error) {
	return u.SignUpUsecase.SignUp(user)
}

func (u *UserUsecase) Login(user domain.User) (string, error) {
	return u.LoginUsecase.Login(user)
}