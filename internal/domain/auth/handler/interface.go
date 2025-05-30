package auth_handler

import "github.com/labstack/echo/v4"

type Handler interface {
	Me(ctx echo.Context) error
}
