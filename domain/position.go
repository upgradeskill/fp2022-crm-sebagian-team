package domain

import "context"

type Position struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PositionResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Timestamp
}

type PositionService interface {
	GetByID(ctx context.Context, id int64) (Position, error)
}

type PositionRepository interface {
	GetByID(ctx context.Context, id int64) (Position, error)
}
