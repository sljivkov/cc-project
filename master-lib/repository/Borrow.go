package repository

import (
	"errors"
	"log"
	"master/domain"
	"time"

	"github.com/jinzhu/gorm"
)

type BorrowModel struct {
	gorm.Model
	ISBN            string
	DateOfBorrowing string
	DateOfReturning string
	UserId          uint `gorm:"user_id"`
}

type BorrowRepository struct {
	db *gorm.DB
}

func NewBorrowRepository(db *gorm.DB) *BorrowRepository {
	return &BorrowRepository{
		db: db,
	}
}

func (repo *BorrowRepository) Borrow(borrow *domain.Borrow) (uint, error) {
	if !repo.valid_borrow(borrow) {
		log.Println("too much books")
		return 0, errors.New("too much books")
	}
	borrowModel := BorrowModel{
		ISBN:            borrow.ISBN,
		DateOfBorrowing: time.Now().String(),
		UserId:          borrow.UserId,
	}
	repo.db.Create(&borrowModel)
	return borrowModel.ID, nil
}

func (repo *BorrowRepository) Return(borrowId uint) error {
	borrow := BorrowModel{}
	result := repo.db.First(&borrow, borrowId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("there is no borrow with that id")
		return result.Error
	}
	borrow.DateOfReturning = time.Now().String()
	repo.db.Save(&borrow)
	return nil
}

func (repo *BorrowRepository) valid_borrow(borrow *domain.Borrow) bool {
	borrows := []BorrowModel{}
	// dodati date of returnign proveru
	repo.db.Where("user_id = ? and date_of_returning = ?", borrow.UserId, "").Find(&borrows)
	if len(borrows) >= 3 {
		log.Println("too much books")
		return false
	}
	return true
}
