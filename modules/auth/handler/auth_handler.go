package handler

import (
	"crm-sebagian-team/domain"
	"crm-sebagian-team/helpers"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthService domain.AuthService
}

func NewAuthHandler(e *echo.Group, r *echo.Group, as domain.AuthService) {
	handler := &AuthHandler{
		AuthService: as,
	}

	e.POST("/auth/login", handler.Login)
}

func isRequestValid(f *domain.LoginRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(f)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (hd *AuthHandler) Login(c echo.Context) error {
	credential := new(domain.LoginRequest)
	if err := c.Bind(credential); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(credential); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()

	resLogin, err := hd.AuthService.Login(ctx, credential)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	res := map[string]interface{}{
		"message": "Login success",
		"data":    resLogin,
	}

	return c.JSON(http.StatusOK, res)
}
