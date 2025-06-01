package review_usecase

import (
	"context"
	review_request "github.com/keanutaufan/anitrackr-server/internal/domain/review/dto/request"
	review_response "github.com/keanutaufan/anitrackr-server/internal/domain/review/dto/response"
	"github.com/keanutaufan/anitrackr-server/pkg/pagination"
)

type UseCase interface {
	Create(ctx context.Context, review review_request.StoreReview) (review_response.ShowReview, error)
	FindOne(ctx context.Context, reviewId int64) (review_response.ShowReview, error)
	FindWithPagination(ctx context.Context, req review_request.IndexReview) (review_response.IndexReview, pagination.PaginationMeta, error)
	Update(ctx context.Context, review review_request.UpdateReview) (review_response.ShowReview, error)
	Delete(ctx context.Context, req review_request.DeleteReview) error
}
