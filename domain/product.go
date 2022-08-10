package domain

import (
	"context"
)

type Product struct {
	ID		int64	`json:"id"`
	Name	string	`json:"name"`
	Qty		int32	`json:"qty"`
}

type ProductRequest struct {
	Name	string	`json:"name" validate:"required"`
	Qty		int32	`json:"email"`
}

type ProductResponse struct {
	ID		int64	`json:"id"`
	Name	string	`json:"name"`
	Qty		int32	`json:"qty"`
}

type ProductService interface {
	Get(ctx context.Context, params *Request) ([]Product, int32, error)
	Store(ctx context.Context, prd *Product) (Product, error)
}

type ProductRepository interface {
	Get(ctx context.Context, params *Request) ([]Product, int64, error)
	GetByName(ctx context.Context, name string) (Product, error)
	Store(ctx context.Context, prd *Product) (Product, error)
}
