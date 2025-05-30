package app

import (
	"context"
	"github.com/keanutaufan/anitrackr-server/internal/app/middlewares"
	"github.com/keanutaufan/anitrackr-server/internal/app/route_group"
	anime_handler "github.com/keanutaufan/anitrackr-server/internal/domain/anime/handler"
	anime_repository "github.com/keanutaufan/anitrackr-server/internal/domain/anime/repository"
	anime_usecase "github.com/keanutaufan/anitrackr-server/internal/domain/anime/usecase"
	auth_handler "github.com/keanutaufan/anitrackr-server/internal/domain/auth/handler"
	auth_usecase "github.com/keanutaufan/anitrackr-server/internal/domain/auth/usecase"
	user_repository "github.com/keanutaufan/anitrackr-server/internal/domain/user/repository"
	"github.com/keanutaufan/anitrackr-server/platform/database"
	"github.com/keanutaufan/anitrackr-server/platform/firebase_app"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func NewServer() *echo.Echo {
	bgCtx := context.Background()

	db := database.NewPostgresDatabase(database.LoadPostgresConfigFromEnv())
	firebaseClient, err := firebase_app.NewFirebaseClient(bgCtx, os.Getenv("FIREBASE_CREDENTIALS_FILE"))
	if err != nil {
		panic(err)
	}

	animeRepository := anime_repository.NewRepository(db)
	userRepository := user_repository.NewRepository(db)

	animeUseCase := anime_usecase.NewUseCase(animeRepository)
	authUseCase := auth_usecase.NewUseCase(userRepository)

	animeHandler := anime_handler.NewHandler(animeUseCase)
	authHandler := auth_handler.NewHandler(authUseCase)

	authMiddleware := middlewares.FirebaseAuthMiddleware(firebaseClient)

	engine := echo.New()
	engine.Use(middleware.Logger())
	route_group.GroupAnimeRoute(engine, animeHandler)
	route_group.GroupAuthRoute(engine, authHandler, authMiddleware)

	return engine
}
