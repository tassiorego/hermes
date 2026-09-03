package campaign

import "hermes/internal/contract"

type Service struct {
	repository Repository
}

func (service *Service) Create(dto contract.CreateCampaignDTO) error {
	return nil
}
