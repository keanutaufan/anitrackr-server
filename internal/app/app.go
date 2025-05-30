package app

import (
	"github.com/keanutaufan/anitrackr-server/internal/app/route_group"
	anime_handler "github.com/keanutaufan/anitrackr-server/internal/domain/anime/handler"
	anime_repository "github.com/keanutaufan/anitrackr-server/internal/domain/anime/repository"
	anime_usecase "github.com/keanutaufan/anitrackr-server/internal/domain/anime/usecase"
	"github.com/keanutaufan/anitrackr-server/platform/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer() *echo.Echo {
	db := database.NewPostgresDatabase(database.LoadPostgresConfigFromEnv())

	animeRepository := anime_repository.NewRepository(db)
	animeUseCase := anime_usecase.NewUseCase(animeRepository)
	animeHandler := anime_handler.NewHandler(animeUseCase)

	engine := echo.New()
	engine.Use(middleware.Logger())
	route_group.GroupAnimeRoute(engine, animeHandler)

	return engine
}
