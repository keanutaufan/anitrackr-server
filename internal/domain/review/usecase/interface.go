package review_usecase

import (
	"context"
	review_request "github.com/keanutaufan/anitrackr-server/internal/domain/review/dto/request"
	review_response "github.com/keanutaufan/anitrackr-server/internal/domain/review/dto/response"
)

type UseCase interface {
	Create(ctx context.Context, review review_request.StoreReview) (review_response.ShowReview, error)
	FindOne(ctx context.Context, reviewId int64) (review_response.ShowReview, error)
}
