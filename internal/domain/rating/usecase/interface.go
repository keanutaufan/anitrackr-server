package rating_usecase

import (
	"context"
	rating_request "github.com/keanutaufan/anitrackr-server/internal/domain/rating/dto/request"
	rating_dto "github.com/keanutaufan/anitrackr-server/internal/domain/rating/dto/response"
)

type UseCase interface {
	Create(ctx context.Context, req rating_request.StoreRating) (rating_dto.ShowRating, error)
	FindOne(ctx context.Context, rating rating_request.ShowRating) (rating_dto.ShowRating, error)
	Update(ctx context.Context, req rating_request.UpdateRating) (rating_dto.ShowRating, error)
	Delete(ctx context.Context, req rating_request.DeleteRating) error
}
