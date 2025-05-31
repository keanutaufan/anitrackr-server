package anime_repository

import (
	"context"
	"github.com/keanutaufan/anitrackr-server/internal/domain/anime/model"
	app_errors "github.com/keanutaufan/anitrackr-server/internal/errors"
	"github.com/keanutaufan/anitrackr-server/pkg/db_error"
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

func (r *repository) FindOne(ctx context.Context, tx bun.IDB, id int64) (anime_model.Anime, error) {
	if tx == nil {
		tx = r.db
	}

	var result anime_model.Anime
	err := tx.NewSelect().
		Model(&result).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		if db_error.IsNotFound(err) {
			return anime_model.Anime{}, app_errors.ErrNotFound
		}
		return anime_model.Anime{}, err
	}

	return result, nil
}

func (r *repository) GetScore(ctx context.Context, tx bun.IDB, id int64) (anime_model.AnimeScore, error) {
	if tx == nil {
		tx = r.db
	}

	var result anime_model.AnimeScore
	err := tx.NewSelect().
		Model(&result).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		if db_error.IsNotFound(err) {
			return anime_model.AnimeScore{}, app_errors.ErrNotFound
		}
		return anime_model.AnimeScore{}, err
	}

	return result, nil
}

func (r *repository) UpdateScore(ctx context.Context, tx bun.IDB, score anime_model.AnimeScore) (anime_model.AnimeScore, error) {
	if tx == nil {
		tx = r.db
	}

	_, err := tx.NewUpdate().
		Model(&score).
		Where("id = ?", score.ID).
		Exec(ctx)

	if err != nil {
		if db_error.IsUniqueViolation(err) {
			return anime_model.AnimeScore{}, app_errors.ErrAlreadyExists
		}
		return anime_model.AnimeScore{}, err
	}

	return score, nil
}
