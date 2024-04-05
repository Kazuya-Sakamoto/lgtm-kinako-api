package user

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type IUserUsecase interface {
	SignUp(user domain.User) (domain.UserResponse, error)
	Login(user domain.User) (string, error)
}

type UserUsecase struct {
	*SignupUsecase
	*LoginUsecase
}

func NewUserUsecase(
	re repository.IUserRepository,
	ha handler.IUserHandler,
) *UserUsecase {
	return &UserUsecase{
		NewSignupUsecase(re, ha),
		NewLoginUsecase(re, ha),
	}
}
