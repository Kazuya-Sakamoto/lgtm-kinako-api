package repository

import (
	"lgtm-kinako-api/domain"

	"gorm.io/gorm"
)

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) ITagRepository {
	return &tagRepository{db}
}

func (tr *tagRepository) GetTags(tags *[]domain.Tag) error {
	if err := tr.db.Find(tags).Error; err != nil {
		return err
	}
	return nil
}

func (tr *tagRepository) CreateTag(tag *domain.Tag) error {
	if err := tr.db.Create(tag).Error; err != nil {
		return err
	}
	return nil
}
