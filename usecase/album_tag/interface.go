package album_tag

import "lgtm-kinako-api/domain"

type IAlbumTagUsecase interface {
	DeleteAlbumTagsByTagID(tagId uint) error
	DeleteAndInsertAlbumTags(albumId uint, tagIds []uint) error
	GetAlbumCountsByTag() ([]domain.TagCount, error)
}
