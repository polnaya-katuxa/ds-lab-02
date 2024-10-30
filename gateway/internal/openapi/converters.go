package openapi

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/polnaya-katuxa/ds-lab-02/gateway/internal/models"
)

func processError(c echo.Context, err error, comment string) error {
	var validationError models.ValidationError
	if errors.As(err, &validationError) {
		validationError.Message = fmt.Sprintf("%s: %s", comment, validationError.Message)

		return c.JSON(http.StatusBadRequest, validationError)
	}

	var internalError models.InternalError
	if errors.As(err, &internalError) {
		internalError.Message = fmt.Sprintf("%s: %s", comment, internalError.Message)

		return c.JSON(internalError.StatusCode, internalError)
	}

	internalError = models.InternalError{
		Message: fmt.Sprintf("%s: %s", comment, err.Error()),
	}
	return c.JSON(http.StatusInternalServerError, internalError)
}
