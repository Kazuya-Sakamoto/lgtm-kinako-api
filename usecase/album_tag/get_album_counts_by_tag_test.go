package album_tag

import (
	"testing"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupGetAlbumCountsByTagUsecase(t *testing.T) (*mock.MockAlbumTagRepository, *GetAlbumCountsByTagUsecase, func()) {
	re := new(mock.MockAlbumTagRepository)
	usecase := NewGetAlbumCountsByTagUsecase(re)

	return re, usecase, func() {
		re.AssertExpectations(t)
	}
}

func Test_GetAlbumCountsByTagUsecase_GetAlbumCountsByTag(t *testing.T) {
	re, usecase, cleanup := setupGetAlbumCountsByTagUsecase(t)
	defer cleanup()

	expectedCounts := []domain.TagCount{
		{TagID: 1, Count: 5},
		{TagID: 2, Count: 3},
		{TagID: 3, Count: 4},
		{TagID: 4, Count: 0},
	}

	re.On("FindCountsByTag").Return(expectedCounts, nil).Once()

	counts, err := usecase.GetAlbumCountsByTag()

	require.NoError(t, err)
	assert.Equal(t, expectedCounts, counts)
	re.AssertCalled(t, "FindCountsByTag")
}
