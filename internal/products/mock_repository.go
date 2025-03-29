package products

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

// MockRepository is a mock implementation of the Repository for unit testing
type MockRepository struct {
	mock.Mock
}

// CreateProduct mocks the CreateProduct method
func (m *MockRepository) CreateProduct(ctx context.Context, name, category, unit string, quantity int, price float64, expiryDate *time.Time) (int, error) {
	args := m.Called(ctx, name, category, unit, quantity, price, expiryDate)
	return args.Int(0), args.Error(1)
}

// GetProductByID mocks the GetProductByID method
func (m *MockRepository) GetProductByID(ctx context.Context, id int) (*Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Product), args.Error(1)
}

// GetProducts mocks the GetProducts method
func (m *MockRepository) GetProducts(ctx context.Context) ([]Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]Product), args.Error(1)
}

// UpdateProduct mocks the UpdateProduct method
func (m *MockRepository) UpdateProduct(ctx context.Context, id int, name, category, unit string, quantity int, price float64, expiryDate *time.Time) error {
	args := m.Called(ctx, id, name, category, unit, quantity, price, expiryDate)
	return args.Error(0)
}

// DeleteProduct mocks the DeleteProduct method
func (m *MockRepository) DeleteProduct(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
