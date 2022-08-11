package repository

import (
	"context"
	"crm-sebagian-team/domain"
	"crm-sebagian-team/helpers"
	"crm-sebagian-team/modules/product"

	"gorm.io/gorm"
)

type productRepository struct {
	Conn *gorm.DB
}

func NewProductRepository(Conn *gorm.DB) domain.ProductRepository {
	return &productRepository{Conn}
}

func (m *productRepository) Get(ctx context.Context, params *domain.Request) (res []domain.Product, total int64, err error) {
	var totalRows int64

	resProduct := []domain.Product{}
	bk := []product.Product{}

	query := m.Conn.Model(&bk).Debug()

	query.Scopes(helpers.PaginateQuery(params)).Find(&bk).Count(&totalRows)
	if query.Error != nil {
		return []domain.Product{}, 0, query.Error
	}

	for _, data := range bk {
		resProduct = append(resProduct, domain.Product{
			ID:        int64(data.ID),
			Name:      data.Name,
			Qty:       int64(data.Qty),
			CreatedAt: data.CreatedAt,
			CreatedBy: data.CreatedBy,
		})
	}

	return resProduct, totalRows, nil
}

func (m *productRepository) GetByName(ctx context.Context, name string) (domain.Product, error) {
	prd := product.Product{}

	query := m.Conn.Model(&prd).Where("name = ?", name).Find(&prd)

	if query.Error != nil {
		return domain.Product{}, query.Error
	}

	resProduct := domain.Product{
		ID:   int64(prd.ID),
		Name: prd.Name,
		Qty:  int64(prd.Qty),
	}

	return resProduct, nil
}

// func (m *productRepository) Store(ctx context.Context, prd *domain.Product) (domain.Product, error) {
// 	prdEntity := product.Product{
// 		Name:      prd.Name,
// 		Qty:       int32(prd.Qty),
// 		CreatedAt: prd.CreatedAt,
// 		CreatedBy: prd.CreatedBy,
// 		UpdatedAt: prd.UpdatedAt,
// 	}

// 	query := m.Conn.Create(&prdEntity)

// 	if query.Error != nil {
// 		return domain.Product{}, query.Error
// 	}

// 	prd.ID = int64(prdEntity.ID)

// 	return *prd, nil
// }
