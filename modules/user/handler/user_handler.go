package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"crm-sebagian-team/domain"
	"crm-sebagian-team/helpers"
	"crm-sebagian-team/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserSvc     domain.UserService
	PositionSvc domain.PositionService
}

func NewUserHandler(e *echo.Group, r *echo.Group, us domain.UserService, ps domain.PositionService) {
	handler := &UserHandler{
		UserSvc:     us,
		PositionSvc: ps,
	}
	r.GET("/users", handler.GetAll)
	r.POST("/users/create", handler.Register)
	r.PUT("/users/update/:id", handler.Update)
	r.DELETE("/users/delete/:id", handler.Delete)
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

func isRequestValidUpdate(u *domain.UserUpdate) (bool, error) {
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

	_, errPosition := h.PositionSvc.GetByID(ctx, req.IdPosition)
	if errPosition != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("ID Position %d %s", req.IdPosition, errPosition.Error()))
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
		CreatedAt:  &usr.CreatedAt,
		IdPosition: usr.IdPosition,
		CreatedBy:  usr.CreatedBy,
	}

	res := map[string]interface{}{
		"message": "create account success",
		"data":    resCreate,
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *UserHandler) Update(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ctx := c.Request().Context()
	_, errNotFound := h.UserSvc.GetByID(ctx, id)
	if errNotFound != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("ID User %d %s", id, errNotFound.Error()))
	}
	req := domain.UserUpdate{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	if ok, err := isRequestValidUpdate(&req); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req.ID = id
	req.UpdatedBy = helpers.GetAuthenticatedUser(c).Email

	_, errPosition := h.PositionSvc.GetByID(ctx, req.IdPosition)
	if errPosition != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("ID Position %d %s", req.IdPosition, errPosition.Error()))
	}

	usr, err := h.UserSvc.UpdateUser(ctx, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	resUpdate := domain.UserResponse{
		ID:         usr.ID,
		Name:       usr.Name,
		Email:      usr.Email,
		Password:   usr.Password,
		Address:    usr.Address,
		IdPosition: usr.IdPosition,
		UpdatedAt:  usr.UpdatedAt,
		UpdatedBy:  usr.UpdatedBy,
	}

	res := map[string]interface{}{
		"message": "update account success",
		"data":    resUpdate,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Delete(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ctx := c.Request().Context()
	usr, errNotFound := h.UserSvc.GetByID(ctx, id)
	if errNotFound != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("ID User %d %s", id, errNotFound.Error()))
	}
	deletedBy := helpers.GetAuthenticatedUser(c).Email

	err := h.UserSvc.DeleteUser(ctx, id, deletedBy)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	res := map[string]string{
		"message": fmt.Sprintf("%s success deleted", usr.Email),
	}

	return c.JSON(http.StatusOK, res)
}
