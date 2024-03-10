package album_tag

type IAlbumTagUsecase interface {
	DeleteAlbumTagsByTagId(tagId uint) error
}
