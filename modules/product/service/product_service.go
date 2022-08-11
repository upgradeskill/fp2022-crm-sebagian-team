package service

import (
	"context"
	"crm-sebagian-team/domain"
	"time"
)

type productService struct {
	productRepo    domain.ProductRepository
	contextTimeout time.Duration
}

func NewProductService(br domain.ProductRepository, timeout time.Duration) domain.ProductService {
	return &productService{
		productRepo:    br,
		contextTimeout: timeout,
	}
}

func (svc *productService) Get(c context.Context, params *domain.Request) ([]domain.Product, int64, error) {

	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()

	res, total, err := svc.productRepo.Get(ctx, params)
	if err != nil {
		return []domain.Product{}, 0, err
	}

	return res, total, nil
}

func (svc *productService) GetByName(c context.Context, name string) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()

	res, err := svc.productRepo.GetByName(ctx, name)
	if err != nil {
		return domain.Product{}, err
	}

	return res, nil
}

// func (svc *productService) Store(c context.Context, request *domain.ProductRequest) (domain.Product, error) {
// 	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
// 	defer cancel()

// 	prd := domain.Product{
// 		Name:      request.Name,
// 		Qty:       request.Qty,
// 		CreatedAt: request.CreatedAt,
// 		CreatedBy: request.CreatedBy,
// 		UpdatedAt: request.UpdatedAt,
// 	}
// 	usrRes, err := svc.ProductRepo.Store(ctx, &prd)

// 	return usrRes, nil
// }
