package otp

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Repository interface {
	SaveOTP(otpRequest OTPRequest) error
}

type otpRepository struct {
	rdb *redis.Client
	ctx context.Context
}

func NewOTPRepository(rdb *redis.Client) *otpRepository {
	return &otpRepository{rdb: rdb, ctx: context.Background()}
}

func (r *otpRepository) SaveOTP(otpRequest OTPRequest) (OTPRequest, error) {
	err := r.rdb.Set(r.ctx, "otp:"+otpRequest.Email, otpRequest.Otp, time.Until(otpRequest.Expiry)).Err()
	if err != nil {
		return otpRequest, err
	}

	return otpRequest, nil
}
