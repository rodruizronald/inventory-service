package products

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Database interface to support pgxpool and mocks
type Database interface {
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
}

// Repository manages database operations for products
type Repository struct {
	db Database
}

// NewRepository initializes a new repository
func NewRepository(db Database) *Repository {
	return &Repository{db: db}
}

// 🔹 Global SQL Queries
const (
	queryCreateProduct = `
		INSERT INTO products (name, category, quantity, unit, price, expiry_date, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW()) RETURNING id`

	queryGetProductByID = `
		SELECT id, name, category, quantity, unit, price, expiry_date, created_at, updated_at 
		FROM products WHERE id = $1`

	queryGetProducts = `
		SELECT id, name, category, quantity, unit, price, expiry_date, created_at, updated_at 
		FROM products`

	queryUpdateProduct = `
		UPDATE products 
		SET name = $1, category = $2, quantity = $3, unit = $4, price = $5, expiry_date = $6, updated_at = NOW() 
		WHERE id = $7`

	queryDeleteProduct = `
		DELETE FROM products WHERE id = $1`
)

// CreateProduct inserts a new product into the inventory and returns its ID
func (r *Repository) CreateProduct(ctx context.Context, name, category, unit string, quantity int, price float64, expiryDate *time.Time) (int, error) {
	var id int
	row := r.db.QueryRow(ctx, queryCreateProduct, name, category, quantity, unit, price, expiryDate)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create product: %w", err)
	}
	return id, nil
}

// GetProductByID retrieves a product by its ID
func (r *Repository) GetProductByID(ctx context.Context, id int) (*Product, error) {
	var product Product
	row := r.db.QueryRow(ctx, queryGetProductByID, id)
	err := row.Scan(&product.ID, &product.Name, &product.Category, &product.Quantity, &product.Unit, &product.Price, &product.ExpiryDate, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}
	return &product, nil
}

// GetProducts retrieves all products in the inventory
func (r *Repository) GetProducts(ctx context.Context) ([]Product, error) {
	rows, err := r.db.Query(ctx, queryGetProducts)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Quantity, &product.Unit, &product.Price, &product.ExpiryDate, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}

// UpdateProduct updates a product's details
func (r *Repository) UpdateProduct(ctx context.Context, id int, name, category, unit string, quantity int, price float64, expiryDate *time.Time) error {
	_, err := r.db.Exec(ctx, queryUpdateProduct, name, category, quantity, unit, price, expiryDate, id)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}
	return nil
}

// DeleteProduct removes a product from the inventory by ID
func (r *Repository) DeleteProduct(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, queryDeleteProduct, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}
