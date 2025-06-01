package route_group

import (
	rating_handler "github.com/keanutaufan/anitrackr-server/internal/domain/rating/handler"
	"github.com/labstack/echo/v4"
)

func GroupRatingRoute(router *echo.Echo, ratingHandler rating_handler.Handler, authMiddleware echo.MiddlewareFunc) *echo.Group {
	group := router.Group("/rating")

	group.POST("", ratingHandler.Store, authMiddleware)
	group.PUT("/:animeId", ratingHandler.Update, authMiddleware)
	group.GET("/:animeId", ratingHandler.Show, authMiddleware)
	group.DELETE("/:animeId", ratingHandler.Delete, authMiddleware)

	return group
}
