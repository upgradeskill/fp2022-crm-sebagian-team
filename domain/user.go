package domain

import (
	"context"
	"time"
)

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Address    string    `json:"address"`
	IdPosition uint8     `json:"id_position"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  string    `json:"updated_by"`
	DeletedAt  time.Time `json:"deleted_at,omitempty"`
	DeletedBy  string    `json:"deleted_by,omitempty"`
}

type UserRequest struct {
	Name       string    `json:"name" validate:"required"`
	Email      string    `json:"email" validate:"required"`
	Password   string    `json:"password" validate:"required"`
	Address    string    `json:"address"validate:"required"`
	IdPosition uint8     `json:"id_position" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Address    string    `json:"address"`
	IdPosition uint8     `json:"id_position"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
}

type UserService interface {
	Get(ctx context.Context, params *Request) ([]User, int64, error)
	Store(ctx context.Context, request *UserRequest) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
}

type UserRepository interface {
	Get(ctx context.Context, params *Request) ([]User, int64, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Store(ctx context.Context, usr *User) (User, error)
}
