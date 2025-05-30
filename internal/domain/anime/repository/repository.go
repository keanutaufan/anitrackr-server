package anime_repository

import (
	"context"
	"github.com/keanutaufan/anitrackr-server/internal/domain/anime"
	"github.com/uptrace/bun"
)

type repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindOne(ctx context.Context, tx bun.IDB, id int64) (anime.Anime, error) {
	if tx == nil {
		tx = r.db
	}

	var result anime.Anime
	err := tx.NewSelect().
		Model(&result).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return anime.Anime{}, err
	}

	return result, nil
}
