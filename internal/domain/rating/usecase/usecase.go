package rating_usecase

import (
	"context"
	anime_model "github.com/keanutaufan/anitrackr-server/internal/domain/anime/model"
	anime_repository "github.com/keanutaufan/anitrackr-server/internal/domain/anime/repository"
	rating_request "github.com/keanutaufan/anitrackr-server/internal/domain/rating/dto/request"
	rating_dto "github.com/keanutaufan/anitrackr-server/internal/domain/rating/dto/response"
	rating_repository "github.com/keanutaufan/anitrackr-server/internal/domain/rating/repository"
	"github.com/keanutaufan/anitrackr-server/platform/database"
	"github.com/shopspring/decimal"
)

type useCase struct {
	txManager  database.TxManager
	ratingRepo rating_repository.Repository
	animeRepo  anime_repository.Repository
}

func NewUseCase(txManager database.TxManager, ratingRepo rating_repository.Repository, animeRepo anime_repository.Repository) UseCase {
	return &useCase{
		txManager:  txManager,
		ratingRepo: ratingRepo,
		animeRepo:  animeRepo,
	}
}

func (uc *useCase) Create(ctx context.Context, req rating_request.StoreRating) (rating_dto.ShowRating, error) {
	tx, err := uc.txManager.Begin(ctx)
	if err != nil {
		return rating_dto.ShowRating{}, err
	}

	oldAnimeScore, err := uc.animeRepo.GetScore(ctx, tx, req.AnimeId)
	if err != nil {
		return rating_dto.ShowRating{}, err
	}

	oldCumulativeScore := oldAnimeScore.Score.Mul(decimal.NewFromInt(oldAnimeScore.ScoredBy))
	_, err = uc.animeRepo.UpdateScore(ctx, tx, anime_model.AnimeScore{
		ID:       oldAnimeScore.ID,
		Score:    oldCumulativeScore.Add(decimal.NewFromInt(int64(req.Score))).Div(decimal.NewFromInt(oldAnimeScore.ScoredBy + 1)),
		ScoredBy: oldAnimeScore.ScoredBy + 1,
	})

	result, err := uc.ratingRepo.Create(ctx, tx, req.ToModel())
	if err != nil {
		return rating_dto.ShowRating{}, err
	}

	err = uc.txManager.Commit(tx)
	if err != nil {
		return rating_dto.ShowRating{}, err
	}

	return (rating_dto.ShowRating{}).FromModel(result), nil
}

func (uc *useCase) FindOne(ctx context.Context, req rating_request.ShowRating) (rating_dto.ShowRating, error) {
	result, err := uc.ratingRepo.FindOne(ctx, nil, req.AnimeId, req.UserId)
	if err != nil {
		return rating_dto.ShowRating{}, err
	}

	return (rating_dto.ShowRating{}).FromModel(result), nil
}

func (uc *useCase) Update(ctx context.Context, req rating_request.UpdateRating) (rating_dto.ShowRating, error) {
	tx, err := uc.txManager.Begin(ctx)
	if err != nil {
		return rating_dto.ShowRating{}, err
	}

	oldAnimeScore, err := uc.animeRepo.GetScore(ctx, tx, req.AnimeId)
	if err != nil {
		return rating_dto.ShowRating{}, err
	}

	oldRating, err := uc.ratingRepo.FindOne(ctx, tx, req.AnimeId, req.UserId)
	if err != nil {
		return rating_dto.ShowRating{}, err
	}

	newAnimeScore := oldAnimeScore.Score.
		Mul(decimal.NewFromInt(oldAnimeScore.ScoredBy)).
		Sub(decimal.NewFromInt(int64(oldRating.Score))).
		Add(decimal.NewFromInt(int64(req.Score))).
		Div(decimal.NewFromInt(oldAnimeScore.ScoredBy))

	_, err = uc.animeRepo.UpdateScore(ctx, tx, anime_model.AnimeScore{
		ID:       oldAnimeScore.ID,
		Score:    newAnimeScore,
		ScoredBy: oldAnimeScore.ScoredBy,
	})

	result, err := uc.ratingRepo.Update(ctx, tx, req.ToModel())
	if err != nil {
		return rating_dto.ShowRating{}, err
	}

	err = uc.txManager.Commit(tx)
	if err != nil {
		return rating_dto.ShowRating{}, err
	}

	return (rating_dto.ShowRating{}).FromModel(result), nil
}
