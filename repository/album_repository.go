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

func (re *albumRepository) FindAll(albums *[]domain.Album) error {
	if err := re.db.
		Joins("INNER JOIN users AS u ON u.id = albums.user_id").
		Order("albums.created_at").
		Preload("Tags").
		Find(albums).Error; err != nil {
		return err
	}
	return nil
}

func (re *albumRepository) FindRandom(albums *[]domain.Album) error {
	if err := re.db.
		Preload("Tags").
		Order("RAND()").Limit(10).
		Find(albums).Error; err != nil {
		return err
	}
	return nil
}

func (re *albumRepository) FindByTag(albums *[]domain.Album, tagId uint) error {
	if err := re.db.
		Joins("JOIN album_tags ON album_tags.album_id = albums.id").
		Where("album_tags.tag_id = ?", tagId).
		Preload("Tags").
		Order("RAND()").
		Find(albums).Error; err != nil {
		return err
	}
	return nil
}

func (re *albumRepository) Create(album *domain.Album) error {
	if err := re.db.Create(album).Error; err != nil {
		return err
	}
	return nil
}

func (re *albumRepository) DeleteByAlbumID(userId uint, albumId uint) error {
	result := re.db.Where("id=? AND user_id=?", albumId, userId).Delete(&domain.Album{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
