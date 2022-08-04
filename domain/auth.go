package domain

import (
	"context"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type AuthService interface {
	Login(ctx context.Context, req *LoginRequest) (LoginResponse, error)
}
