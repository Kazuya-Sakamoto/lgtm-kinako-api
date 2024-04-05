package album

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type IAlbumUsecase interface {
	GetAllAlbums() ([]domain.AlbumResponse, error)
	GetRandomAlbums() ([]domain.AlbumResponse, error)
	GetAlbumsByTag(tagId uint) ([]domain.AlbumResponse, error)
	CreateAlbum(album domain.Album) (domain.AlbumResponse, error)
	UploadImageToS3(encodedImage []byte) (string, error)
	DeleteAlbum(userId uint, albumId uint) error
}

type AlbumUsecase struct {
	*GetAllAlbumsUsecase
	*GetRandomAlbumsUsecase
	*GetAlbumsByTagUsecase
	*UploadImageToS3Usecase
	*CreateAlbumUsecase
	*DeleteAlbumUsecase
}

func NewAlbumUsecase(
	re repository.IAlbumRepository,
	ha handler.IAlbumHandler,
) *AlbumUsecase {
	return &AlbumUsecase{
		NewGetAllAlbumsUsecase(re),
		NewGetRandomAlbumsUsecase(re),
		NewGetAlbumsByTagUsecase(re),
		NewUploadImageToS3Usecase(),
		NewCreateAlbumUsecase(re, ha),
		NewDeleteAlbumUsecase(re),
	}
}
