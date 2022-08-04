package handler

import (
	"net/http"

	"crm-sebagian-team/domain"
	"crm-sebagian-team/helpers"
	"crm-sebagian-team/modules/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserSvc domain.UserService
}

func NewUserHandler(e *echo.Group, bs domain.UserService) {
	handler := &UserHandler{
		UserSvc: bs,
	}
	e.GET("/users", handler.GetAll)
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
