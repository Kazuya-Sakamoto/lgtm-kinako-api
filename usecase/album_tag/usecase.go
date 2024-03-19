package album_tag

import (
	"lgtm-kinako-api/repository"
)

type AlbumTagUsecase struct {
	DeleteAlbumTagsByTagIDUsecase *DeleteAlbumTagsByTagIDUsecase
	ResetAndSetAlbumTagsUsecase   *ResetAndSetAlbumTagsUsecase
}

func NewAlbumTagUsecase(atr repository.IAlbumTagRepository) *AlbumTagUsecase {
	return &AlbumTagUsecase{
		DeleteAlbumTagsByTagIDUsecase: NewDeleteAlbumTagsByTagIDUsecase(atr),
		ResetAndSetAlbumTagsUsecase:   NewResetAndSetAlbumTagsUsecase(atr),
	}
}

func (atu *AlbumTagUsecase) DeleteAlbumTagsByTagID(tagId uint) error {
	return atu.DeleteAlbumTagsByTagIDUsecase.DeleteAlbumTagsByTagID(tagId)
}

func (atu *AlbumTagUsecase) ResetAndSetAlbumTags(albumId uint, tagIds []uint) error {
	return atu.ResetAndSetAlbumTagsUsecase.ResetAndSetAlbumTags(albumId, tagIds)
}
