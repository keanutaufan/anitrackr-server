package anime_repository

import (
	"context"
	"github.com/keanutaufan/anitrackr-server/internal/domain/anime"
	"github.com/uptrace/bun"
)

type Repository interface {
	FindOne(ctx context.Context, tx bun.IDB, id int64) (anime.Anime, error)
}
