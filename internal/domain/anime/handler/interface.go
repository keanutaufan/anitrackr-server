package anime_handler

import (
	"github.com/labstack/echo/v4"
)

type Handler interface {
	Show(ctx echo.Context) error
}
