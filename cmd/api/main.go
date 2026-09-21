package main

import (
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

	// router.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
	// 		var responseTime int64

	// 		start := time.Now()
	// 		next.ServeHTTP(res, req)
	// 		responseTime = time.Since(start).Milliseconds()
	// 		log.Println(req.Method, req.URL.Path, responseTime, "ms")
	// 	})
	// })
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

	router.Post("/products", func(res http.ResponseWriter, req *http.Request) {
		var product Product
		render.DecodeJSON(req.Body, &product)
		product.ID = 1
		render.JSON(res, req, product)
	})

	http.ListenAndServe(":3000", router)
}
