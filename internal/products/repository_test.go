package products

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestRepository_CreateProduct tests the CreateProduct method
func TestRepository_CreateProduct(t *testing.T) {
	mockDB, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mockDB.Close()

	repo := NewRepository(mockDB)

	expectedID := 1
	name := "Milk"
	category := "Dairy"
	quantity := 100
	unit := "liters"
	price := 2.50
	expiryDate := time.Now().AddDate(0, 1, 0)

	query := regexp.QuoteMeta(queryCreateProduct)

	mockDB.ExpectQuery(query).
		WithArgs(name, category, quantity, unit, price, &expiryDate).
		WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(expectedID))

	id, err := repo.CreateProduct(context.Background(), name, category, unit, quantity, price, &expiryDate)

	assert.NoError(t, err)
	assert.Equal(t, expectedID, id)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

// TestRepository_GetProductByID tests GetProductByID method
func TestRepository_GetProductByID(t *testing.T) {
	mockDB, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mockDB.Close()

	repo := NewRepository(mockDB)

	expectedProduct := &Product{
		ID:         1,
		Name:       "Milk",
		Category:   "Dairy",
		Quantity:   100,
		Unit:       "liters",
		Price:      2.50,
		ExpiryDate: nil,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	query := regexp.QuoteMeta(queryGetProductByID)

	mockDB.ExpectQuery(query).
		WithArgs(expectedProduct.ID).
		WillReturnRows(pgxmock.NewRows([]string{"id", "name", "category", "quantity", "unit", "price", "expiry_date", "created_at", "updated_at"}).
			AddRow(expectedProduct.ID, expectedProduct.Name, expectedProduct.Category, expectedProduct.Quantity, expectedProduct.Unit, expectedProduct.Price, expectedProduct.ExpiryDate, expectedProduct.CreatedAt, expectedProduct.UpdatedAt))

	product, err := repo.GetProductByID(context.Background(), expectedProduct.ID)

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

// TestRepository_GetProducts tests GetProducts method
func TestRepository_GetProducts(t *testing.T) {
	mockDB, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mockDB.Close()

	repo := NewRepository(mockDB)

	query := regexp.QuoteMeta(queryGetProducts)

	mockDB.ExpectQuery(query).
		WillReturnRows(pgxmock.NewRows([]string{"id", "name", "category", "quantity", "unit", "price", "expiry_date", "created_at", "updated_at"}).
			AddRow(1, "Milk", "Dairy", 100, "liters", 2.50, nil, time.Now(), time.Now()).
			AddRow(2, "Cheese", "Dairy", 50, "kg", 5.00, nil, time.Now(), time.Now()))

	products, err := repo.GetProducts(context.Background())

	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

// TestRepository_UpdateProduct tests UpdateProduct method
func TestRepository_UpdateProduct(t *testing.T) {
	mockDB, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mockDB.Close()

	repo := NewRepository(mockDB)

	id := 1
	newName := "Yogurt"
	newCategory := "Dairy"
	newQuantity := 80
	newUnit := "liters"
	newPrice := 3.00
	newExpiryDate := time.Now().AddDate(0, 2, 0)

	query := regexp.QuoteMeta(queryUpdateProduct)

	mockDB.ExpectExec(query).
		WithArgs(newName, newCategory, newQuantity, newUnit, newPrice, &newExpiryDate, id).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = repo.UpdateProduct(context.Background(), id, newName, newCategory, newUnit, newQuantity, newPrice, &newExpiryDate)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

// TestRepository_DeleteProduct tests DeleteProduct method
func TestRepository_DeleteProduct(t *testing.T) {
	mockDB, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mockDB.Close()

	repo := NewRepository(mockDB)

	id := 1
	query := regexp.QuoteMeta(queryDeleteProduct)

	mockDB.ExpectExec(query).
		WithArgs(id).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = repo.DeleteProduct(context.Background(), id)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

// TestRepository_GetProductByID_NotFound tests GetProductByID for a non-existing product
func TestRepository_GetProductByID_NotFound(t *testing.T) {
	mockDB, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mockDB.Close()

	repo := NewRepository(mockDB)

	id := 999 // Non-existent product ID

	query := regexp.QuoteMeta(queryGetProductByID)

	mockDB.ExpectQuery(query).
		WithArgs(id).
		WillReturnError(errors.New("no rows in result set"))

	product, err := repo.GetProductByID(context.Background(), id)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}
