package dto

import (
	"time"

	"gorm.io/gorm"
)

type AddProductRequest struct {
	Name        string `binding:"required" json:"name"`
	Price       int64  `binding:"required" json:"price"`
	Description string `binding:"required" json:"description"`
	Quantity    int    `binding:"required" json:"quantity"`
}

type AddProductResponse struct {
	ID          uint           `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Name        string         `json:"name"`
	Price       int64          `json:"price"`
	Description string         `json:"description"`
	Quantity    int            `json:"quantity"`
}

type GetProductResponse struct {
	ID          uint           `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Name        string         `json:"name"`
	Price       int64          `json:"price"`
	Description string         `json:"description"`
	Quantity    int            `json:"quantity"`
}
