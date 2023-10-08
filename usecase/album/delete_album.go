// usecase/album/delete_album.go
package album

import "lgtm-kinako-api/repository"

type DeleteAlbumUsecase struct {
	ar repository.IAlbumRepository
}

func NewDeleteAlbumUsecase(ar repository.IAlbumRepository) *DeleteAlbumUsecase {
	return &DeleteAlbumUsecase{ar}
}

func (au *DeleteAlbumUsecase) DeleteAlbum(userId uint, albumId uint) error {
	if err := au.ar.DeleteAlbum(userId, albumId); err != nil {
		return err
	}
	return nil
}
