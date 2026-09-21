package endpoints

import (
	"hermes/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CreateCampaign(res http.ResponseWriter, req *http.Request) (response any, err error) {
	var newCampaign contract.CreateCampaignDTO
	render.Decode(req, &newCampaign)
	id, err := h.CampaignService.Create(newCampaign)
	if err != nil {
		return nil, err
	}

	return map[string]any{"id": id}, nil
}
