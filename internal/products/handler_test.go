package products

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	tests := []struct {
		name         string
		input        any
		setMock      func(m *MockRepository, input any)
		expectStatus int
	}{
		{
			name:  "Valid product creation",
			input: Product{Name: "Milk", Category: "Dairy", Unit: "Liter", Price: 2.5, Quantity: 10},
			setMock: func(m *MockRepository, input any) {
				p, _ := input.(Product)
				m.On("CreateProduct", mock.Anything, p.Name, p.Category, p.Unit,
					p.Quantity, p.Price, p.ExpiryDate).Return(1, nil).Once()
			},
			expectStatus: http.StatusCreated,
		},
		{
			name:  "Database failure - missing product name",
			input: Product{Category: "Dairy", Unit: "Liter", Price: 2.5, Quantity: 10},
			setMock: func(m *MockRepository, input any) {
				p, _ := input.(Product)
				m.On("CreateProduct", mock.Anything, p.Name, p.Category, p.Unit,
					p.Quantity, p.Price, p.ExpiryDate).Return(0, errors.New("invalid input")).Once()
			},
			expectStatus: http.StatusInternalServerError,
		},
		{
			name:         "Not a valid Prodcut JSON",
			input:        struct{ field string }{""},
			setMock:      func(m *MockRepository, input any) {},
			expectStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			h := NewProductHandler(mockRepo)
			tt.setMock(mockRepo, tt.input)

			reqBody, err := json.Marshal(tt.input)
			require.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			h.CreateProduct(w, req)

			resp := w.Result()
			assert.Equal(t, tt.expectStatus, resp.StatusCode)
			mockRepo.AssertExpectations(t)
		})
	}
}
