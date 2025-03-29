package products

import "time"

// Product represents an item in the inventory system
type Product struct {
	ID         int        `json:"id" db:"id"`                             // Unique identifier
	Name       string     `json:"name" db:"name"`                         // Product name
	Category   string     `json:"category" db:"category"`                 // Product category (e.g., Dairy, Meat, Grains)
	Quantity   int        `json:"quantity" db:"quantity"`                 // Available stock
	Unit       string     `json:"unit" db:"unit"`                         // Measurement unit (kg, liters, pieces)
	Price      float64    `json:"price" db:"price"`                       // Price per unit
	ExpiryDate *time.Time `json:"expiry_date,omitempty" db:"expiry_date"` // Expiration date (optional, for perishable goods)
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`             // When the product was added
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`             // Last update timestamp
}
