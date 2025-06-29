package list_handler

import "github.com/labstack/echo/v4"

type Handler interface {
	Store(c echo.Context) error
	Show(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
