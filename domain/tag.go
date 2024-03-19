package domain

type Tag struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}

type TagResponse struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}

type ITagRepository interface {
	FindAll(tags *[]Tag) error
	Create(tag *Tag) error
	DeleteByTagID(tagId uint) error
}
