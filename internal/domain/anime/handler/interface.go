package anime_handler

import (
	"github.com/labstack/echo/v4"
)

type Handler interface {
	Index(c echo.Context) error
	Show(c echo.Context) error
}
