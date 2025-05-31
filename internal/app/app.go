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
	rating_handler "github.com/keanutaufan/anitrackr-server/internal/domain/rating/handler"
	rating_repository "github.com/keanutaufan/anitrackr-server/internal/domain/rating/repository"
	rating_usecase "github.com/keanutaufan/anitrackr-server/internal/domain/rating/usecase"
	review_handler "github.com/keanutaufan/anitrackr-server/internal/domain/review/handler"
	review_repository "github.com/keanutaufan/anitrackr-server/internal/domain/review/repository"
	review_usecase "github.com/keanutaufan/anitrackr-server/internal/domain/review/usecase"
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

	txManager := database.NewPostgresTxManager(db)

	animeRepository := anime_repository.NewRepository(db)
	userRepository := user_repository.NewRepository(db)
	reviewRepository := review_repository.NewRepository(db)
	ratingRepository := rating_repository.NewRepository(db)

	animeUseCase := anime_usecase.NewUseCase(animeRepository)
	authUseCase := auth_usecase.NewUseCase(userRepository)
	reviewUseCase := review_usecase.NewUseCase(reviewRepository)
	ratingUseCase := rating_usecase.NewUseCase(txManager, ratingRepository, animeRepository)

	animeHandler := anime_handler.NewHandler(animeUseCase)
	authHandler := auth_handler.NewHandler(authUseCase)
	reviewHandler := review_handler.NewHandler(reviewUseCase)
	ratingHandler := rating_handler.NewHandler(ratingUseCase)

	authMiddleware := middlewares.FirebaseAuthMiddleware(firebaseClient)

	engine := echo.New()
	engine.HTTPErrorHandler = middlewares.ErrorHandler
	engine.Use(middleware.Logger())
	route_group.GroupAnimeRoute(engine, animeHandler)
	route_group.GroupAuthRoute(engine, authHandler, authMiddleware)
	route_group.GroupReviewRoute(engine, reviewHandler, authMiddleware)
	route_group.GroupRatingRoute(engine, ratingHandler, authMiddleware)

	return engine
}
