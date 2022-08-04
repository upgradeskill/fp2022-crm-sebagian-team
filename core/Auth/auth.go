func (hand *UserHandler) UserDelete(c echo.Context) error {
	//claims from jwtToken to UserClaims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*dto.UserClaims)
	resp := make(map[string]interface{})

	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		resp["message"] = "invalid id"
		return c.JSON(http.StatusBadRequest, resp)
	}

	data, validate, err := hand.userUseCase.UserDelete(uint(id), claims)

	if validate != nil {
		resp["message"] = "invalid parameters"
		resp["error_validation"] = validate
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err != nil {
		if err.Error() == "you don't have right" {
			resp["message"] = err.Error()
			return c.JSON(http.StatusUnauthorized, resp)
		}
		resp["message"] = err.Error()
		return c.JSON(http.StatusNotFound, resp)
	}

	resp["message"] = "Success Delete Data"
	resp["data"] = data
	return c.JSON(http.StatusOK, resp)
}