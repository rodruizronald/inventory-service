package products

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMockCreateProduct(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := context.Background()
	mockRepo.On("CreateProduct", ctx, "Milk", "Dairy", "liters", 10, 2.5, mock.Anything).Return(1, nil)

	id, err := mockRepo.CreateProduct(ctx, "Milk", "Dairy", "liters", 10, 2.5, nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)

	mockRepo.AssertExpectations(t)
}

func TestMockGetProductByID(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := context.Background()
	product := &Product{ID: 1, Name: "Milk", Category: "Dairy", Quantity: 10}
	mockRepo.On("GetProductByID", ctx, 1).Return(product, nil)

	result, err := mockRepo.GetProductByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, product, result)

	mockRepo.AssertExpectations(t)
}

func TestMockGetProducts(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := context.Background()
	mockRepo.On("GetProducts", ctx).Return([]Product{{ID: 1, Name: "Milk"}}, nil)

	result, err := mockRepo.GetProducts(ctx)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Milk", result[0].Name)

	mockRepo.AssertExpectations(t)
}

func TestMockUpdateProduct(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := context.Background()
	mockRepo.On("UpdateProduct", ctx, 1, "Milk", "Dairy", "liters", 10, 2.5, mock.Anything).Return(nil)

	err := mockRepo.UpdateProduct(ctx, 1, "Milk", "Dairy", "liters", 10, 2.5, nil)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestMockDeleteProduct(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := context.Background()
	mockRepo.On("DeleteProduct", ctx, 1).Return(nil)

	err := mockRepo.DeleteProduct(ctx, 1)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
