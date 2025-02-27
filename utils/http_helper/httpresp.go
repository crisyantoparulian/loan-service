package httphelper

import (
	"net/http"

	apperror "github.com/crisyantoparulian/loansvc/utils/app_error"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Errors  *map[string]string `json:"errors,omitempty"`
	Message string             `json:"message"`
	Success bool               `json:"success"`
}

type ResponseSuccess struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

func HttpRespError(c echo.Context, srcErr error) (err error) {
	// Check if it's an APIError
	if apiErr, ok := srcErr.(*apperror.AppError); ok {
		err = c.JSON(apiErr.Code, ErrorResponse{
			Message: srcErr.Error(),
		})
	} else {
		err = c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: srcErr.Error(),
		})
	}
	return
}

func HttpSuccessOk(c echo.Context, msg string, data interface{}) (err error) {
	return c.JSON(http.StatusOK, ResponseSuccess{
		Success: true,
		Message: msg,
		Data:    data,
	})
}

func HttpSuccessCreated(c echo.Context, msg string, data interface{}) (err error) {
	return c.JSON(http.StatusCreated, ResponseSuccess{
		Success: true,
		Message: msg,
		Data:    data,
	})
}
