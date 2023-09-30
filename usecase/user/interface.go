package user

import (
	"lgtm-kinako-api/model"
)

type IUserUsecase interface {
    SignUp(user model.User) (model.UserResponse, error)
    Login(user model.User) (string, error)
}
