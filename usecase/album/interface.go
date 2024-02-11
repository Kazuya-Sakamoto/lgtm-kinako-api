package album

import "lgtm-kinako-api/domain"

type IAlbumUsecase interface {
	GetAllAlbums(userId uint) ([]domain.AlbumResponse, error)
	GetRandomAlbums() ([]domain.AlbumResponse, error)
	GetAlbumsByTag(tagId uint) ([]domain.AlbumResponse, error)
	CreateAlbum(album domain.Album) (domain.AlbumResponse, error)
	UploadImageToS3(encodedImage []byte) (string, error)
	DeleteAlbum(userId uint, albumId uint) error
}
