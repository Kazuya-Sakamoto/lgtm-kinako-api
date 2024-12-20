package album

import (
	"testing"
	"time"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func setupGetRandomAlbumsUsecase(t *testing.T) (*mock.MockAlbumRepository, *GetRandomAlbumsUsecase, func()) {
	re := new(mock.MockAlbumRepository)
	usecase := NewGetRandomAlbumsUsecase(re)

	return re, usecase, func() {
		re.AssertExpectations(t)
	}
}

func Test_AlbumUsecase_GetRandomAlbums(t *testing.T) {
	t.Run("正常にアルバムが取得できること", func(t *testing.T) {
		re, usecase, cleanup := setupGetRandomAlbumsUsecase(t)
		defer cleanup()

		ea := []domain.Album{
			{
				ID:        1,
				Title:     "Title1",
				Image:     "image1.jpg",
				Tags:      []domain.Tag{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:        2,
				Title:     "Title2",
				Image:     "image2.jpg",
				Tags:      []domain.Tag{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		re.On("FindRandom", testify_mock.AnythingOfType("*[]domain.Album")).Return(nil).Run(func(args testify_mock.Arguments) {
			arg := args.Get(0).(*[]domain.Album)
			*arg = ea
		})

		res, err := usecase.GetRandomAlbums()

		require.NoError(t, err)
		assert.Equal(t, len(ea), len(res))
		for i, r := range res {
			assert.Equal(t, ea[i].ID, r.ID)
			assert.Equal(t, ea[i].Title, r.Title)
			assert.Equal(t, ea[i].Image, r.Image)
		}
	})
}
