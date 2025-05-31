package middlewares

import (
	"errors"
	app_errors "github.com/keanutaufan/anitrackr-server/internal/errors"
	"github.com/keanutaufan/anitrackr-server/pkg/app_error"
	"github.com/keanutaufan/anitrackr-server/pkg/http_response"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	var appError app_error.AppError
	if !errors.As(err, &appError) {
		c.Logger().Error(err)
		appError = app_errors.ErrInternalServer
	} else {
		c.Logger().Error(err)
	}

	err = c.JSON(appError.HttpStatus, http_response.Response{
		Success: false,
		Message: appError.Message,
	})
	if err != nil {
		c.Logger().Error(err)
	}
}
