package user

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ForgotPasswordUserInput struct {
	Email string `json:"email" binding:"required"`
}

type OTPUserInput struct {
	Otp string `json:"otp" binding:"required"`
}
