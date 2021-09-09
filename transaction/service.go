package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
}

type service struct {
	repository   Repository
	campaignRepo campaign.Repository
}

func InstanceService(repository Repository, campaignRepo campaign.Repository) *service {
	return &service{repository: repository, campaignRepo: campaignRepo}
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