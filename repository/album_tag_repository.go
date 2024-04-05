package repository

import (
	"lgtm-kinako-api/domain"

	"gorm.io/gorm"
)

type albumTagRepository struct {
	db *gorm.DB
}

func NewAlbumTagRepository(db *gorm.DB) IAlbumTagRepository {
	return &albumTagRepository{db}
}

func (re *albumTagRepository) DeleteByTagID(tagId uint) error {
	result := re.db.Where("tag_id=?", tagId).Delete(&domain.AlbumTag{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (re *albumTagRepository) DeleteByAlbumID(albumId uint) error {
	return re.db.Where("album_id = ?", albumId).Delete(&domain.AlbumTag{}).Error
}

func (re *albumTagRepository) Create(albumId uint, tagIds []uint) error {
	for _, tagId := range tagIds {
		albumTag := domain.AlbumTag{AlbumId: albumId, TagId: tagId}
		if err := re.db.Create(&albumTag).Error; err != nil {
			return err
		}
	}
	return nil
}

func (re *albumTagRepository) DeleteAndInsert(albumId uint, tagIds []uint) error {
	return re.db.Transaction(func(tx *gorm.DB) error {
		if err := re.DeleteByAlbumID(albumId); err != nil {
			return err
		}
		if err := re.Create(albumId, tagIds); err != nil {
			return err
		}

		return nil
	})
}

func (re *albumTagRepository) FindCountsByTag() ([]domain.TagCount, error) {
	var counts []domain.TagCount
	if err := re.db.
		Table("album_tags").
		Select("tag_id, COUNT(album_id) AS count").
		Group("tag_id").
		Scan(&counts).Error; err != nil {
		return nil, err
	}
	return counts, nil
}
