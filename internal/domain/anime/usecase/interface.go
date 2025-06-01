package anime_usecase

import (
	"context"
	anime_request "github.com/keanutaufan/anitrackr-server/internal/domain/anime/dto/request"
	anime_response "github.com/keanutaufan/anitrackr-server/internal/domain/anime/dto/response"
	"github.com/keanutaufan/anitrackr-server/pkg/pagination"
)

type UseCase interface {
	FindOne(ctx context.Context, req anime_request.ShowWithUser) (anime_response.ShowWithUser, error)
	FindWithPagination(ctx context.Context, req anime_request.IndexAnime) (anime_response.IndexAnime, pagination.PaginationMeta, error)
}
