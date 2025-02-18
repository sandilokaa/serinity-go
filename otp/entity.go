package otp

import "time"

type OTPRequest struct {
	Email  string    `json:"email" binding:"required"`
	Otp    string    `json:"otp"`
	Expiry time.Time `json:"expiry"`
}
