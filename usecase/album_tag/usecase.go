package album_tag

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository"
)

type AlbumTagUsecase struct {
	DeleteAlbumTagsByTagIDUsecase   *DeleteAlbumTagsByTagIDUsecase
	DeleteAndInsertAlbumTagsUsecase *DeleteAndInsertAlbumTagsUsecase
	GetAlbumCountsByTagUsecase      *GetAlbumCountsByTagUsecase
}

func NewAlbumTagUsecase(atr repository.IAlbumTagRepository) *AlbumTagUsecase {
	return &AlbumTagUsecase{
		DeleteAlbumTagsByTagIDUsecase:   NewDeleteAlbumTagsByTagIDUsecase(atr),
		DeleteAndInsertAlbumTagsUsecase: NewDeleteAndInsertAlbumTagsUsecase(atr),
		GetAlbumCountsByTagUsecase:      NewGetAlbumCountsByTagUsecase(atr),
	}
}

func (atu *AlbumTagUsecase) DeleteAlbumTagsByTagID(tagId uint) error {
	return atu.DeleteAlbumTagsByTagIDUsecase.DeleteAlbumTagsByTagID(tagId)
}

func (atu *AlbumTagUsecase) DeleteAndInsertAlbumTags(albumId uint, tagIds []uint) error {
	return atu.DeleteAndInsertAlbumTagsUsecase.DeleteAndInsertAlbumTags(albumId, tagIds)
}

func (atu *AlbumTagUsecase) GetAlbumCountsByTag() ([]domain.TagCount, error) {
	return atu.GetAlbumCountsByTagUsecase.GetAlbumCountsByTag()
}
