package otp

type Service interface {
	SaveOTP(input OTPRequest) error
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
