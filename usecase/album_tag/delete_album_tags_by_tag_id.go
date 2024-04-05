package album_tag

import (
	"lgtm-kinako-api/repository"
)

type DeleteAlbumTagsByTagIDUsecase struct {
	re repository.IAlbumTagRepository
}

func NewDeleteAlbumTagsByTagIDUsecase(re repository.IAlbumTagRepository) *DeleteAlbumTagsByTagIDUsecase {
	return &DeleteAlbumTagsByTagIDUsecase{re}
}

func (u *DeleteAlbumTagsByTagIDUsecase) DeleteAlbumTagsByTagID(tagId uint) error {
	if err := u.re.DeleteByTagID(tagId); err != nil {
		return err
	}
	return nil
}
