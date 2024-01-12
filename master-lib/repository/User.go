package repository

import (
	"master/domain"
	"fmt"
	"errors"
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

	userExist := UserModel{}
	// maybe test this later
	result := r.DB.Where("ssn = ?", user.SSN).First(&userExist)
	if result.RowsAffected > 0 {
		return 0, errors.New("primary key constraint")
	}
	fmt.Println(userExist)
	result  = r.DB.Create(&userModel)
	if result.Error != nil {
		return 0, result.Error
	}
	return userModel.ID, nil
}
