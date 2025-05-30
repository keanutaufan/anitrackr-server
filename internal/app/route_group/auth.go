package route_group

import (
	auth_handler "github.com/keanutaufan/anitrackr-server/internal/domain/auth/handler"
	"github.com/labstack/echo/v4"
)

func GroupAuthRoute(router *echo.Echo, authHandler auth_handler.Handler, authMiddleware echo.MiddlewareFunc) *echo.Group {
	group := router.Group("/auth")

	group.GET("/me", authHandler.Me, authMiddleware)

	return group
}
