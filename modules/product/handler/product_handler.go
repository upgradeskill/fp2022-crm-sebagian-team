package handler

import (
	"net/http"

	"crm-sebagian-team/domain"
	"crm-sebagian-team/helpers"
	"crm-sebagian-team/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductSvc domain.ProductService
}