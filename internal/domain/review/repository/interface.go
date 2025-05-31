package review_repository

import (
	"context"
	review_model "github.com/keanutaufan/anitrackr-server/internal/domain/review/model"
	"github.com/uptrace/bun"
)

type Repository interface {
	Create(ctx context.Context, tx bun.IDB, review review_model.Review) (review_model.Review, error)
	FindOne(ctx context.Context, tx bun.IDB, id int64) (review_model.Review, error)
	Update(ctx context.Context, tx bun.IDB, review review_model.Review) (review_model.Review, error)
}
