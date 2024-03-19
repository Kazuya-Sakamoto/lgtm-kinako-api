package mock

import (
	"lgtm-kinako-api/domain"

	"github.com/stretchr/testify/mock"
)

/*
repository
*/
type MockAlbumRepository struct {
	mock.Mock
}

func (m *MockAlbumRepository) FindAll(albums *[]domain.Album) error {
	args := m.Called(albums)
	if args.Get(0) != nil {
		*albums = args.Get(0).([]domain.Album)
	}
	return args.Error(0)
}

func (m *MockAlbumRepository) FindRandom(albums *[]domain.Album) error {
	args := m.Called(albums)
	if args.Get(0) != nil {
		*albums = args.Get(0).([]domain.Album)
	}
	return args.Error(0)
}

func (m *MockAlbumRepository) FindByTag(albums *[]domain.Album, tagId uint) error {
	args := m.Called(albums)
	if args.Get(0) != nil {
		*albums = args.Get(0).([]domain.Album)
	}
	return args.Error(0)
}

func (m *MockAlbumRepository) Create(album *domain.Album) error {
	args := m.Called(album)
	return args.Error(0)
}

func (m *MockAlbumRepository) DeleteByAlbumID(userId, albumId uint) error {
	args := m.Called(albumId)
	return args.Error(0)
}

/*
handler
*/
type MockAlbumHandler struct {
	mock.Mock
}

func (m *MockAlbumHandler) AlbumHandler(album domain.Album) error {
	args := m.Called(album)
	return args.Error(0)
}
