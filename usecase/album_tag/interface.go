package album_tag

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository"
)

type IAlbumTagUsecase interface {
	DeleteAlbumTagsByTagID(tagId uint) error
	DeleteAndInsertAlbumTags(albumId uint, tagIds []uint) error
	GetAlbumCountsByTag() ([]domain.TagCount, error)
}

type AlbumTagUsecase struct {
	*DeleteAlbumTagsByTagIDUsecase
	*DeleteAndInsertAlbumTagsUsecase
	*GetAlbumCountsByTagUsecase
}

func NewAlbumTagUsecase(re repository.IAlbumTagRepository) *AlbumTagUsecase {
	return &AlbumTagUsecase{
		NewDeleteAlbumTagsByTagIDUsecase(re),
		NewDeleteAndInsertAlbumTagsUsecase(re),
		NewGetAlbumCountsByTagUsecase(re),
	}
}
