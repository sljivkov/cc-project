package application

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"worker/domain"
	"worker/repository"
)

type BorrowId struct {
	BorrowId uint `json:"borrow_id"`
}
type BorrowService struct {
	repo      *repository.BorrowRepository
	serverUrl string
}

func NewBorrowService(url string, repo *repository.BorrowRepository) *BorrowService {
	return &BorrowService{
		repo:      repo,
		serverUrl: url,
	}
}

func (service *BorrowService) Borrow(borrow *domain.Borrow) (uint, error) {
	targetUrl, _ := url.Parse("http://" + service.serverUrl + "/borrow")
	jsonPayload, _ := json.Marshal(borrow)
	response, err := http.Post(targetUrl.String(), "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return 0, err
	}
	if response.StatusCode == 400 {
		return 0, errors.New("you can't borrow anymore")
	}

	responseBody, _ := io.ReadAll(response.Body)

	var apiResponse BorrowId
	if err := json.Unmarshal(responseBody, &apiResponse); err != nil {
		return 0, err
	}
	err = service.repo.Borrow(borrow)
	if err != nil {
		return 0, err
	}
	return apiResponse.BorrowId, nil
}

func (service *BorrowService) Return(borrowId uint) error {
	id := BorrowId{
		BorrowId: borrowId,
	}
	jsonPayload, _ := json.Marshal(id)
	targetUrl, _ := url.Parse("http://" + service.serverUrl + "/return")
	response, _ := http.Post(targetUrl.String(), "application/json", bytes.NewBuffer(jsonPayload))
	if response.StatusCode == 400 {
		return errors.New("you can't return ")
	}
	return nil
}
