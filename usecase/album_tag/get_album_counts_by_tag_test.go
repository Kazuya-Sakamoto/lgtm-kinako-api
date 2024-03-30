package album_tag

import (
	"testing"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupGetAlbumCountsByTagUsecase(t *testing.T) (*mock.MockAlbumTagRepository, *GetAlbumCountsByTagUsecase, func()) {
	mr := new(mock.MockAlbumTagRepository)
	usecase := NewGetAlbumCountsByTagUsecase(mr)

	return mr, usecase, func() {
		mr.AssertExpectations(t)
	}
}

func Test_GetAlbumCountsByTagUsecase_GetAlbumCountsByTag(t *testing.T) {
	mr, usecase, cleanup := setupGetAlbumCountsByTagUsecase(t)
	defer cleanup()

	expectedCounts := []domain.TagCount{
		{TagID: 1, Count: 5},
		{TagID: 2, Count: 3},
		{TagID: 3, Count: 4},
		{TagID: 4, Count: 0},
	}

	mr.On("FindAlbumCountsByTag").Return(expectedCounts, nil).Once()

	counts, err := usecase.GetAlbumCountsByTag()

	require.NoError(t, err)
	assert.Equal(t, expectedCounts, counts)
	mr.AssertCalled(t, "FindAlbumCountsByTag")
}
