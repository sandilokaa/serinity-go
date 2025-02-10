package user

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OtpRequest struct {
	ID         int
	Email      string
	Otp        string
	IsVerified bool
	ExpiredAt  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (OtpRequest) TableName() string {
	return "OtpRequests"
}
