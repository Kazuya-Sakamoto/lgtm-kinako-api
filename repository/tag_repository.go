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

func (re *tagRepository) FindAll(tags *[]domain.Tag) error {
	if err := re.db.Find(tags).Error; err != nil {
		return err
	}
	return nil
}

func (re *tagRepository) Create(tag *domain.Tag) error {
	if err := re.db.Create(tag).Error; err != nil {
		return err
	}
	return nil
}

func (re *tagRepository) DeleteByTagID(tagId uint) error {
	result := re.db.Where("id=?", tagId).Delete(&domain.Tag{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
