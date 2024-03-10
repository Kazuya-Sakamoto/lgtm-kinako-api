package handler

import (
	"lgtm-kinako-api/domain"

	handler "github.com/go-ozzo/ozzo-validation/v4"
)

type IAlbumHandler interface {
	AlbumHandler(album domain.Album) error
}

type albumHandler struct{}

func NewAlbumHandler() IAlbumHandler {
	return &albumHandler{}
}

func (ah *albumHandler) AlbumHandler(album domain.Album) error {
	return handler.ValidateStruct(&album,
		handler.Field(
			&album.Title,
			handler.Required.Error(ErrTitleRequired.Error()),
			handler.RuneLength(1, 10).Error(ErrTitleLength.Error()),
		),
	)
}
