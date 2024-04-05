package album

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type CreateAlbumUsecase struct {
	re repository.IAlbumRepository
	ha handler.IAlbumHandler
}

func NewCreateAlbumUsecase(re repository.IAlbumRepository, ha handler.IAlbumHandler) *CreateAlbumUsecase {
	return &CreateAlbumUsecase{re, ha}
}

func (u *CreateAlbumUsecase) CreateAlbum(album domain.Album) (domain.AlbumResponse, error) {
	if err := u.ha.AlbumHandler(album); err != nil {
		return domain.AlbumResponse{}, err
	}
	if err := u.re.Create(&album); err != nil {
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
