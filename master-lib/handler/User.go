package handler

import (
	"master/application"
	"master/domain"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	user := &domain.User{}
	c.BindJSON(user)
	id, err := h.userService.Register(user)
	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user_id": id,
	})
}
