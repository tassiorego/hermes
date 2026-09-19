package campaign

import (
	"hermes/internal/contract"
)

type Service struct {
	repository Repository
}

func (service *Service) Create(dto contract.CreateCampaignDTO) (string, error) {
	campaign, err := New(dto.Name, dto.Content, dto.Contacts)

	if err != nil {
		return "", err
	}

	err = service.repository.Save(campaign)

	if err != nil {
		return "", err
	}

	return campaign.ID, nil
}
