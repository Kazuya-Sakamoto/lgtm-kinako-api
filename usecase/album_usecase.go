package usecase

import (
	"lgtm-kinako-api/model"
	"lgtm-kinako-api/repository"
)

type IAlbumUsecase interface {
	GetAllAlbums() ([]model.AlbumResponse, error)
	CreateAlbum(task model.Album) (model.AlbumResponse, error)
}

type albumUsecase struct {
	ar repository.IAlbumRepository
}

func NewAlbumUsecase(ar repository.IAlbumRepository) IAlbumUsecase {
	return &albumUsecase{ar}
}

func (au *albumUsecase) GetAllAlbums() ([]model.AlbumResponse, error) {
	albums := []model.Album{}
	if err := au.ar.GetAllAlbums(&albums); err != nil {
		return nil, err
	}
	res := []model.AlbumResponse{}
	for _, v := range albums {
		a := model.AlbumResponse{
			ID:          v.ID,
			Title:       v.Title,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		res = append(res, a)
	}
	return res, nil
}

func (au *albumUsecase) CreateAlbum(album model.Album) (model.AlbumResponse, error) {
	if err := au.ar.CreateAlbum(&album); err != nil {
		return model.AlbumResponse{}, err
	}
	res := model.AlbumResponse{
		ID:          album.ID,
		Title:       album.Title,
		CreatedAt:   album.CreatedAt,
		UpdatedAt:   album.UpdatedAt,
	}

	return res, nil
}
