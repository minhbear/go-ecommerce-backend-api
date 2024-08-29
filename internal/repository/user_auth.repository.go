package repository

import (
	"fmt"
	"time"

	"github.com/go-ecommerce-backend-api/global"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64 ) error 
}

type userAuthRepository struct {}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}

func (u *userAuthRepository) AddOTP(email string, otp int, expirationTime int64 ) error {
	key := fmt.Sprintf("usr:%v:otp", email)

	return global.RDb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}