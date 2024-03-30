package mock

import (
	"lgtm-kinako-api/domain"

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

func (m *MockAlbumTagRepository) FindCountsByTag() ([]domain.TagCount, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]domain.TagCount), args.Error(1)
	}
	return []domain.TagCount{}, args.Error(1)
}
