package album_tag

import (
	"lgtm-kinako-api/repository"
)

type ResetAndSetAlbumTagsUsecase struct {
	atr repository.IAlbumTagRepository
}

func NewResetAndSetAlbumTagsUsecase(atr repository.IAlbumTagRepository) *ResetAndSetAlbumTagsUsecase {
	return &ResetAndSetAlbumTagsUsecase{atr}
}

func (atu *ResetAndSetAlbumTagsUsecase) ResetAndSetAlbumTags(albumId uint, tagIds []uint) error {
	if err := atu.atr.ResetAndSet(albumId, tagIds); err != nil {
		return err
	}
	return nil
}
