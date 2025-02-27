package handler

import (
	"fmt"
	"net/http"
	"time"

	apperror "github.com/crisyantoparulian/loansvc/utils/app_error"
	httphelper "github.com/crisyantoparulian/loansvc/utils/http_helper"
	"github.com/crisyantoparulian/loansvc/utils/str"
	utilvalidator "github.com/crisyantoparulian/loansvc/utils/validator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func generateLoanCode(borrowerType string) string {
	code := "RGL" // default regular
	switch borrowerType {
	case "premium":
		code = "PRM"
	case "ultimate":
		code = "ULT"
	}
	datePart := time.Now().Format("060102")   // YYMMDD format
	randomPart := str.GenerateRandomString(5) // Generate 5-character random string
	return fmt.Sprintf("%s%s%s", code, datePart, randomPart)
}

func (s *Server) bindAndValidatePayload(c echo.Context, payload interface{}) error {
	if err := c.Bind(&payload); err != nil {
		httphelper.HttpRespError(c,
			apperror.WrapWithCode(fmt.Errorf("failed to unmarshall request: %w", err),
				http.StatusBadRequest))
		return err
	}

	// Validate payload
	if err := s.Validator.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		validationErr := utilvalidator.FormatValidationErrors(validationErrors)
		c.JSON(http.StatusBadRequest, httphelper.ErrorResponse{
			Message: "Validation failed",
			Errors:  &validationErr,
		})
		return err
	}

	return nil
}
