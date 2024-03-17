package album_tag

import (
	"lgtm-kinako-api/repository"
)

type AlbumTagUsecase struct {
	DeleteAlbumTagsByTagIdUsecase *DeleteAlbumTagsByTagIdUsecase
	ResetAndSetAlbumTagsUsecase   *ResetAndSetAlbumTagsUsecase
}

func NewAlbumTagUsecase(atr repository.IAlbumtagRepository) *AlbumTagUsecase {
	return &AlbumTagUsecase{
		DeleteAlbumTagsByTagIdUsecase: NewDeleteAlbumTagsByTagIdUsecase(atr),
		ResetAndSetAlbumTagsUsecase:   NewResetAndSetAlbumTagsUsecase(atr),
	}
}

func (atu *AlbumTagUsecase) DeleteAlbumTagsByTagId(tagId uint) error {
	return atu.DeleteAlbumTagsByTagIdUsecase.DeleteAlbumTagsByTagId(tagId)
}

func (atu *AlbumTagUsecase) ResetAndSetAlbumTags(albumId uint, tagIds []uint) error {
	return atu.ResetAndSetAlbumTagsUsecase.ResetAndSetAlbumTags(albumId, tagIds)
}
