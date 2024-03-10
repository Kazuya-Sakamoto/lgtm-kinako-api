package domain

type AlbumTag struct {
	ID      uint `json:"id" gorm:"primaryKey"`
	AlbumId uint `json:"album_id" gorm:"not null"`
	TagId   uint `json:"tag_id" gorm:"not null"`
}

type AlbumTagResponse struct {
	ID      uint `json:"id" gorm:"primaryKey"`
	AlbumId uint `json:"album_id" gorm:"not null"`
	TagId   uint `json:"tag_id" gorm:"not null"`
}

type IAlbumTagRepository interface {
	DeleteAlbumTagsByTagId(tagId uint) error
}
