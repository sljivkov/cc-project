package handler

import (
	"log"
	"master/application"
	"master/domain"

	"github.com/gin-gonic/gin"
)

type BorrowHandler struct {
	borrowService *application.BorrowService
}

type ReturnBody struct {
	BorrowId uint `json:"borrow_id"`
}

func NewBorrowHandler(service *application.BorrowService) *BorrowHandler {
	return &BorrowHandler{
		borrowService: service,
	}
}

func (h *BorrowHandler) Borrow(c *gin.Context) {
	borrow := &domain.Borrow{}
	c.BindJSON(borrow)
	id, err := h.borrowService.Borrow(borrow)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "You can not borrow more books",
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
	log.Println(body)
	err := h.borrowService.Return(body.BorrowId)
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
