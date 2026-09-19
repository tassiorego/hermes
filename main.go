package main

import (
	"hermes/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {
	campaign := &campaign.Campaign{}
	validate := validator.New()
	err := validate.Struct(campaign)
	if err != nil {
		// Handle validation errors
		println("Validation errors:", err.Error())
	}

}
