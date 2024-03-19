package mock

import (
	"github.com/stretchr/testify/mock"
)

type MockAlbumTagUsecase struct {
	mock.Mock
}

func (mu *MockAlbumTagUsecase) DeleteAlbumTagsByTagID(tagId uint) error {
	args := mu.Called(tagId)
	return args.Error(0)
}

func (mu *MockAlbumTagUsecase) ResetAndSetAlbumTags(albumId uint, tagIds []uint) error {
	args := mu.Called(albumId, tagIds)
	return args.Error(0)
}
