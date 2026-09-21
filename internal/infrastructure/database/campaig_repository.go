package database

import "hermes/internal/domain/campaign"

type CampaignRepository struct {
	campaigns []*campaign.Campaign
}

func (repo *CampaignRepository) Save(campaign *campaign.Campaign) error {
	repo.campaigns = append(repo.campaigns, campaign)
	return nil
}
