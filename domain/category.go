package domain

import (
	"context"
)

type Category struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	Timestamp
}

type CategoryRequest struct {
	Name     string `json:"name" validate:"required"`
	IsActive bool   `json:"is_active" validate:"required"`
}

type CategoryResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	Timestamp
}

type CategoryService interface {
	Get(ctx context.Context, params *Request) ([]Category, int64, error)
	Store(ctx context.Context, category *CategoryRequest) (CategoryResponse, error)
	// Update(ctx context.Context, usr *Category) (Category, error)
	// Delete(ctx context.Context, usr *Category) (Category, error)
}

type CategoryRepository interface {
	Get(ctx context.Context, params *Request) ([]Category, int64, error)
	// GetByID(ctx context.Context, email string) (Category, error)
	Store(ctx context.Context, category *Category) (Category, error)
	// Update(ctx context.Context, email string) (Category, error)
	// Delete(ctx context.Context, email string) (Category, error)
}
