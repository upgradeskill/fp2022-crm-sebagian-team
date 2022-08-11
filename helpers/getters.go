package helpers

import (
	"math"
	"net/http"
	"strconv"

	"crm-sebagian-team/domain"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type ResponseError struct {
	Message string `json:"message"`
}

func GetRequestParams(c echo.Context) domain.Request {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	perPage, _ := strconv.ParseInt(c.QueryParam("per_page"), 10, 64)

	if page == 0 {
		page = 1
	}

	if perPage == 0 {
		perPage = 10
	}

	offset := (page - 1) * perPage

	sortOrder := c.QueryParam("sort_order")
	if sortOrder == "" {
		sortOrder = "ASC"
	}

	params := domain.Request{
		Keyword:   c.QueryParam("q"),
		Page:      page,
		PerPage:   perPage,
		Offset:    offset,
		SortBy:    c.QueryParam("sort_by"),
		SortOrder: sortOrder,
		StartDate: c.QueryParam("start_date"),
		EndDate:   c.QueryParam("end_date"),
	}

	return params
}

func Paginate(c echo.Context, data interface{}, total int64, params domain.Request) *domain.ResultsData {
	return &domain.ResultsData{
		Data: data,
		Meta: &domain.MetaData{
			TotalCount:  total,
			TotalPage:   math.Ceil(float64(total) / float64(params.PerPage)),
			CurrentPage: params.Page,
			PerPage:     params.PerPage,
		},
	}
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrInvalidCredentials:
		return http.StatusUnauthorized
	case domain.ErrUserIsNotActive:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

// GetAuthenticatedUser ...
func GetAuthenticatedUser(c echo.Context) *domain.JwtCustomClaims {
	auth := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &auth)
	return &auth
}
