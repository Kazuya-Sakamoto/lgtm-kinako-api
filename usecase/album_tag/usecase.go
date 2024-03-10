package album_tag

import (
	"lgtm-kinako-api/repository"
)

type AlbumTagUsecase struct {
	DeleteAlbumTagsByTagIdUsecase *DeleteAlbumTagsByTagIdUsecase
}

func NewAlbumTagUsecase(atr repository.IAlbumtagRepository) *AlbumTagUsecase {
	return &AlbumTagUsecase{
		DeleteAlbumTagsByTagIdUsecase: NewDeleteAlbumTagsByTagIdUsecase(atr),
	}
}

func (atu *AlbumTagUsecase) DeleteAlbumTagsByTagId(tagId uint) error {
	return atu.DeleteAlbumTagsByTagIdUsecase.DeleteAlbumTagsByTagId(tagId)
}
