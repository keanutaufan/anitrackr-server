package anime_repository

import (
	"context"
	"github.com/keanutaufan/anitrackr-server/internal/domain/anime/model"
	"github.com/uptrace/bun"
)

type Repository interface {
	FindOne(ctx context.Context, tx bun.IDB, id int64) (anime_model.Anime, error)
}
