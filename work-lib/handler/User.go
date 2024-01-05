package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"worker/domain"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	UserID uint `json:"user_id"`
}

type UserHandler struct {
	ServerUrl string
}

func (h *UserHandler) Register(c *gin.Context) {
	user := &domain.User{}
	c.BindJSON(user)
	targetUrl, _ := url.Parse("http://" + h.ServerUrl + "/user/register")
	user.LastName = user.LastName + "worker"
	jsonPayload, _ := json.Marshal(user)
	response, err := http.Post(targetUrl.String(), "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if response.StatusCode == 400 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "This user already exist"})
		return
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var apiResponse APIResponse
	if err := json.Unmarshal(responseBody, &apiResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"user_id": apiResponse.UserID,
	})
}
