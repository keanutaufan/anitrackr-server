package anime_repository

import (
	"context"
	anime_request "github.com/keanutaufan/anitrackr-server/internal/domain/anime/dto/request"
	"github.com/keanutaufan/anitrackr-server/internal/domain/anime/model"
	app_errors "github.com/keanutaufan/anitrackr-server/internal/errors"
	"github.com/keanutaufan/anitrackr-server/pkg/db_error"
	"github.com/keanutaufan/anitrackr-server/pkg/pagination"
	"github.com/uptrace/bun"
	"math"
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
		ColumnExpr("lists.name AS user_list_name").
		Join("LEFT JOIN reviews ON reviews.anime_id = anime.id AND reviews.user_id = ?", userId).
		Join("LEFT JOIN ratings ON ratings.anime_id = anime.id AND ratings.user_id = ?", userId).
		Join("LEFT JOIN lists ON lists.anime_id = anime.id AND lists.user_id = ?", userId).
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

func (r *repository) FindWithPagination(ctx context.Context, tx bun.IDB, req anime_request.IndexAnime) ([]anime_model.UserAnime, pagination.PaginationMeta, error) {
	if tx == nil {
		tx = r.db
	}

	var result []anime_model.UserAnime
	query := tx.NewSelect().
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
		ColumnExpr("lists.name AS user_list_name").
		Join("LEFT JOIN reviews ON reviews.anime_id = anime.id AND reviews.user_id = ?", req.UserId).
		Join("LEFT JOIN ratings ON ratings.anime_id = anime.id AND ratings.user_id = ?", req.UserId).
		Join("LEFT JOIN lists ON lists.anime_id = anime.id AND lists.user_id = ?", req.UserId)

	if req.Search != "" {
		query = query.Where(
			"anime.title || ' ' || COALESCE(anime.title_english, '') || ' ' || anime.title_japanese || ' ' || COALESCE(anime.title_synonyms) ILIKE ?",
			"%"+req.Search+"%",
		)
	}

	if req.MinUserScore != 0 {
		query = query.Where("COALESCE(ratings.score, 0) >= ?", req.MinUserScore)
	}

	if req.ListName != "" {
		query = query.Where("lists.name = ?", req.ListName)
	}

	if req.SortBy != "" && req.SortDir != "" {
		query.OrderExpr("? ?", bun.Ident(req.SortBy), bun.Safe(req.SortDir))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, pagination.PaginationMeta{}, err
	}

	query = query.Limit(req.PageSize).Offset(req.PageSize * (req.Page - 1))
	err = query.Scan(ctx, &result)
	if err != nil {
		return nil, pagination.PaginationMeta{}, err
	}

	meta := pagination.PaginationMeta{
		Page:     req.Page,
		PageSize: req.PageSize,
		MaxPage:  int(math.Ceil(float64(count) / float64(req.PageSize))),
		Count:    count,
	}

	return result, meta, nil
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
