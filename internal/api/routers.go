package api

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/rodruizronald/inventory-service/internal/products"
)

// NewRouter initializes and returns a new router
func NewRouter(repo *products.Repository) *chi.Mux {
	r := chi.NewRouter()
	handler := products.NewProductHandler(repo)

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// Product Routes
	r.Route("/api/v1/products", func(r chi.Router) {
		r.Post("/", handler.CreateProduct)
		r.Get("/", handler.GetProducts)
		r.Get("/{id}", handler.GetProductByID)
		r.Put("/{id}", handler.UpdateProduct)
		r.Delete("/{id}", handler.DeleteProduct)
	})

	return r
}
