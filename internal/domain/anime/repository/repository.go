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

func (r *repository) FindOne(ctx context.Context, tx bun.IDB, animeId int64) (anime_model.Anime, error) {
	if tx == nil {
		tx = r.db
	}

	var result anime_model.Anime
	err := tx.NewSelect().
		Model(&result).
		Where("id = ?", animeId).
		Scan(ctx)

	if err != nil {
		if db_error.IsNotFound(err) {
			return anime_model.Anime{}, app_errors.ErrNotFound
		}
		return anime_model.Anime{}, err
	}

	return result, nil
}

func (r *repository) FindOneWithUserProperties(ctx context.Context, tx bun.IDB, animeId, userId int64) (anime_model.UserAnime, error) {
	if tx == nil {
		tx = r.db
	}

	var result anime_model.UserAnime
	err := tx.NewSelect().
		Table("anime").
		Column("anime.*").
		ColumnExpr("reviews.id AS user_review__id").
		ColumnExpr("reviews.title AS user_review__title").
		ColumnExpr("reviews.body AS user_review__body").
		ColumnExpr("reviews.anime_id AS user_review__anime_id").
		ColumnExpr("reviews.user_id AS user_review__user_id").
		ColumnExpr("reviews.created_at AS user_review__created_at").
		ColumnExpr("reviews.updated_at AS user_review__updated_at").
		ColumnExpr("COALESCE(ratings.score, 0) AS user_score").
		Join("LEFT JOIN reviews ON reviews.anime_id = anime.id AND reviews.user_id = ?", userId).
		Join("LEFT JOIN ratings ON ratings.anime_id = anime.id AND ratings.user_id = ?", userId).
		Where("anime.id = ?", animeId).
		Scan(ctx, &result)

	if err != nil {
		if db_error.IsNotFound(err) {
			return anime_model.UserAnime{}, app_errors.ErrNotFound
		}
		return anime_model.UserAnime{}, err
	}

	return result, nil
}

func (r *repository) GetScore(ctx context.Context, tx bun.IDB, animeId int64) (anime_model.AnimeScore, error) {
	if tx == nil {
		tx = r.db
	}

	var result anime_model.AnimeScore
	err := tx.NewSelect().
		Model(&result).
		Where("id = ?", animeId).
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
