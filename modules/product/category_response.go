package product

import "crm-sebagian-team/domain"

func NewListCategoryResponse(datas []domain.Category) []domain.CategoryResponse {
	resp := []domain.CategoryResponse{}

	for _, data := range datas {
		resp = append(resp, NewCategoryResponse(data))
	}

	return resp
}

func NewCategoryResponse(category domain.Category) domain.CategoryResponse {
	return domain.CategoryResponse{
		ID:       category.ID,
		Name:     category.Name,
		IsActive: category.IsActive,
	}
}
