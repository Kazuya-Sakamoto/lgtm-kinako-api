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

type TagCount struct {
	TagID uint  `json:"tag_id" gorm:"primaryKey"`
	Count int64 `json:"count"  gorm:"not null"`
}

type IAlbumTagRepository interface {
	DeleteByTagID(tagId uint) error
	DeleteByAlbumID(albumId uint) error
	Create(albumId uint, tagIds []uint) error
	DeleteAndInsert(albumId uint, tagIds []uint) error
	FindCountsByTag() ([]TagCount, error)
}
