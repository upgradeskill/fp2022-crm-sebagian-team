package domain

import (
	"context"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserService interface {
	Get(ctx context.Context, params *Request) ([]User, int64, error)
	Store(ctx context.Context, usr *User) (User, error)
}

type UserRepository interface {
	Get(ctx context.Context, params *Request) ([]User, int64, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Store(ctx context.Context, usr *User) (User, error)
}
