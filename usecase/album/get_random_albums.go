package album

import (
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/model"
	"lgtm-kinako-api/repository"
)

type GetRandomAlbumsUsecase struct {
	ar repository.IAlbumRepository
	ah handler.IAlbumHandler
}

func NewGetRandomAlbumsUsecase(ar repository.IAlbumRepository, ah handler.IAlbumHandler) *GetRandomAlbumsUsecase {
	return &GetRandomAlbumsUsecase{ar, ah}
}

func (au *GetRandomAlbumsUsecase) GetRandomAlbums() ([]model.AlbumResponse, error) {
	albums := []model.Album{}
	if err := au.ar.GetRandomAlbums(&albums); err != nil {
		return nil, err
	}
	res := []model.AlbumResponse{}
	for _, v := range albums {
		a := model.AlbumResponse{
			ID:          v.ID,
			Title:       v.Title,
			Image: 		   v.Image,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		res = append(res, a)
	}
	return res, nil
}