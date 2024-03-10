package mock

import (
	"github.com/stretchr/testify/mock"
)

/*
	repository
*/
type MockAlbumTagRepository struct {
	mock.Mock
}

func (m *MockAlbumTagRepository) DeleteAlbumTagsByTagId(tagId uint) error {
	args := m.Called(tagId)
	return args.Error(0)
}
