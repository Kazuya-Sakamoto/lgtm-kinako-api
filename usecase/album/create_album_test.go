package album

import (
	"testing"
	"time"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/require"
)

func setupCreateAlbumUsecase(t *testing.T) (*mock.MockAlbumRepository, *mock.MockAlbumHandler, *CreateAlbumUsecase, func()) {
	re := new(mock.MockAlbumRepository)
	ha := new(mock.MockAlbumHandler)
	usecase := NewCreateAlbumUsecase(re, ha)

	return re, ha, usecase, func() {
		re.AssertExpectations(t)
		ha.AssertExpectations(t)
	}
}

func Test_AlbumUsecase_CreateAlbum(t *testing.T) {
	t.Run("アルバムが正常に作成されること", func(t *testing.T) {
		re, ha, usecase, cleanup := setupCreateAlbumUsecase(t)
		defer cleanup()

		album := domain.Album{
			ID:        1,
			Title:     "Test Album",
			Image:     "test.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		ha.On("AlbumHandler", album).Return(nil)
		re.On("Create", &album).Return(nil)

		res, err := usecase.CreateAlbum(album)

		require.NoError(t, err)
		require.Equal(t, album.ID, res.ID)
		require.Equal(t, album.Title, res.Title)
		require.Equal(t, album.Image, res.Image)
	})
}
