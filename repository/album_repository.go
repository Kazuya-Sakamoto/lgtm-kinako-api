package repository

import (
	"lgtm-kinako-api/model"

	"gorm.io/gorm"
)

type IAlbumRepository interface {
	GetAllAlbums(albums *[]model.Album) error
	CreateAlbum(task *model.Album) error
}

type albumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) IAlbumRepository {
	return &albumRepository{db}
}

func (ar *albumRepository) GetAllAlbums(albums *[]model.Album) error {
	if err := ar.db.Order("created_at").Find(albums).Error; err != nil {
		return err
	}
	return nil
}

func (ar *albumRepository) CreateAlbum(album *model.Album) error {
	if err := ar.db.Create(album).Error; err != nil {
		return err
	}
	return nil
}