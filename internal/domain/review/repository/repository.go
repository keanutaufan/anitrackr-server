package review_repository

import (
	"context"
	review_model "github.com/keanutaufan/anitrackr-server/internal/domain/review/model"
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

func (r *repository) Create(ctx context.Context, tx bun.IDB, review review_model.Review) (review_model.Review, error) {
	if tx == nil {
		tx = r.db
	}

	_, err := tx.NewInsert().
		Model(&review).
		Exec(ctx)

	if err != nil {
		if db_error.IsUniqueViolation(err) {
			return review_model.Review{}, app_errors.ErrAlreadyExists
		}
		return review_model.Review{}, err
	}

	return review, nil
}

func (r *repository) FindOne(ctx context.Context, tx bun.IDB, id int64) (review_model.Review, error) {
	if tx == nil {
		tx = r.db
	}

	var result review_model.Review
	err := tx.NewSelect().
		Model(&result).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		if db_error.IsNotFound(err) {
			return review_model.Review{}, app_errors.ErrNotFound
		}
		return review_model.Review{}, err
	}

	return result, nil
}

func (r *repository) Update(ctx context.Context, tx bun.IDB, review review_model.Review) (review_model.Review, error) {
	if tx == nil {
		tx = r.db
	}

	_, err := tx.NewUpdate().
		Model(&review).
		Set("title = ?", review.Title).
		Set("body = ?", review.Body).
		Where("id = ?", review.ID).
		Returning("*").
		Exec(ctx)

	if err != nil {
		return review_model.Review{}, err
	}

	return review, nil
}
