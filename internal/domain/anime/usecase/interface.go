package anime_usecase

import (
	"context"
	anime_request "github.com/keanutaufan/anitrackr-server/internal/domain/anime/dto/request"
	anime_response "github.com/keanutaufan/anitrackr-server/internal/domain/anime/dto/response"
)

type UseCase interface {
	FindOne(ctx context.Context, req anime_request.ShowWithUser) (anime_response.ShowWithUser, error)
}
