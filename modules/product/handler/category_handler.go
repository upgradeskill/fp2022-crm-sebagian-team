package handler

import (
	"net/http"

	"crm-sebagian-team/domain"
	"crm-sebagian-team/helpers"
	"crm-sebagian-team/modules/product"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	CategorySvc domain.CategoryService
}

func NewCategoryHandler(e *echo.Group, r *echo.Group, svc domain.CategoryService) {
	handler := &CategoryHandler{
		CategorySvc: svc,
	}
	r.GET("/category-product", handler.GetAll)
	r.POST("/category-product", handler.Store)
}

func isRequestValid(u *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (h *CategoryHandler) GetAll(c echo.Context) error {

	cx := c.Request().Context()

	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"name": c.QueryParam("name"),
	}

	listData, total, err := h.CategorySvc.Get(cx, &params)

	if err != nil {
		return err
	}

	res := helpers.Paginate(c, product.NewListCategoryResponse(listData), total, params)

	return c.JSON(http.StatusOK, res)
}

func (h *CategoryHandler) Store(c echo.Context) error {
	req := new(domain.CategoryRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	data, err := h.CategorySvc.Store(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	resCreate := domain.CategoryResponse{
		ID:       data.ID,
		Name:     data.Name,
		IsActive: data.IsActive,
	}

	res := map[string]interface{}{
		"message": "create category success",
		"data":    resCreate,
	}

	return c.JSON(http.StatusCreated, res)
}
