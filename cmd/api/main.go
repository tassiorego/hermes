package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func main() {
	router := chi.NewRouter()
	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		res.Write([]byte("Hello, " + name + "!"))
	})

	router.Get("/{name}", func(res http.ResponseWriter, req *http.Request) {
		name := chi.URLParam(req, "name")
		res.Write([]byte("Hello, " + name + "!"))
	})

	router.Get("/v1/users", func(res http.ResponseWriter, req *http.Request) {
		users := []string{"Alice", "Bob", "Charlie"}
		response := map[string][]string{"users": users}
		render.JSON(res, req, response)
	})

	http.ListenAndServe(":3000", router)
}
