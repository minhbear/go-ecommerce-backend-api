package service

import "github.com/go-ecommerce-backend-api/internal/repository"

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repository.NewUserRepository(),
	}
}

func (userService *UserService) GetInfoUser() string {
	return userService.userRepository.GetInfoUser()
}
