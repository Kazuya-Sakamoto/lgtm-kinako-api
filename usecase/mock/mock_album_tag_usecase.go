package mock

import (
	"github.com/stretchr/testify/mock"
)

type MockAlbumTagUsecase struct {
	mock.Mock
}

func (mu *MockAlbumTagUsecase) DeleteAlbumTagsByTagId(tagId uint) error {
	args := mu.Called(tagId)
	return args.Error(0)
}
