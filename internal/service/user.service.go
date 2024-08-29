package service

import (
	"fmt"
	"time"

	"github.com/go-ecommerce-backend-api/global"
	"github.com/go-ecommerce-backend-api/internal/repository"
	"github.com/go-ecommerce-backend-api/pkg/response"
	"github.com/go-ecommerce-backend-api/pkg/utils/crypto"
	"github.com/go-ecommerce-backend-api/pkg/utils/random"
	sendto "github.com/go-ecommerce-backend-api/pkg/utils/send_to"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepository repository.IUserRepository
	userAuthRepository repository.IUserAuthRepository
}

func NewUserService(
	userRepository repository.IUserRepository, 
	userAuthRepository repository.IUserAuthRepository,
) IUserService {
	return &userService{
		userRepository: userRepository,
		userAuthRepository: userAuthRepository, 
	}
}

func (userService *userService) Register(email string, purpose string) int {
	if userService.userRepository.GetUserByEmail(email) {
		return response.ErrCodeUserHasExit
	}

	hashEmail := crypto.GetHash(email)
	fmt.Printf("hashEmail: %s", hashEmail)
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}
	fmt.Printf("otp: %d", otp)

	err := userService.userAuthRepository.AddOTP(hashEmail, otp, int64(10 * time.Minute))
	if err != nil {
		return response.ErrInvalidOTP 
	}

	err = sendto.SendTextEmailOtp([]string{email}, global.Config.SMTP.User, fmt.Sprintf("%d", otp))
	if err != nil {
		return response.ErrInvalidOTP
	}

	return response.ErrCodeSuccess
}