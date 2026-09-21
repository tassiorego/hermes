package endpoints

import (
	"hermes/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CreateCampaign(res http.ResponseWriter, req *http.Request) {
	var newCampaign contract.CreateCampaignDTO
	render.Decode(req, &newCampaign)
	id, err := h.CampaignService.Create(newCampaign)

	if err != nil {
		render.Status(req, http.StatusBadRequest)
		render.JSON(res, req, map[string]interface{}{"error": err.Error()})
		return
	}
	render.Status(req, http.StatusCreated)
	render.JSON(res, req, map[string]interface{}{"id": id})
}
