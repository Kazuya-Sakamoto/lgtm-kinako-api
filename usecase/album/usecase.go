package album

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type AlbumUsecase struct {
	GetAllAlbumsUsecase 	*GetAllAlbumsUsecase
	GetRandomAlbumsUsecase  *GetRandomAlbumsUsecase
	UploadImageToS3Usecase  *UploadImageToS3Usecase
	CreateAlbumUsecase	    *CreateAlbumUsecase
	DeleteAlbumUsecase	    *DeleteAlbumUsecase
}

func NewAlbumUsecase(ar repository.IAlbumRepository, uh handler.IAlbumHandler) *AlbumUsecase {
	return &AlbumUsecase{
		GetAllAlbumsUsecase: 	NewGetAllAlbumsUsecase(ar, uh),
		GetRandomAlbumsUsecase: NewGetRandomAlbumsUsecase(ar, uh),
		UploadImageToS3Usecase: NewUploadImageToS3Usecase(ar, uh),
		CreateAlbumUsecase:  	NewCreateAlbumUsecase(ar, uh),
		DeleteAlbumUsecase: 	NewDeleteAlbumUsecase(ar),
	}
}

func (au *AlbumUsecase) GetAllAlbums(userId uint) ([]domain.AlbumResponse, error) {
	return au.GetAllAlbumsUsecase.GetAllAlbums(userId)
}

func (au *AlbumUsecase) GetRandomAlbums() ([]domain.AlbumResponse, error) {
	return au.GetRandomAlbumsUsecase.GetRandomAlbums()
}

func (au *AlbumUsecase) UploadImageToS3(encodedImage []byte) (string, error) {
	return au.UploadImageToS3Usecase.UploadImageToS3(encodedImage)
}

func (au *AlbumUsecase) CreateAlbum(album domain.Album) (domain.AlbumResponse, error) {
	return au.CreateAlbumUsecase.CreateAlbum(album)
}

func (au *AlbumUsecase) DeleteAlbum(userId uint, albumId uint) error {
	return au.DeleteAlbumUsecase.DeleteAlbum(userId, albumId)
}