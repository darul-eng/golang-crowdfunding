package campaign

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) FindCampaigns(userID int) ([]Campaign, error) {
	//TODO implement me
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
