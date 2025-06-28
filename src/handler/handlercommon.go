package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiResponse struct {
	Message string `json:"message"`
}

// TODO:Do not send internal err msg instead send correct error code and error msg
func httpErrResp(e error, c echo.Context) (err error) {
	var errResp ApiResponse
	errResp.Message = e.Error()
	return c.JSON(http.StatusBadRequest, errResp)
}
