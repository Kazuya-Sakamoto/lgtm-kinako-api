package handler

import (
	"lgtm-kinako-api/domain"

	handler "github.com/go-ozzo/ozzo-validation/v4"
)

type ITagHandler interface {
	TagHandler(tag domain.Tag) error
}

type tagHandler struct{}

func NewTagHandler() ITagHandler {
	return &tagHandler{}
}

func (th *tagHandler) TagHandler(tag domain.Tag) error {
	return handler.ValidateStruct(&tag,
		handler.Field(
			&tag.Name,
			handler.Required.Error(ErrTagRequired.Error()),
			handler.RuneLength(1, 10).Error(ErrTagLength.Error()),
		),
	)
}
