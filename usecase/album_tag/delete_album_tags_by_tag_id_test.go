package album_tag

import (
	"testing"

	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/require"
)

func setupDeleteAlbumTagsByTagIdUsecase(t *testing.T) (*mock.MockAlbumTagRepository, *AlbumTagUsecase, func()) {
	mr := new(mock.MockAlbumTagRepository)
	usecase := NewAlbumTagUsecase(mr)

	return mr, usecase, func() {
		mr.AssertExpectations(t)
	}
}

func Test_AlbumTagUsecase_DeleteAlbumTagsByTagId(t *testing.T) {
	t.Run("アルバムタグが正常に削除されること", func(t *testing.T) {
		mr, usecase, cleanup := setupDeleteAlbumTagsByTagIdUsecase(t)
		defer cleanup()

		tagId := uint(1)

		mr.On("DeleteAlbumTagsByTagId", tagId).Return(nil)

		err := usecase.DeleteAlbumTagsByTagId(tagId)

		require.NoError(t, err)
	})
}
