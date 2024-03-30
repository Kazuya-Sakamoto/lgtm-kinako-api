package mock

import (
	"lgtm-kinako-api/domain"

	"github.com/stretchr/testify/mock"
)

type MockAlbumTagUsecase struct {
	mock.Mock
}

func (mu *MockAlbumTagUsecase) DeleteAlbumTagsByTagID(tagId uint) error {
	args := mu.Called(tagId)
	return args.Error(0)
}

func (mu *MockAlbumTagUsecase) DeleteAndInsertAlbumTags(albumId uint, tagIds []uint) error {
	args := mu.Called(albumId, tagIds)
	return args.Error(0)
}

func (m *MockAlbumTagUsecase) GetAlbumCountsByTag() ([]domain.TagCount, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]domain.TagCount), args.Error(1)
	}
	return []domain.TagCount{}, args.Error(1)
}
