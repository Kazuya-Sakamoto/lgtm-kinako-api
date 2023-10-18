package repository

import (
	"fmt"
	"lgtm-kinako-api/domain"

	"gorm.io/gorm"
)

type albumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) IAlbumRepository {
	return &albumRepository{db}
}

func (ar *albumRepository) GetAllAlbums(albums *[]domain.Album) error {
	if err := ar.db.Order("created_at").Find(albums).Error; err != nil {
		return err
	}
	return nil
}

func (ar *albumRepository) GetRandomAlbums(albums *[]domain.Album) error {
	if err := ar.db.Order("RANDOM()").Limit(8).Find(albums).Error; err != nil {
			return err
	}
	return nil
}

func (ar *albumRepository) CreateAlbum(album *domain.Album) error {
	if err := ar.db.Create(album).Error; err != nil {
		return err
	}
	return nil
}

func (tr *albumRepository) DeleteAlbum(userId uint, albumId uint) error {
	result := tr.db.Where("id=? AND user_id=?", albumId, userId).Delete(&domain.Album{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}