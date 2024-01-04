package handler

import (
	"master/application"
	"master/domain"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService *application.BookService
}

func NewBookHandler(service *application.BookService) *BookHandler {
	return &BookHandler{
		bookService: service,
	}
}

func (h *BookHandler) Create(c *gin.Context) {
	book := &domain.Book{}
	c.BindJSON(book)
	err := h.bookService.Create(book)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "Book already exist",
		})
		return
	}
	c.JSON(200, book)
}
