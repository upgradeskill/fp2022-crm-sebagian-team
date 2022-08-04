package handlers

import (
	"fp2022-crm-sebagian-team/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Result struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetUsers(ctx echo.Context) error {
	books, err := models.All()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, Result{Message: "Success Get Data", Data: users})
}

func DeleteUser(ctx echo.Context) error {
	id := ctx.Param("id")
	err := models.GetBookByIsbn(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, Result{Message: err.Error(), Data: nil})
	}
	errorDelete := models.Delete(id)
	if errorDelete != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, Result{Message: "Success Deleted Data", Data: nil})
}