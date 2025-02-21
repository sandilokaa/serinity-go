package otp

type VerifyOTPInput struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}
