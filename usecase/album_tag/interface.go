package album_tag

type IAlbumTagUsecase interface {
	DeleteAlbumTagsByTagId(tagId uint) error
	ResetAndSetAlbumTags(albumId uint, tagIds []uint) error
}
