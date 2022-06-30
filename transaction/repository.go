package transaction

import "gorm.io/gorm"

type Repository interface {
	FindTransactionByCampaignID(campaignID int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindTransactionByCampaignID(campaignID int) ([]Transaction, error) {
	//TODO implement me
	var transactions []Transaction
	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, err
}
