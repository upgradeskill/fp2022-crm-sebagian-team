package repository

import (
	"context"

	"crm-sebagian-team/domain"
	"crm-sebagian-team/helpers"
	"crm-sebagian-team/modules/product"

	"gorm.io/gorm"
)

type categoryRepository struct {
	Conn *gorm.DB
}

func NewCategoryRepository(Conn *gorm.DB) domain.CategoryRepository {
	return &categoryRepository{Conn}
}

func (rp *categoryRepository) Get(ctx context.Context, params *domain.Request) (res []domain.Category, total int64, err error) {
	var totalRows int64

	resCategory := []domain.Category{}
	category := []product.Category{}

	query := rp.Conn.Model(&category).Debug()

	query.Scopes(helpers.PaginateQuery(params)).Find(&category).Count(&totalRows)
	if query.Error != nil {
		return []domain.Category{}, 0, query.Error
	}

	for _, data := range category {
		resCategory = append(resCategory, domain.Category{
			ID:       int64(data.ID),
			Name:     data.Name,
			IsActive: data.IsActive,
		})
	}

	return resCategory, totalRows, nil
}

func (rp *categoryRepository) Store(ctx context.Context, dom *domain.Category) (domain.Category, error) {
	catEntity := product.Category{
		Name:     dom.Name,
		IsActive: dom.IsActive,
	}

	query := rp.Conn.Create(&catEntity)

	if query.Error != nil {
		return domain.Category{}, query.Error
	}

	dom.ID = int64(catEntity.ID)

	return *dom, nil
}
