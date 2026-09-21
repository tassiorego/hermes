package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) GetCampaigns(res http.ResponseWriter, req *http.Request) {
	campaigns := h.CampaignService.Get()
	render.Status(req, http.StatusOK)
	render.JSON(res, req, campaigns)
}
