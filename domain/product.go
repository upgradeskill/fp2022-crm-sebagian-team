package domain

import (
	"context"
	"time"
)

type Product struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Qty       int64     `json:"qty"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	DeletedBy string    `json:"deleted_by,omitempty"`
}

type ProductRequest struct {
	Name      string    `json:"name" validate:"required"`
	Qty       int64     `json:"qty"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Qty       int64     `json:"qty"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}

type ProductService interface {
	Get(ctx context.Context, params *Request) ([]Product, int64, error)
	GetByName(ctx context.Context, name string) (Product, error)
	// Store(ctx context.Context, prd *Product) (Product, error)
}

type ProductRepository interface {
	Get(ctx context.Context, params *Request) ([]Product, int64, error)
	GetByName(ctx context.Context, name string) (Product, error)
	Delete(ctx context.Context, ID int64) (Product, error)
	// Store(ctx context.Context, prd *Product) (Product, error)
}
