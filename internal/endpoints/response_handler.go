package endpoints

import (
	"net/http"

	"github.com/go-chi/render"

	internalerrors "hermes/internal/internal-errors"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (response any, err error)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ResponseHandler(successStatus int, endpoint EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := endpoint(w, r)
		if err != nil {
			status := internalerrors.StatusFor(err)
			render.Status(r, status)
			render.JSON(w, r, ErrorResponse{
				Status:  status,
				Message: err.Error(),
			})
			return
		}
		render.Status(r, successStatus)
		if response != nil {
			render.JSON(w, r, response)
		}
	})
}
