package rating_repository

import (
	"context"
	rating_model "github.com/keanutaufan/anitrackr-server/internal/domain/rating/model"
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

func (r *repository) Create(ctx context.Context, tx bun.IDB, rating rating_model.Rating) (rating_model.Rating, error) {
	if tx == nil {
		tx = r.db
	}

	_, err := tx.NewInsert().
		Model(&rating).
		Exec(ctx)

	if err != nil {
		if db_error.IsUniqueViolation(err) {
			return rating_model.Rating{}, app_errors.ErrAlreadyExists
		}
		return rating_model.Rating{}, err
	}

	return rating, nil
}

func (r *repository) Update(ctx context.Context, tx bun.IDB, rating rating_model.Rating) (rating_model.Rating, error) {
	if tx == nil {
		tx = r.db
	}

	_, err := tx.NewUpdate().
		Model(&rating).
		WherePK().
		Exec(ctx)

	if err != nil {
		return rating_model.Rating{}, err
	}

	return rating, nil
}

func (r *repository) FindOne(ctx context.Context, tx bun.IDB, animeId, userId int64) (rating_model.Rating, error) {
	if tx == nil {
		tx = r.db
	}

	var result rating_model.Rating
	err := tx.NewSelect().
		Model(&result).
		Where("anime_id = ? AND user_id = ?", animeId, userId).
		Scan(ctx)

	if err != nil {
		if db_error.IsNotFound(err) {
			return rating_model.Rating{}, app_errors.ErrNotFound
		}
		return rating_model.Rating{}, err
	}

	return result, nil
}
