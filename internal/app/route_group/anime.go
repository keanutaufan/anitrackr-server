package route_group

import (
	anime_handler "github.com/keanutaufan/anitrackr-server/internal/domain/anime/handler"
	"github.com/labstack/echo/v4"
)

func GroupAnimeRoute(router *echo.Echo, animeHandler anime_handler.Handler, authMiddleware echo.MiddlewareFunc) *echo.Group {
	group := router.Group("/anime")

	group.GET("", animeHandler.Index, authMiddleware)
	group.GET("/:animeId", animeHandler.Show, authMiddleware)

	return group
}
