package album_tag

type IAlbumTagUsecase interface {
	DeleteAlbumTagsByTagID(tagId uint) error
	ResetAndSetAlbumTags(albumId uint, tagIds []uint) error
}
