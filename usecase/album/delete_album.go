package album

import "lgtm-kinako-api/repository"

type DeleteAlbumUsecase struct {
	ar repository.IAlbumRepository
}

func NewDeleteAlbumUsecase(ar repository.IAlbumRepository) *DeleteAlbumUsecase {
	return &DeleteAlbumUsecase{ar}
}

func (au *DeleteAlbumUsecase) DeleteAlbum(userId, albumId uint) error {
	if err := au.ar.DeleteByAlbumID(userId, albumId); err != nil {
		return err
	}
	return nil
}
