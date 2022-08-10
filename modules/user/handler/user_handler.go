package handler

import (
	"fmt"
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
	r.POST("/users/create", handler.Register)
	e.POST("/users/register", handler.Register)
}

func isRequestValid(u *domain.UserRequest) (bool, error) {
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
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := helpers.Paginate(c, user.NewListUserResponse(listUser), total, params)

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Register(c echo.Context) error {
	req := new(domain.UserRequest)
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	fmt.Println(helpers.GetAuthenticatedUser(c).Email)
	if helpers.GetAuthenticatedUser(c).Email == "" {
		req.CreatedBy = req.Email
	} else {
		req.CreatedBy = helpers.GetAuthenticatedUser(c).Email
	}

	if ok, err := isRequestValid(req); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()

	email, _ := h.UserSvc.GetByEmail(ctx, req.Email)
	if req.Email == email.Email {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s Sudah terdaftar", req.Email))
	}

	usr, err := h.UserSvc.Store(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	resCreate := domain.UserResponse{
		ID:         usr.ID,
		Name:       usr.Name,
		Email:      usr.Email,
		Password:   usr.Password,
		Address:    usr.Address,
		CreatedAt:  usr.CreatedAt,
		IdPosition: usr.IdPosition,
		CreatedBy:  usr.CreatedBy,
	}

	res := map[string]interface{}{
		"message": "create account success",
		"data":    resCreate,
	}

	return c.JSON(http.StatusCreated, res)
}
