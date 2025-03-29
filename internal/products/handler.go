package products

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, name, category, unit string, quantity int, price float64, expiryDate *time.Time) (int, error)
	GetProductByID(ctx context.Context, id int) (*Product, error)
	GetProducts(ctx context.Context) ([]Product, error)
	UpdateProduct(ctx context.Context, id int, name, category, unit string, quantity int, price float64, expiryDate *time.Time) error
	DeleteProduct(ctx context.Context, id int) error
}

// ProductHandler provides HTTP handlers for product operations
type ProductHandler struct {
	Repo ProductRepository
}

// NewProductHandler initializes a new ProductHandler
func NewProductHandler(repo ProductRepository) *ProductHandler {
	return &ProductHandler{Repo: repo}
}

// CreateProduct handles product creation
// @Summary Create a new product
// @Description Add a new product to the inventory
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body Product true "Product data"
// @Success 201 {object} Product
// @Failure 400 {object} map[string]string
// @Router /products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.Repo.CreateProduct(r.Context(), p.Name, p.Category, p.Unit, p.Quantity, p.Price, p.ExpiryDate)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	p.ID = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

// GetProductByID handles retrieving a product by ID
// @Summary Get a product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} Product
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	product, err := h.Repo.GetProductByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// GetProducts handles retrieving all products
// @Summary Get all products
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {array} Product
// @Router /products [get]
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Repo.GetProducts(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

// UpdateProduct handles updating a product
// @Summary Update a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param product body Product true "Updated product data"
// @Success 200 {object} Product
// @Failure 400 {object} map[string]string
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var updated Product
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.Repo.UpdateProduct(r.Context(), id, updated.Name, updated.Category, updated.Unit, updated.Quantity, updated.Price, updated.ExpiryDate)
	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	updated.ID = id
	json.NewEncoder(w).Encode(updated)
}

// DeleteProduct handles deleting a product
// @Summary Delete a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.Repo.DeleteProduct(r.Context(), id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
