package album_tag

import (
	"lgtm-kinako-api/repository"
)

type DeleteAndInsertAlbumTagsUsecase struct {
	atr repository.IAlbumTagRepository
}

func NewDeleteAndInsertAlbumTagsUsecase(atr repository.IAlbumTagRepository) *DeleteAndInsertAlbumTagsUsecase {
	return &DeleteAndInsertAlbumTagsUsecase{atr}
}

func (atu *DeleteAndInsertAlbumTagsUsecase) DeleteAndInsertAlbumTags(albumId uint, tagIds []uint) error {
	if err := atu.atr.DeleteAndInsert(albumId, tagIds); err != nil {
		return err
	}
	return nil
}
