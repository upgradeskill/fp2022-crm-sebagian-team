package service

import (
	"crm-sebagian-team/domain"
	"time"
)

type productService struct {
	productRepo    domain.ProductRepository
	contextTimeout time.Duration
}
