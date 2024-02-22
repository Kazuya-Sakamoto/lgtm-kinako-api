package album

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type GetAllAlbumsUsecase struct {
	ar repository.IAlbumRepository
	ah handler.IAlbumHandler
}

func NewGetAllAlbumsUsecase(ar repository.IAlbumRepository, ah handler.IAlbumHandler) *GetAllAlbumsUsecase {
	return &GetAllAlbumsUsecase{ar, ah}
}

func (au *GetAllAlbumsUsecase) GetAllAlbums(userId uint) ([]domain.AlbumResponse, error) {
	albums := []domain.Album{}
	if err := au.ar.GetAllAlbums(&albums, userId); err != nil {
		return nil, err
	}
	res := []domain.AlbumResponse{}
	for _, v := range albums {
		a := domain.AlbumResponse{
			ID:          v.ID,
			Title:       v.Title,
			Image: 		 v.Image,
			Tags: 	  	 v.Tags,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		res = append(res, a)
	}
	return res, nil
}
