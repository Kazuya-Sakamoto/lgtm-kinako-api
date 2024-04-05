package album

import "lgtm-kinako-api/repository"

type DeleteAlbumUsecase struct {
	re repository.IAlbumRepository
}

func NewDeleteAlbumUsecase(re repository.IAlbumRepository) *DeleteAlbumUsecase {
	return &DeleteAlbumUsecase{re}
}

func (u *DeleteAlbumUsecase) DeleteAlbum(userId, albumId uint) error {
	if err := u.re.DeleteByAlbumID(userId, albumId); err != nil {
		return err
	}
	return nil
}
