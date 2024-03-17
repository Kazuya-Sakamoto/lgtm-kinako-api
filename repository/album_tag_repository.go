package repository

import (
	"lgtm-kinako-api/domain"

	"gorm.io/gorm"
)

type albumtagRepository struct {
	db *gorm.DB
}

func NewAlbumTagRepository(db *gorm.DB) IAlbumtagRepository {
	return &albumtagRepository{db}
}

func (atr *albumtagRepository) DeleteAlbumTagsByTagId(tagId uint) error {
	result := atr.db.Where("tag_id=?", tagId).Delete(&domain.AlbumTag{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (atr *albumtagRepository) DeleteByAlbumID(albumId uint) error {
	return atr.db.Where("album_id = ?", albumId).Delete(&domain.AlbumTag{}).Error
}

func (atr *albumtagRepository) Create(albumId uint, tagIds []uint) error {
	for _, tagId := range tagIds {
		albumTag := domain.AlbumTag{AlbumId: albumId, TagId: tagId}
		if err := atr.db.Create(&albumTag).Error; err != nil {
			return err
		}
	}
	return nil
}

func (atr *albumtagRepository) ResetAndSet(albumId uint, tagIds []uint) error {
	return atr.db.Transaction(func(tx *gorm.DB) error {
		if err := atr.DeleteByAlbumID(albumId); err != nil {
			return err
		}
		if err := atr.Create(albumId, tagIds); err != nil {
			return err
		}

		return nil
	})
}
