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

func (tr *albumtagRepository) DeleteAlbumTagsByTagId(tagId uint) error {
	result := tr.db.Where("tag_id=?", tagId).Delete(&domain.AlbumTag{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
