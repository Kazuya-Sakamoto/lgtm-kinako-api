package handler

import (
	"lgtm-kinako-api/model"

	handler "github.com/go-ozzo/ozzo-validation/v4"
)

type IAlbumHandler interface {
	AlbumHandler(album model.Album) error
}

type albumHandler struct{}

func NewAlbumHandler() IAlbumHandler {
	return &albumHandler{}
}


func (ah *albumHandler) AlbumHandler(album model.Album) error {
	return handler.ValidateStruct(&album,
		handler.Field(
			&album.Title,
			handler.Required.Error(ErrTitleRequired.Error()),
			handler.RuneLength(1, 10).Error(ErrTitleLength.Error()),
		),
	)
}
