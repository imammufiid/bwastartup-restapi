package transaction

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/payment"
	"errors"
)

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository     Repository
	campaignRepo   campaign.Repository
	paymentService payment.Service
}

func InstanceService(repository Repository, campaignRepo campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepo, paymentService}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {

	// get campaign
	campaign, err := s.campaignRepo.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	// check authorization
	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	campaign, err := s.campaignRepo.FindByID(input.CampaignID)
	if err != nil {
		return Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return Transaction{}, errors.New("not an owner of the campaign")
	}

	transaction := Transaction{
		Amount:     input.Amount,
		CampaignID: input.CampaignID,
		Status:     "pending",
		UserID:     input.User.ID,
		Code:       helper.GenerateCodeTransaction(input.User.ID),
	}

	// save transaction
	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	// mapping transaction
	paymentTransaction := payment.Transaction{
		ID: newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	// get payment url
	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	// add payment url to transaction
	newTransaction.PaymentURL = paymentURL

	// update transaction with add payment url
	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}
	
	return newTransaction, nil
}
