package campaign

import (
	"hermes/internal/contract"
)

type Service struct {
	repository Repository
}

func (service *Service) Create(dto contract.CreateCampaignDTO) (string, error) {
	campaign, _ := New(dto.Name, dto.Content, dto.Contacts)

	service.repository.Save(campaign)

	return campaign.ID, nil
}
