package application

import (
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
	return service.borrowRepository.Borrow(borrow)
}

func (service *BorrowService) Return(borrowId uint) error {
	return service.borrowRepository.Return(borrowId)
}
