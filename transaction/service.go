package transaction

import (
	"cheggstore/cloth"
	"cheggstore/payment"
	"fmt"
	"strconv"
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

	transaction := Transaction{}
	transaction.ClothID = input.ClothID
	transaction.UserID = input.User.ID
	transaction.Quantity = input.Quantity
	transaction.Amount = price * input.Quantity
	transaction.Status = "pending"
	transaction.Code = "12345"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	err = s.clothRepository.UpdateStockByClothID(cloth.ID, cloth.Stock-transaction.Quantity)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
