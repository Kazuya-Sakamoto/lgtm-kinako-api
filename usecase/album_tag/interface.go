package album_tag

type IAlbumTagUsecase interface {
	DeleteAlbumTagsByTagID(tagId uint) error
	DeleteAndInsertAlbumTags(albumId uint, tagIds []uint) error
}
