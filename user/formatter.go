package user

type UserFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return formatter
}

type OTPFormatter struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Otp        string `json:"otp"`
	IsVerified bool   `json:"is_verified"`
}

func FormatOTP(otpRequest OtpRequest) OTPFormatter {
	formatter := OTPFormatter{
		ID:         otpRequest.ID,
		Email:      otpRequest.Email,
		Otp:        otpRequest.Otp,
		IsVerified: otpRequest.IsVerified,
	}

	return formatter
}
