package album_tag

import (
	"testing"

	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/require"
)

func setupDeleteAlbumTagsByTagIDUsecase(t *testing.T) (*mock.MockAlbumTagRepository, *AlbumTagUsecase, func()) {
	mr := new(mock.MockAlbumTagRepository)
	usecase := NewAlbumTagUsecase(mr)

	return mr, usecase, func() {
		mr.AssertExpectations(t)
	}
}

func Test_AlbumTagUsecase_DeleteAlbumTagsByTagID(t *testing.T) {
	t.Run("アルバムタグが正常に削除されること", func(t *testing.T) {
		mr, usecase, cleanup := setupDeleteAlbumTagsByTagIDUsecase(t)
		defer cleanup()

		tagId := uint(1)

		mr.On("DeleteByTagID", tagId).Return(nil)

		err := usecase.DeleteAlbumTagsByTagID(tagId)

		require.NoError(t, err)
	})
}
