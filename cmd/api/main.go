package main

import (
	"hermes/internal/domain/campaign"
	"hermes/internal/endpoints"
	"hermes/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	handler := &endpoints.Handler{
		CampaignService: campaignService,
	}

	router.Post("/campaigns", handler.CreateCampaign)
	router.Get("/campaigns", handler.GetCampaigns)

	http.ListenAndServe(":3000", router)
}
