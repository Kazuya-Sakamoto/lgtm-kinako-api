package album

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository"
)

type GetAllAlbumsUsecase struct {
	re repository.IAlbumRepository
}

func NewGetAllAlbumsUsecase(re repository.IAlbumRepository) *GetAllAlbumsUsecase {
	return &GetAllAlbumsUsecase{re}
}

func (u *GetAllAlbumsUsecase) GetAllAlbums() ([]domain.AlbumResponse, error) {
	albums := []domain.Album{}
	if err := u.re.FindAll(&albums); err != nil {
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
