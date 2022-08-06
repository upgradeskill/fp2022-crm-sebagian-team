package handler

import (
	"net/http"

	"crm-sebagian-team/domain"
	"crm-sebagian-team/helpers"
	"crm-sebagian-team/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserSvc domain.UserService
}

func NewUserHandler(e *echo.Group, r *echo.Group, us domain.UserService) {
	handler := &UserHandler{
		UserSvc: us,
	}
	r.GET("/users", handler.GetAll)
	e.POST("/users/register", handler.Register)
}

func isRequestValid(u *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (h *UserHandler) GetAll(c echo.Context) error {

	cx := c.Request().Context()

	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"code": c.QueryParam("code"),
		"name": c.QueryParam("name"),
	}

	listUser, total, err := h.UserSvc.Get(cx, &params)

	if err != nil {
		return err
	}

	res := helpers.Paginate(c, user.NewListUserResponse(listUser), total, params)

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Register(c echo.Context) error {
	req := new(domain.User)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(req); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	usr, err := h.UserSvc.Store(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	resCreate := domain.UserResponse{
		ID:    usr.ID,
		Name:  usr.Name,
		Email: usr.Email,
	}

	res := map[string]interface{}{
		"message": "create account success",
		"data":    resCreate,
	}

	return c.JSON(http.StatusCreated, res)
}
