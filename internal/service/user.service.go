package service

import (
	"github.com/go-ecommerce-backend-api/internal/repository"
	"github.com/go-ecommerce-backend-api/pkg/response"
)

type IUserService interface {
	Register(email string, password string) int
}

type userService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (userService *userService) Register(email string, password string) int {
	if !userService.userRepository.GetUserByEmail(email) {
		return response.ErrCodeUserHasExit
	}
	
	return response.ErrCodeSuccess
}