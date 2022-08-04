package domain

import (
	"context"
)

type User struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"email"`
}

type UserRequest struct {
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UserResponse struct {
	ID          int64  `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"email"`
}

type UserService interface {
	Get(ctx context.Context, params *Request) ([]User, int64, error)
}

type UserRepository interface {
	Get(ctx context.Context, params *Request) (new []User, total int64, err error)
}
