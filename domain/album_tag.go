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
	DeleteByAlbumID(albumId uint) error
	Create(albumId uint, tagIds []uint) error
	ResetAndSet(albumId uint, tagIds []uint) error
}
