package service

import (
	"context"
	"time"

	"crm-sebagian-team/domain"

	"golang.org/x/crypto/bcrypt"
)

type productService struct {
	productRepo			domain.ProductRepository
	contextTimeout 		time.Duration
}
