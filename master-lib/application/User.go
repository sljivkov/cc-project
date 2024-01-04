package application

import (
	"master/domain"
	"master/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}

func (service *UserService) Register(user *domain.User) (uint, error) {
	return service.UserRepository.Register(user)
}
