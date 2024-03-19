package album

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type GetRandomAlbumsUsecase struct {
	ar repository.IAlbumRepository
	ah handler.IAlbumHandler
}

func NewGetRandomAlbumsUsecase(ar repository.IAlbumRepository, ah handler.IAlbumHandler) *GetRandomAlbumsUsecase {
	return &GetRandomAlbumsUsecase{ar, ah}
}

func (au *GetRandomAlbumsUsecase) GetRandomAlbums() ([]domain.AlbumResponse, error) {
	albums := []domain.Album{}
	if err := au.ar.FindRandom(&albums); err != nil {
		return nil, err
	}
	res := []domain.AlbumResponse{}
	for _, v := range albums {
		a := domain.AlbumResponse{
			ID:        v.ID,
			Title:     v.Title,
			Image:     v.Image,
			Tags:      v.Tags,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		res = append(res, a)
	}
	return res, nil
}
