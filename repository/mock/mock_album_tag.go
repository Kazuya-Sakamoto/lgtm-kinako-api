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

func (m *MockAlbumTagRepository) DeleteByTagID(tagId uint) error {
	args := m.Called(tagId)
	return args.Error(0)
}

func (m *MockAlbumTagRepository) DeleteByAlbumID(albumId uint) error {
	args := m.Called(albumId)
	return args.Error(0)
}

func (m *MockAlbumTagRepository) Create(albumId uint, tagIds []uint) error {
	args := m.Called(albumId, tagIds)
	return args.Error(0)
}

func (m *MockAlbumTagRepository) DeleteAndInsert(albumId uint, tagIds []uint) error {
	args := m.Called(albumId, tagIds)
	return args.Error(0)
}
