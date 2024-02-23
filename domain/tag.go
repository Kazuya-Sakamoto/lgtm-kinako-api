package domain

type Tag struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
}

type TagResponse struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
}

type ITagRepository interface {
	GetTags(tags *[]Tag) error
	CreateTag(tag *Tag) error
}
