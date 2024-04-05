package album_tag

import (
	"testing"

	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/require"
)

func setupDeleteAlbumTagsByTagIDUsecase(t *testing.T) (*mock.MockAlbumTagRepository, *AlbumTagUsecase, func()) {
	re := new(mock.MockAlbumTagRepository)
	usecase := NewAlbumTagUsecase(re)

	return re, usecase, func() {
		re.AssertExpectations(t)
	}
}

func Test_AlbumTagUsecase_DeleteAlbumTagsByTagID(t *testing.T) {
	t.Run("アルバムタグが正常に削除されること", func(t *testing.T) {
		re, usecase, cleanup := setupDeleteAlbumTagsByTagIDUsecase(t)
		defer cleanup()

		tagId := uint(1)

		re.On("DeleteByTagID", tagId).Return(nil)

		err := usecase.DeleteAlbumTagsByTagID(tagId)

		require.NoError(t, err)
	})
}
