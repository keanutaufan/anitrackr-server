package rating_handler

import "github.com/labstack/echo/v4"

type Handler interface {
	Store(c echo.Context) error
	Show(c echo.Context) error
}
