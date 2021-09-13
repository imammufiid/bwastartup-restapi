package transaction

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/payment"
	"errors"
	"strconv"
)

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
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
		ID:     newTransaction.ID,
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

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transactionID, _ := strconv.Atoi(input.OrderID)

	// get transaction by order id from midtrans
	transaction, err := s.repository.GetByID(transactionID)
	if err != nil {
		return err
	}

	// change status payment in the db
	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "canceled"
	}

	// update status payment in the db
	updatedTrx, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	// get campaign
	campaign, err := s.campaignRepo.FindByID(updatedTrx.CampaignID)
	if err != nil {
		return nil
	}

	if updatedTrx.Status == "paid" {
		// update field backer count
		campaign.BackerCount += 1
		campaign.CurrentAmount += updatedTrx.Amount

		// update campaign
		_, err := s.campaignRepo.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}
