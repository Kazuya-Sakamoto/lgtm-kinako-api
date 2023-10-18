package user

import "lgtm-kinako-api/domain"

type IUserUsecase interface {
    SignUp(user domain.User) (domain.UserResponse, error)
    Login(user domain.User) (string, error)
}
