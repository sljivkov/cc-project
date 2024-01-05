package repository

import (
	"worker/domain"

	"github.com/jinzhu/gorm"
)

type BorrowModel struct {
	gorm.Model
	Name            string
	Author          string
	ISBN            string
	DateOfBorrowing string
	UserId          uint
}

type BorrowRepository struct {
	db *gorm.DB
}

func NewBorrowRepository(db *gorm.DB) *BorrowRepository {
	return &BorrowRepository{
		db: db,
	}

}

func (repo *BorrowRepository) Borrow(borrow *domain.Borrow) error {
	borrowModel := BorrowModel{
		Name:            borrow.Name,
		Author:          borrow.Author,
		ISBN:            borrow.ISBN,
		DateOfBorrowing: borrow.DateOfBorrowing,
		UserId:          borrow.UserId,
	}
	repo.db.Create(&borrowModel)
	return nil
}
