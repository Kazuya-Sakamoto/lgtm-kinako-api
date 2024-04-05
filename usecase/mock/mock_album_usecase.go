package mock

import (
	"lgtm-kinako-api/domain"

	"github.com/stretchr/testify/mock"
)

type MockAlbumUsecase struct {
	mock.Mock
}

func (m *MockAlbumUsecase) GetAllAlbums() ([]domain.AlbumResponse, error) {
	args := m.Called()
	return args.Get(0).([]domain.AlbumResponse), args.Error(1)
}

func (m *MockAlbumUsecase) GetRandomAlbums() ([]domain.AlbumResponse, error) {
	args := m.Called()
	return args.Get(0).([]domain.AlbumResponse), args.Error(1)
}

func (m *MockAlbumUsecase) GetAlbumsByTag(tagId uint) ([]domain.AlbumResponse, error) {
	args := m.Called(tagId)
	return args.Get(0).([]domain.AlbumResponse), args.Error(1)
}

func (m *MockAlbumUsecase) CreateAlbum(album domain.Album) (domain.AlbumResponse, error) {
	args := m.Called(album)
	return args.Get(0).(domain.AlbumResponse), args.Error(1)
}

func (m *MockAlbumUsecase) DeleteAlbum(userId, albumId uint) error {
	args := m.Called(userId, albumId)
	return args.Error(0)
}

func (m *MockAlbumUsecase) UploadImageToS3(data []byte) (string, error) {
	args := m.Called(data)
	return args.String(0), args.Error(1)
}
