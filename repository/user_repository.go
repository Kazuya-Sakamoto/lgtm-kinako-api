package repository

import (
	"lgtm-kinako-api/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (re *userRepository) FindByEmail(user *domain.User, email string) error {
	if err := re.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (re *userRepository) Create(user *domain.User) error {
	if err := re.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
