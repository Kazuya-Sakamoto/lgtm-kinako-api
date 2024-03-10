package album_tag

import (
	"lgtm-kinako-api/repository"
)

type DeleteAlbumTagsByTagIdUsecase struct {
	atr repository.IAlbumtagRepository
}

func NewDeleteAlbumTagsByTagIdUsecase(atr repository.IAlbumtagRepository) *DeleteAlbumTagsByTagIdUsecase {
	return &DeleteAlbumTagsByTagIdUsecase{atr}
}

func (atu *DeleteAlbumTagsByTagIdUsecase) DeleteAlbumTagsByTagId(tagId uint) error {
	if err := atu.atr.DeleteAlbumTagsByTagId(tagId); err != nil {
		return err
	}
	return nil
}
