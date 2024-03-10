package mock

import (
	"lgtm-kinako-api/domain"

	"github.com/stretchr/testify/mock"
)

type MockAlbumUsecase struct {
	mock.Mock
}

func (mu *MockAlbumUsecase) GetAllAlbums(userId uint) ([]domain.AlbumResponse, error) {
	args := mu.Called(userId)
	return args.Get(0).([]domain.AlbumResponse), args.Error(1)
}

func (mu *MockAlbumUsecase) GetRandomAlbums() ([]domain.AlbumResponse, error) {
	args := mu.Called()
	return args.Get(0).([]domain.AlbumResponse), args.Error(1)
}

func (mu *MockAlbumUsecase) GetAlbumsByTag(tagId uint) ([]domain.AlbumResponse, error) {
	args := mu.Called(tagId)
	return args.Get(0).([]domain.AlbumResponse), args.Error(1)
}

func (mu *MockAlbumUsecase) CreateAlbum(album domain.Album) (domain.AlbumResponse, error) {
	args := mu.Called(album)
	return args.Get(0).(domain.AlbumResponse), args.Error(1)
}

func (mu *MockAlbumUsecase) DeleteAlbum(userId, albumId uint) error {
	args := mu.Called(userId, albumId)
	return args.Error(0)
}

func (mu *MockAlbumUsecase) UploadImageToS3(data []byte) (string, error) {
	args := mu.Called(data)
	return args.String(0), args.Error(1)
}
