package album_tag

import (
	"lgtm-kinako-api/repository"
)

type DeleteAlbumTagsByTagIDUsecase struct {
	atr repository.IAlbumTagRepository
}

func NewDeleteAlbumTagsByTagIDUsecase(atr repository.IAlbumTagRepository) *DeleteAlbumTagsByTagIDUsecase {
	return &DeleteAlbumTagsByTagIDUsecase{atr}
}

func (atu *DeleteAlbumTagsByTagIDUsecase) DeleteAlbumTagsByTagID(tagId uint) error {
	if err := atu.atr.DeleteByTagID(tagId); err != nil {
		return err
	}
	return nil
}
