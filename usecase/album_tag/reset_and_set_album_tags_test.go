package album_tag

import (
	"testing"

	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/assert"
)

func setupResetAndSetAlbumTagsUsecase(t *testing.T) (*mock.MockAlbumTagRepository, *ResetAndSetAlbumTagsUsecase, func()) {
	mr := new(mock.MockAlbumTagRepository)
	usecase := NewResetAndSetAlbumTagsUsecase(mr)

	return mr, usecase, func() {
		mr.AssertExpectations(t)
	}
}

func Test_AlbumTagUsecase_ResetAndSetAlbumTags(t *testing.T) {
	mr, usecase, cleanup := setupResetAndSetAlbumTagsUsecase(t)
	defer cleanup()

	albumId := uint(1)
	tagIds := []uint{2, 3, 4}

	mr.On("ResetAndSet", albumId, tagIds).Return(nil).Once()

	err := usecase.ResetAndSetAlbumTags(albumId, tagIds)

	assert.NoError(t, err)
	mr.AssertCalled(t, "ResetAndSet", albumId, tagIds)
}
