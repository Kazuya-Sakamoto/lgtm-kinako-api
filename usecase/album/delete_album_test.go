package album

import (
	"testing"

	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/require"
)

func setupDeleteAlbumUsecase(t *testing.T) (*mock.MockAlbumRepository, *DeleteAlbumUsecase, func()) {
	mr := new(mock.MockAlbumRepository)
	usecase := NewDeleteAlbumUsecase(mr)

	return mr, usecase, func() {
		mr.AssertExpectations(t)
	}
}

func Test_AlbumUsecase_DeleteAlbum(t *testing.T) {
	t.Run("アルバムが正常に削除されること", func(t *testing.T) {
		mr, usecase, cleanup := setupDeleteAlbumUsecase(t)
		defer cleanup()

		userId, albumId := uint(1), uint(1)

		mr.On("DeleteAlbum", albumId).Return(nil)

		err := usecase.DeleteAlbum(userId, albumId)

		require.NoError(t, err)
	})
}
