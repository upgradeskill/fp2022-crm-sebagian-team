package handler

import (
	"net/http"

	"crm-sebagian-team/domain"
	"crm-sebagian-team/helpers"
	"crm-sebagian-team/modules/product"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductSvc domain.ProductService
}

func NewProductHandler(e *echo.Group, r *echo.Group, prds domain.ProductService) {
	handler := &ProductHandler{
		ProductSvc: prds,
	}
	r.GET("/products", handler.GetAll)
	// r.POST("/product/create", handler.CreateNew)
}

func (h *ProductHandler) GetAll(c echo.Context) error {

	cx := c.Request().Context()

	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"name": c.QueryParam("name"),
	}

	listProduct, total, err := h.ProductSvc.Get(cx, &params)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := helpers.Paginate(c, product.NewListProductResponse(listProduct), total, params)

	return c.JSON(http.StatusOK, res)
}

// func (h *ProductHandler) CreateNew(c echo.Context) error {
// 	req := new(domain.ProductRequest)
// 	if err := c.Bind(&req); err != nil {
// 		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
// 	}

// 	usr, err := h.ProductSvc.Store(ctx, req)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
// 	}

// 	resCreate := domain.ProductResponse{
// 		ID:        usr.ID,
// 		Name:      usr.Name,
// 		Qty:       usr.Qty,
// 		CreatedAt: usr.CreatedAt,
// 		CreatedBy: usr.CreatedBy,
// 	}

// 	res := map[string]interface{}{
// 		"message": "create product success",
// 		"data":    resCreate,
// 	}

// 	return c.JSON(http.StatusCreated, res)
// }
