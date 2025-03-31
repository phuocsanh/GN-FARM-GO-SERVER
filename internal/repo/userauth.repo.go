package repo

import (
	"fmt"
	"go_ecommerce/global"
	"time"
)

type IUserAuthRepository interface {
	AddOtp(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct {
}

// AddOtp implements IUserAuthRepository.
func (u *userAuthRepository) AddOtp(email string, otp int, expirationTime int64) error {
	// panic("unimplemented")
	key:= fmt.Sprintf("usr:%s:otp", email)
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
