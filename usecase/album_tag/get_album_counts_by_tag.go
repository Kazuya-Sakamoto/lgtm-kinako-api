package album_tag

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository"
)

type GetAlbumCountsByTagUsecase struct {
	re repository.IAlbumTagRepository
}

func NewGetAlbumCountsByTagUsecase(re repository.IAlbumTagRepository) *GetAlbumCountsByTagUsecase {
	return &GetAlbumCountsByTagUsecase{re}
}

func (u *GetAlbumCountsByTagUsecase) GetAlbumCountsByTag() ([]domain.TagCount, error) {
	res, err := u.re.FindCountsByTag()
	if err != nil {
		return nil, err
	}
	return res, nil
}
