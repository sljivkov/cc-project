package repository

import (
	"errors"
	"master/domain"

	"github.com/jinzhu/gorm"
)

type UserModel struct {
	gorm.Model
	FirstName string
	LastName  string
	SSN       string `gorm:"primaryKey;autoIncrement:false"`
	Address   string
}

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Register(user *domain.User) (uint, error) {
	userModel := UserModel{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		SSN:       user.SSN,
		Address:   user.Address,
	}
	// maybe test this later
	result := r.DB.Find(&userModel)
	if result.RowsAffected > 0 {
		return 0, errors.New("primary key constraint")
	}
	result = r.DB.Create(&userModel)
	if result.Error != nil {
		return 0, result.Error
	}
	return userModel.ID, nil
}
