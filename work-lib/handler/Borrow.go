package handler

import (
	"worker/application"
	"worker/domain"

	"github.com/gin-gonic/gin"
)

type ReturnBody struct {
	BorrowId uint `json:"borrow_id"`
}
type BorrowHandler struct {
	service *application.BorrowService
}

func NewBorrowHandler(service *application.BorrowService) *BorrowHandler {
	return &BorrowHandler{
		service: service,
	}
}
func (h *BorrowHandler) Borrow(c *gin.Context) {
	borrow := &domain.Borrow{}
	c.BindJSON(borrow)
	id, err := h.service.Borrow(borrow)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"borrow_id": id,
	})
}
func (h *BorrowHandler) Return(c *gin.Context) {
	body := ReturnBody{}
	c.BindJSON(&body)

	err := h.service.Return(body.BorrowId)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "There was an error",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "Book returned",
	})
}
