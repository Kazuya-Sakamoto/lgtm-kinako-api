package album

import (
	"lgtm-kinako-api/model"
)

type IAlbumUsecase interface {
	GetAllAlbums() ([]model.AlbumResponse, error)
	GetRandomAlbums() ([]model.AlbumResponse, error)
	CreateAlbum(album model.Album) (model.AlbumResponse, error)
	UploadImageToS3(encodedImage []byte) (string, error)
	DeleteAlbum(userId uint, albumId uint) error
}
