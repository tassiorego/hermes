package database

import (
	"hermes/internal/domain/campaign"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (repo *CampaignRepository) Save(campaign *campaign.Campaign) error {
	repo.campaigns = append(repo.campaigns, *campaign)

	return nil
}

func (repo *CampaignRepository) Get() ([]campaign.Campaign, error) {
	if repo.campaigns == nil {
		return []campaign.Campaign{}, nil
	}
	return repo.campaigns, nil
}
