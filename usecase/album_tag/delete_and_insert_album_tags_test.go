package album_tag

import (
	"testing"

	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/assert"
)

func setupDeleteAndInsertAlbumTagsUsecase(t *testing.T) (*mock.MockAlbumTagRepository, *DeleteAndInsertAlbumTagsUsecase, func()) {
	re := new(mock.MockAlbumTagRepository)
	usecase := NewDeleteAndInsertAlbumTagsUsecase(re)

	return re, usecase, func() {
		re.AssertExpectations(t)
	}
}

func Test_AlbumTagUsecase_DeleteAndInsertAlbumTags(t *testing.T) {
	re, usecase, cleanup := setupDeleteAndInsertAlbumTagsUsecase(t)
	defer cleanup()

	albumId := uint(1)
	tagIds := []uint{2, 3, 4}

	re.On("DeleteAndInsert", albumId, tagIds).Return(nil).Once()

	err := usecase.DeleteAndInsertAlbumTags(albumId, tagIds)

	assert.NoError(t, err)
	re.AssertCalled(t, "DeleteAndInsert", albumId, tagIds)
}
