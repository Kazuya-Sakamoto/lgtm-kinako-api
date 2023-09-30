package user

import (
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/model"
	"lgtm-kinako-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type SignupUsecase struct {
	ur repository.IUserRepository
	uh handler.IUserHandler
}

func NewSignupUsecase(ur repository.IUserRepository, uh handler.IUserHandler) *SignupUsecase {
	return &SignupUsecase{ur, uh}
}

func (su *SignupUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := su.uh.UserHandler(user); err != nil {
		return model.UserResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := su.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}

	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}