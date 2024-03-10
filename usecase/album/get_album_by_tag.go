package album

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type GetAlbumsByTagUsecase struct {
	ar repository.IAlbumRepository
	ah handler.IAlbumHandler
}

func NewGetAlbumsByTagUsecase(ar repository.IAlbumRepository, ah handler.IAlbumHandler) *GetAlbumsByTagUsecase {
	return &GetAlbumsByTagUsecase{ar, ah}
}

func (au *GetAlbumsByTagUsecase) GetAlbumsByTag(tagId uint) ([]domain.AlbumResponse, error) {
	albums := []domain.Album{}
	if err := au.ar.GetAlbumsByTag(&albums, tagId); err != nil {
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
