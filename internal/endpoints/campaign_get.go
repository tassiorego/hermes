package endpoints

import (
	"net/http"
)

func (h *Handler) GetCampaigns(res http.ResponseWriter, req *http.Request) (response any, err error) {
	campaigns, err := h.CampaignService.Get()

	if err != nil {
		return nil, err
	}
	return campaigns, nil
}
