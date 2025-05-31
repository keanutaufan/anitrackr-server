package rating_repository

import (
	"context"
	rating_model "github.com/keanutaufan/anitrackr-server/internal/domain/rating/model"
	"github.com/uptrace/bun"
)

type Repository interface {
	Create(ctx context.Context, tx bun.IDB, rating rating_model.Rating) (rating_model.Rating, error)
	Update(ctx context.Context, tx bun.IDB, rating rating_model.Rating) (rating_model.Rating, error)
	FindOne(ctx context.Context, tx bun.IDB, animeId, userId int64) (rating_model.Rating, error)
}
