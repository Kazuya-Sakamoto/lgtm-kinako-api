// usecase/album/create_album.go
package album

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type CreateAlbumUsecase struct {
	ar repository.IAlbumRepository
	ah handler.IAlbumHandler
}

func NewCreateAlbumUsecase(ar repository.IAlbumRepository, ah handler.IAlbumHandler) *CreateAlbumUsecase {
	return &CreateAlbumUsecase{ar, ah}
}

func (au *CreateAlbumUsecase) CreateAlbum(album domain.Album) (domain.AlbumResponse, error) {
	if err := au.ah.AlbumHandler(album); err != nil {
		return domain.AlbumResponse{}, err
	}
	if err := au.ar.CreateAlbum(&album); err != nil {
		return domain.AlbumResponse{}, err
	}
	res := domain.AlbumResponse{
		ID:        album.ID,
		Title:     album.Title,
		Image:     album.Image,
		CreatedAt: album.CreatedAt,
		UpdatedAt: album.UpdatedAt,
	}

	return res, nil
}
