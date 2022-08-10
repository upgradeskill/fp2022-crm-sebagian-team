package service

import (
	"context"
	"net/http"
	"time"

	"crm-sebagian-team/domain"
	"crm-sebagian-team/modules/product"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type categoryService struct {
	categoryRepo   domain.CategoryRepository
	contextTimeout time.Duration
	validate       *validator.Validate
}

func NewCategoryService(rp domain.CategoryRepository, validate *validator.Validate, timeout time.Duration) domain.CategoryService {
	return &categoryService{
		categoryRepo:   rp,
		contextTimeout: timeout,
		validate:       validate,
	}
}

func (svc *categoryService) Get(c context.Context, params *domain.Request) ([]domain.Category, int64, error) {

	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()

	res, total, err := svc.categoryRepo.Get(ctx, params)
	if err != nil {
		return []domain.Category{}, 0, err
	}

	return res, total, nil
}

func (svc *categoryService) Store(c context.Context, req *domain.CategoryRequest) (domain.CategoryResponse, error) {
	ctx, cancel := context.WithTimeout(c, svc.contextTimeout)
	defer cancel()

	err := svc.validate.Struct(req)
	if err != nil {
		return domain.CategoryResponse{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	categoryData := domain.Category{
		Name:     req.Name,
		IsActive: req.IsActive,
	}

	categoryRes, err := svc.categoryRepo.Store(ctx, &categoryData)
	if err != nil {
		return domain.CategoryResponse{}, err
	}

	return product.NewCategoryResponse(categoryRes), nil
}
