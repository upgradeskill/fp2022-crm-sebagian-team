package server

import (
	"crm-sebagian-team/config"
	"crm-sebagian-team/middleware"
	_authHandler "crm-sebagian-team/modules/auth/handler"
	_productHandler "crm-sebagian-team/modules/product/handler"
	_userHandler "crm-sebagian-team/modules/user/handler"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func newAppHandler(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"app": "Bootcamp Apps",
		})
	})
}

// NewHandler will create a new handler for the given usecase
func NewHandler(cfg *config.Config, svc *Service) {

	e := echo.New()
	e.HTTPErrorHandler = ErrorHandler
	middleware := middleware.InitMiddleware(cfg)

	v1 := e.Group("/v1")
	route := e.Group("")
	route.Use(middleware.JWT)

	newAppHandler(e)
	_authHandler.NewAuthHandler(v1, route, svc.AuthService)
	_userHandler.NewUserHandler(v1, route, svc.UserService, svc.PositionService)
	_productHandler.NewCategoryHandler(v1, route, svc.CategoryService)

	log.Fatal(e.Start(":3000"))
}

// ErrorHandler ...
func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Logger().Error(report)
	c.JSON(report.Code, report)
}
