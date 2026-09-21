package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) GetCampaigns(res http.ResponseWriter, req *http.Request) {
	campaigns, err := h.CampaignService.Get()

	if err != nil {
		render.Status(req, http.StatusInternalServerError)
		render.JSON(res, req, map[string]string{"error": err.Error()})
		return
	}
	render.Status(req, http.StatusOK)
	render.JSON(res, req, campaigns)
}
