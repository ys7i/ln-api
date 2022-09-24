package echoutil

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrInternal(ec echo.Context, err error) error {
	ec.Logger().Error()
	return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
		"error":     "internal server error",
		"error_msg": err.Error(),
	})
}
