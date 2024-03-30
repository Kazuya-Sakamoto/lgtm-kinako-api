package album_tag

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository"
)

type GetAlbumCountsByTagUsecase struct {
	atr repository.IAlbumTagRepository
}

func NewGetAlbumCountsByTagUsecase(atr repository.IAlbumTagRepository) *GetAlbumCountsByTagUsecase {
	return &GetAlbumCountsByTagUsecase{atr}
}

func (atu *GetAlbumCountsByTagUsecase) GetAlbumCountsByTag() ([]domain.TagCount, error) {
	res, err := atu.atr.FindCountsByTag()
	if err != nil {
		return nil, err
	}
	return res, nil
}
