package otp

import (
	"fmt"
)

type Service interface {
	SaveOTP(input OTPRequest) error
	VerifyOTP(email, userOTP string) (bool, error)
}

type otpService struct {
	repository Repository
}

func NewOTPService(repository Repository) *otpService {
	return &otpService{repository}
}

func (s *otpService) SaveOTP(input OTPRequest) error {
	otpRequest := OTPRequest{
		Email:  input.Email,
		Otp:    input.Otp,
		Expiry: input.Expiry,
	}

	return s.repository.SaveOTP(otpRequest)
}

func (s *otpService) VerifyOTP(email, userOTP string) (bool, error) {
	storedOTP, err := s.repository.GetOTP(email)
	if err != nil {
		return false, fmt.Errorf("failed to get OTP: %v", err)
	}

	if userOTP == storedOTP {
		return true, nil
	}

	return false, nil
}
