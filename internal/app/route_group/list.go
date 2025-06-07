package route_group

import (
	list_handler "github.com/keanutaufan/anitrackr-server/internal/domain/list/handler"
	"github.com/labstack/echo/v4"
)

func GroupListRoute(router *echo.Echo, listHandler list_handler.Handler, authMiddleware echo.MiddlewareFunc) *echo.Group {
	group := router.Group("/list")

	group.POST("", listHandler.Store, authMiddleware)
	group.PUT("/:animeId", listHandler.Update, authMiddleware)
	group.GET("/:animeId", listHandler.Show, authMiddleware)
	group.DELETE("/:animeId", listHandler.Delete, authMiddleware)

	return group
}
