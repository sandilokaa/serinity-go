package transaction

import (
	"cheggstore/cloth"
	"cheggstore/helper"
	"cheggstore/payment"
	"fmt"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type service struct {
	repository      Repository
	clothRepository cloth.Repository
	paymentService  payment.Service
}

type Service interface {
	FindAllTransaction(search string) ([]Transaction, error)
	GetTransactionByUserID(userID int, requestedUserID int) ([]Transaction, error)
	GetTransactionByID(input TransactionInputDetail) (Transaction, error)
	GetTransactionUserIDByID(input TransactionInputDetail, userID int) (Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
}

func NewService(repository Repository, clothRepository cloth.Repository, paymentService payment.Service) *service {
	return &service{repository, clothRepository, paymentService}
}

func (s *service) FindAllTransaction(search string) ([]Transaction, error) {
	transactions, err := s.repository.FindAllTransaction(search)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionByUserID(userID int, requestedUserID int) ([]Transaction, error) {

	if userID != requestedUserID {
		return nil, fmt.Errorf("access denied")
	}

	transactions, err := s.repository.GetTransactionByUserID(requestedUserID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionByID(input TransactionInputDetail) (Transaction, error) {
	transaction, err := s.repository.GetTransactionByID(input.ID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) GetTransactionUserIDByID(input TransactionInputDetail, userID int) (Transaction, error) {

	transaction, err := s.repository.GetTransactionUserIDByID(input.ID, userID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	cloth, err := s.clothRepository.FindClothByID(input.ClothID)
	if err != nil {
		return Transaction{}, err
	}

	if cloth.Price == "" {
		return Transaction{}, fmt.Errorf("price cannot be empty")
	}

	price, err := strconv.Atoi(cloth.Price)
	if err != nil {
		return Transaction{}, err
	}

	clothVariation, err := s.clothRepository.FindClothVariationByID(input.ClothVariationID)
	if err != nil {
		return Transaction{}, err
	}

	transaction := Transaction{}
	transaction.ClothID = input.ClothID
	transaction.ClothVariationID = input.ClothVariationID
	transaction.UserID = input.User.ID
	transaction.Quantity = input.Quantity
	transaction.Amount = price * input.Quantity
	transaction.Status = "pending"
	transaction.Code = helper.GenerateTransactionCode()

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
		Cloth:  payment.ClothDetail{Name: cloth.Name},
	}

	itemDetail := []midtrans.ItemDetail{
		{
			ID:    strconv.Itoa(cloth.ID),
			Name:  fmt.Sprintf("%s - %s - %s", cloth.Name, clothVariation.Color, clothVariation.Size),
			Price: int64(price),
			Qty:   int32(input.Quantity),
		},
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User, itemDetail)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetTransactionByID(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	clothVariation, err := s.clothRepository.FindClothVariationByID(transaction.ClothVariationID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {

		err = s.clothRepository.UpdateStockByClothID(clothVariation.ID, clothVariation.Stock-transaction.Quantity)
		if err != nil {
			return err
		}
	}

	return nil
}
