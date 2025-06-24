package review_repository

import (
	"context"
	review_request "github.com/keanutaufan/anitrackr-server/internal/domain/review/dto/request"
	review_model "github.com/keanutaufan/anitrackr-server/internal/domain/review/model"
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

func (r *repository) FindOne(ctx context.Context, tx bun.IDB, reviewId int64) (review_model.ReviewDenormalized, error) {
	if tx == nil {
		tx = r.db
	}

	var result review_model.ReviewDenormalized
	err := tx.NewSelect().
		Model(&result).
		ColumnExpr("review_denormalized.*").
		ColumnExpr("anime.title as anime_title").
		Relation("User").
		Join("JOIN anime ON anime.id = review_denormalized.anime_id").
		Where("review_denormalized.id = ?", reviewId).
		Scan(ctx)

	if err != nil {
		if db_error.IsNotFound(err) {
			return review_model.ReviewDenormalized{}, app_errors.ErrNotFound
		}
		return review_model.ReviewDenormalized{}, err
	}

	return result, nil
}

func (r *repository) FindWithPagination(ctx context.Context, tx bun.IDB, req review_request.IndexReview) ([]review_model.ReviewDenormalized, pagination.PaginationMeta, error) {
	if tx == nil {
		tx = r.db
	}

	var result []review_model.ReviewDenormalized
	query := tx.NewSelect().
		Model(&result).
		ColumnExpr("review_denormalized.*").
		ColumnExpr("anime.title as anime_title").
		Relation("User").
		Join("JOIN anime ON anime.id = review_denormalized.anime_id")

	if req.AnimeId != 0 {
		query.Where("anime_id = ?", req.AnimeId)
	}

	if req.UserId != 0 {
		query.Where("user_id = ?", req.UserId)
	}

	if req.SortBy != "" && req.SortDir != "" {
		query.OrderExpr("? ?", bun.Ident(req.SortBy), bun.Safe(req.SortDir))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return []review_model.ReviewDenormalized{}, pagination.PaginationMeta{}, err
	}

	query = query.Limit(req.PageSize).Offset(req.PageSize * (req.Page - 1))
	err = query.Scan(ctx, &result)
	if err != nil {
		return []review_model.ReviewDenormalized{}, pagination.PaginationMeta{}, err
	}

	meta := pagination.PaginationMeta{
		Page:     req.Page,
		PageSize: req.PageSize,
		MaxPage:  int(math.Ceil(float64(count) / float64(req.PageSize))),
		Count:    count,
	}

	return result, meta, nil

}

func (r *repository) Update(ctx context.Context, tx bun.IDB, review review_model.Review) (review_model.Review, error) {
	if tx == nil {
		tx = r.db
	}

	_, err := tx.NewUpdate().
		Model(&review).
		Set("title = ?", review.Title).
		Set("body = ?", review.Body).
		Set("is_liked = ?", review.IsLiked).
		Where("id = ?", review.ID).
		Returning("*").
		Exec(ctx)

	if err != nil {
		return review_model.Review{}, err
	}

	return review, nil
}

func (r *repository) Delete(ctx context.Context, tx bun.IDB, reviewId int64) error {
	if tx == nil {
		tx = r.db
	}

	_, err := tx.NewDelete().
		Model((*review_model.Review)(nil)).
		Where("id = ?", reviewId).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
