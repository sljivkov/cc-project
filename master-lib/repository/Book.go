package repository

import (
	"errors"
	"master/domain"

	"github.com/jinzhu/gorm"
)

type BookModel struct {
	gorm.Model
	Name   string
	Author string
	ISBN   string `gorm:"primaryKey;autoIncrement:false"`
}

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) Create(book *domain.Book) error {
	bookModel := BookModel{
		Name:   book.Name,
		Author: book.Author,
		ISBN:   book.ISBN,
	}
	result := r.DB.Find(&bookModel)
	if result.RowsAffected > 0 {
		return errors.New("primary key constraint")
	}
	result = r.DB.Create(&bookModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *BookRepository) BookWithIsbn(isbn string) bool {
	bookModel := BookModel{}
	result := r.DB.Where("isbn = ?", isbn).First(&bookModel)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}
