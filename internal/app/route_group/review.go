package route_group

import (
	review_handler "github.com/keanutaufan/anitrackr-server/internal/domain/review/handler"
	"github.com/labstack/echo/v4"
)

func GroupReviewRoute(router *echo.Echo, reviewHandler review_handler.Handler, authMiddleware echo.MiddlewareFunc) *echo.Group {
	group := router.Group("/review")

	group.POST("", reviewHandler.Store, authMiddleware)
	group.GET("/:reviewId", reviewHandler.Show)

	return group
}
