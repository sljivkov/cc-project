package application

import (
	"errors"
	"master/domain"
	"master/repository"
)

type BorrowService struct {
	borrowRepository *repository.BorrowRepository
	bookRepository   *repository.BookRepository
}

func NewBorrowService(repo *repository.BorrowRepository, bookRepo *repository.BookRepository) *BorrowService {
	return &BorrowService{
		borrowRepository: repo,
		bookRepository:   bookRepo,
	}
}

func (service *BorrowService) Borrow(borrow *domain.Borrow) (uint, error) {
	if service.bookRepository.BookWithIsbn(borrow.ISBN) {
		return service.borrowRepository.Borrow(borrow)
	}
	return 0, errors.New("No book with that isbn")
}

func (service *BorrowService) Return(borrowId uint) error {
	return service.borrowRepository.Return(borrowId)
}
