package album

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository"
)

type GetAlbumsByTagUsecase struct {
	re repository.IAlbumRepository
}

func NewGetAlbumsByTagUsecase(re repository.IAlbumRepository) *GetAlbumsByTagUsecase {
	return &GetAlbumsByTagUsecase{re}
}

func (u *GetAlbumsByTagUsecase) GetAlbumsByTag(tagId uint) ([]domain.AlbumResponse, error) {
	albums := []domain.Album{}
	if err := u.re.FindByTag(&albums, tagId); err != nil {
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
