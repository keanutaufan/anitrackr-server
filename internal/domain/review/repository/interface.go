package review_repository

import (
	"context"
	review_request "github.com/keanutaufan/anitrackr-server/internal/domain/review/dto/request"
	review_model "github.com/keanutaufan/anitrackr-server/internal/domain/review/model"
	"github.com/keanutaufan/anitrackr-server/pkg/pagination"
	"github.com/uptrace/bun"
)

type Repository interface {
	Create(ctx context.Context, tx bun.IDB, review review_model.Review) (review_model.Review, error)
	FindOne(ctx context.Context, tx bun.IDB, reviewId int64) (review_model.Review, error)
	FindWithPagination(ctx context.Context, tx bun.IDB, req review_request.IndexReview) ([]review_model.Review, pagination.PaginationMeta, error)
	Update(ctx context.Context, tx bun.IDB, review review_model.Review) (review_model.Review, error)
	Delete(ctx context.Context, tx bun.IDB, reviewId int64) error
}
