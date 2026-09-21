package endpoints

import (
	"hermes/internal/domain/campaign"
)

type Handler struct {
	CampaignService campaign.Service
}
