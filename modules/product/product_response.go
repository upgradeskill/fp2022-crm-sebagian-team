package product

import "crm-sebagian-team/domain"

func NewListProductResponse(datas []domain.Product) []domain.ProductResponse {
	resp := []domain.ProductResponse{}

	for _, data := range datas {
		resp = append(resp, NewProductResponse(data))
	}

	return resp
}

func NewProductResponse(product domain.Product) domain.ProductResponse {
	return domain.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Qty:       product.Qty,
		CreatedAt: product.CreatedAt,
		CreatedBy: product.CreatedBy,
	}
}
