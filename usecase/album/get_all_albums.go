package album

import (
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/model"
	"lgtm-kinako-api/repository"
)

type GetAllAlbumsUsecase struct {
	ar repository.IAlbumRepository
	ah handler.IAlbumHandler
}

func NewGetAllAlbumsUsecase(ar repository.IAlbumRepository, ah handler.IAlbumHandler) *GetAllAlbumsUsecase {
	return &GetAllAlbumsUsecase{ar, ah}
}

func (au *GetAllAlbumsUsecase) GetAllAlbums() ([]model.AlbumResponse, error) {
	albums := []model.Album{}
	if err := au.ar.GetAllAlbums(&albums); err != nil {
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