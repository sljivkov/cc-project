package application

import (
	"master/domain"
	"master/repository"
)

type BookService struct {
	BookRepository repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{
		BookRepository: repo,
	}
}

func (service *BookService) Create(book *domain.Book) error {
	return service.BookRepository.Create(book)
}
