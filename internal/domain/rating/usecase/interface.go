package rating_usecase

import (
	"context"
	rating_request "github.com/keanutaufan/anitrackr-server/internal/domain/rating/dto/request"
	rating_dto "github.com/keanutaufan/anitrackr-server/internal/domain/rating/dto/response"
)

type UseCase interface {
	Create(ctx context.Context, req rating_request.StoreRating) (rating_dto.GetResponse, error)
	FindOne(ctx context.Context, animeId, userId int64) (rating_dto.GetResponse, error)
	Update(ctx context.Context, req rating_request.UpdateRating) (rating_dto.GetResponse, error)
}
