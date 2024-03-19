package domain

import "time"

type Album struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	Tags      []Tag     `json:"tags" gorm:"many2many:album_tags;"`
}

type AlbumResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Image     string    `json:"image"`
	Tags      []Tag     `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IAlbumRepository interface {
	FindAll(albums *[]Album) error
	FindRandom(albums *[]Album) error
	FindByTag(albums *[]Album, tagId uint) error
	Create(album *Album) error
	DeleteByAlbumID(userId uint, albumId uint) error
}
