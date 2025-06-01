package anime_repository

import (
	"context"
	anime_request "github.com/keanutaufan/anitrackr-server/internal/domain/anime/dto/request"
	"github.com/keanutaufan/anitrackr-server/internal/domain/anime/model"
	"github.com/keanutaufan/anitrackr-server/pkg/pagination"
	"github.com/uptrace/bun"
)

type Repository interface {
	FindOne(ctx context.Context, tx bun.IDB, animeId int64) (anime_model.Anime, error)
	FindOneWithUserProperties(ctx context.Context, tx bun.IDB, animeId, userId int64) (anime_model.UserAnime, error)
	FindWithPagination(ctx context.Context, tx bun.IDB, req anime_request.IndexAnime) ([]anime_model.UserAnime, pagination.PaginationMeta, error)
	GetScore(ctx context.Context, tx bun.IDB, animeId int64) (anime_model.AnimeScore, error)
	UpdateScore(ctx context.Context, tx bun.IDB, score anime_model.AnimeScore) (anime_model.AnimeScore, error)
}
