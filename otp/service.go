package otp

type Service interface {
	SaveOTP(input OTPRequest) (OTPRequest, error)
}

type otpService struct {
	repository Repository
}

func NewOTPService(repository Repository) *otpService {
	return &otpService{repository}
}

func (s *otpService) SaveOTP(input OTPRequest) (OTPRequest, error) {
	otpRequest := OTPRequest{
		Email:  input.Email,
		Otp:    input.Otp,
		Expiry: input.Expiry,
	}

	err := s.repository.SaveOTP(otpRequest)
	if err != nil {
		return otpRequest, err
	}

	return otpRequest, nil
}
