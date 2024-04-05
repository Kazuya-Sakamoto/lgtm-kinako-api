package album

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository"
)

type GetRandomAlbumsUsecase struct {
	re repository.IAlbumRepository
}

func NewGetRandomAlbumsUsecase(re repository.IAlbumRepository) *GetRandomAlbumsUsecase {
	return &GetRandomAlbumsUsecase{re}
}

func (u *GetRandomAlbumsUsecase) GetRandomAlbums() ([]domain.AlbumResponse, error) {
	albums := []domain.Album{}
	if err := u.re.FindRandom(&albums); err != nil {
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
