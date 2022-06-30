package transaction

import (
	"errors"
	"golang-vue-nuxtjs-crowdfunding/campaign"
)

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository: repository, campaignRepository: campaignRepository}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	//TODO implement me
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}
	transactionByCampaignID, err := s.repository.FindTransactionByCampaignID(input.ID)
	if err != nil {
		return nil, err
	}

	return transactionByCampaignID, nil
}
