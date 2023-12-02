package album

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func setupCreateAlbumUsecase(t *testing.T) (*mock.MockAlbumRepository, *mock.MockAlbumHandler, *CreateAlbumUsecase, func()) {
	mr := new(mock.MockAlbumRepository)
	mh := new(mock.MockAlbumHandler)
	usecase := NewCreateAlbumUsecase(mr, mh)

	return mr, mh, usecase, func() {
		mr.AssertExpectations(t)
		mh.AssertExpectations(t)
	}
}

func Test_AlbumUsecase_CreateAlbum(t *testing.T) {
	t.Run("アルバムが正常に作成されること", func(t *testing.T) {
		mr, mh, usecase, cleanup := setupCreateAlbumUsecase(t)
		defer cleanup()

		album := domain.Album{
			ID:        1,
			Title:     "Test Album",
			Image:     "test.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mh.On("AlbumHandler", album).Return(nil)
		mr.On("CreateAlbum", &album).Return(nil)

		res, err := usecase.CreateAlbum(album)

		require.NoError(t, err)
		require.Equal(t, album.ID, res.ID)
		require.Equal(t, album.Title, res.Title)
		require.Equal(t, album.Image, res.Image)
	})
}
