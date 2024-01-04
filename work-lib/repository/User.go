package repository

import (
	"worker/domain"

	"github.com/jinzhu/gorm"
)

type UserModel struct {
	gorm.Model
	domain.User
}

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Register(user *UserModel) error {
	// call another service to check if he exist
	result := r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
