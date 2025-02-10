package user

import (
	"errors"
	"reflect"
	"serinitystore/helper"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	GetUserById(ID int) (User, error)
	SaveOTPRequest(input ForgotPasswordUserInput) (OtpRequest, error)
	UpdateIsVerifiedOTP(inputData OTPUserInput) (OtpRequest, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Role = "buyer"
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(password)

	email := input.Email

	userByEmail, err := s.repository.FindByEmail(email)
	if err == nil && !reflect.DeepEqual(userByEmail, User{}) {
		return User{}, errors.New("email already in use")
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil

}

func (s *service) LoginUser(input LoginUserInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserById(ID int) (User, error) {
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, nil
	}

	if user.ID == 0 {
		return user, errors.New("no user found on with that ID")
	}

	return user, nil
}

func (s *service) SaveOTPRequest(input ForgotPasswordUserInput) (OtpRequest, error) {
	otpRequest := OtpRequest{}
	otpRequest.Email = input.Email
	otpRequest.Otp = helper.GenerateOTP()
	otpRequest.ExpiredAt = time.Now().Add(3 * time.Minute)
	otpRequest.IsVerified = false

	newOtpRequest, err := s.repository.SaveOTPRequest(otpRequest)
	if err != nil {
		return newOtpRequest, err
	}

	return newOtpRequest, nil
}

func (s *service) UpdateIsVerifiedOTP(inputData OTPUserInput) (OtpRequest, error) {
	otpRequest, err := s.repository.FindOTPByOTP(inputData.Otp)
	if err != nil {
		return otpRequest, err
	}

	otpRequest.Otp = inputData.Otp
	otpRequest.IsVerified = true

	updatedIsVerfied, err := s.repository.UpdateIsVerifiedOTP(otpRequest)
	if err != nil {
		return updatedIsVerfied, err
	}

	return updatedIsVerfied, nil
}
