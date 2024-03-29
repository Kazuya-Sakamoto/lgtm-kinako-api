package album_tag

import (
	"testing"

	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/assert"
)

func setupDeleteAndInsertAlbumTagsUsecase(t *testing.T) (*mock.MockAlbumTagRepository, *DeleteAndInsertAlbumTagsUsecase, func()) {
	mr := new(mock.MockAlbumTagRepository)
	usecase := NewDeleteAndInsertAlbumTagsUsecase(mr)

	return mr, usecase, func() {
		mr.AssertExpectations(t)
	}
}

func Test_AlbumTagUsecase_DeleteAndInsertAlbumTags(t *testing.T) {
	mr, usecase, cleanup := setupDeleteAndInsertAlbumTagsUsecase(t)
	defer cleanup()

	albumId := uint(1)
	tagIds := []uint{2, 3, 4}

	mr.On("DeleteAndInsert", albumId, tagIds).Return(nil).Once()

	err := usecase.DeleteAndInsertAlbumTags(albumId, tagIds)

	assert.NoError(t, err)
	mr.AssertCalled(t, "DeleteAndInsert", albumId, tagIds)
}
