package handler

import (
	"lgtm-kinako-api/model"

	handler "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)


type IUserHandler interface {
	UserHandler(user model.User) error
}

type userHandler struct{}

func NewUserHandler() IUserHandler {
	return &userHandler{}
}

func (uv *userHandler) UserHandler(user model.User) error {
	return handler.ValidateStruct(&user,
		handler.Field(
			&user.Email,
			handler.Required.Error(ErrEmailRequired.Error()),
			handler.RuneLength(1, 30).Error(ErrEmailLength.Error()),
			is.Email.Error(ErrInvalidEmailFormat.Error()),
		),
		handler.Field(
			&user.Password,
			handler.Required.Error(ErrPasswordRequired.Error()),
			handler.RuneLength(6, 30).Error(ErrPasswordLength.Error()),
		),
	)
}