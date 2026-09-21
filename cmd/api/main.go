package main

import (
	"hermes/internal/contract"
	"hermes/internal/domain/campaign"
	"hermes/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
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

	service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	router.Post("/campaigns", func(res http.ResponseWriter, req *http.Request) {
		var newCampaign contract.CreateCampaignDTO
		render.Decode(req, &newCampaign)
		id, err := service.Create(newCampaign)

		if err != nil {
			render.Status(req, http.StatusBadRequest)
			render.JSON(res, req, map[string]interface{}{"error": err.Error()})
			return
		}
		render.Status(req, http.StatusCreated)
		render.JSON(res, req, map[string]interface{}{"id": id})
	})

	http.ListenAndServe(":3000", router)
}
