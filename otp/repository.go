package otp

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Repository interface {
	SaveOTP(otpRequest OTPRequest) error
	GetOTP(email string) (string, error)
}

type otpRepository struct {
	rdb *redis.Client
	ctx context.Context
}

func NewOTPRepository(rdb *redis.Client) *otpRepository {
	return &otpRepository{rdb: rdb, ctx: context.Background()}
}

func (r *otpRepository) SaveOTP(otpRequest OTPRequest) error {
	err := r.rdb.Set(r.ctx, "otp:"+otpRequest.Email, otpRequest.Otp, time.Until(otpRequest.Expiry)).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *otpRepository) GetOTP(email string) (string, error) {
	otp, err := r.rdb.Get(r.ctx, "otp:"+email).Result()
	if err != nil {
		return "", err
	}

	return otp, nil
}
