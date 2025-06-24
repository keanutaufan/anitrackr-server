package list_repository

import (
	"context"
	list_model "github.com/keanutaufan/anitrackr-server/internal/domain/list/model"
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

func (r *repository) Create(ctx context.Context, tx bun.IDB, list list_model.List) (list_model.List, error) {
	if tx == nil {
		tx = r.db
	}

	_, err := tx.NewInsert().
		Model(&list).
		Exec(ctx)

	if err != nil {
		if db_error.IsUniqueViolation(err) {
			return list_model.List{}, app_errors.ErrAlreadyExists
		}
		return list_model.List{}, err
	}

	return list, nil
}

func (r *repository) Update(ctx context.Context, tx bun.IDB, list list_model.List) (list_model.List, error) {
	if tx == nil {
		tx = r.db
	}

	_, err := tx.NewUpdate().
		Model(&list).
		Column("name").
		Column("episode_watched").
		WherePK().
		Exec(ctx)

	if err != nil {
		return list_model.List{}, err
	}

	return list, nil
}

func (r *repository) FindOne(ctx context.Context, tx bun.IDB, animeId, userId int64) (list_model.List, error) {
	if tx == nil {
		tx = r.db
	}

	var result list_model.List
	err := tx.NewSelect().
		Model(&result).
		Where("anime_id = ? AND user_id = ?", animeId, userId).
		Scan(ctx)

	if err != nil {
		if db_error.IsNotFound(err) {
			return list_model.List{}, app_errors.ErrNotFound
		}
		return list_model.List{}, err
	}

	return result, nil
}

func (r *repository) Delete(ctx context.Context, tx bun.IDB, animeId, userId int64) error {
	if tx == nil {
		tx = r.db
	}

	_, err := tx.NewDelete().
		Model(&list_model.List{}).
		Where("anime_id = ? AND user_id = ?", animeId, userId).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
