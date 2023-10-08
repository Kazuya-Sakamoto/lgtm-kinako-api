package mock

import (
	"lgtm-kinako-api/model"

	"github.com/stretchr/testify/mock"
)

type MockAlbumUsecase struct {
    mock.Mock
}

func (mu *MockAlbumUsecase) GetAllAlbums() ([]model.AlbumResponse, error) {
    args := mu.Called()
    return args.Get(0).([]model.AlbumResponse), args.Error(1)
}

func (mu *MockAlbumUsecase) GetRandomAlbums() ([]model.AlbumResponse, error) {
    args := mu.Called()
    return args.Get(0).([]model.AlbumResponse), args.Error(1)
}

func (mu *MockAlbumUsecase) CreateAlbum(album model.Album) (model.AlbumResponse, error) {
    args := mu.Called(album)
    return args.Get(0).(model.AlbumResponse), args.Error(1)
}

func (mu *MockAlbumUsecase) DeleteAlbum(userId, albumId uint) error {
    args := mu.Called(userId, albumId)
    return args.Error(0)
}

func (mu *MockAlbumUsecase) UploadImageToS3(data []byte) (string, error) {
    args := mu.Called(data)
    return args.String(0), args.Error(1)
}