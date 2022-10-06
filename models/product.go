package models

import "time"

type Product struct {
	ID          int       `gorm:"primary_key" json:"id"`
	ProductName string    `json:"product_name"`
	Price       int       `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
