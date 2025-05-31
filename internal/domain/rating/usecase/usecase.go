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

func (uc *useCase) Create(ctx context.Context, req rating_request.StoreRating) (rating_dto.GetResponse, error) {
	tx, err := uc.txManager.Begin(ctx)
	if err != nil {
		return rating_dto.GetResponse{}, err
	}

	oldAnimeScore, err := uc.animeRepo.GetScore(ctx, tx, req.AnimeId)
	if err != nil {
		return rating_dto.GetResponse{}, err
	}

	oldCumulativeScore := oldAnimeScore.Score.Mul(decimal.NewFromInt(oldAnimeScore.ScoredBy))
	_, err = uc.animeRepo.UpdateScore(ctx, tx, anime_model.AnimeScore{
		ID:       oldAnimeScore.ID,
		Score:    oldCumulativeScore.Add(decimal.NewFromInt(int64(req.Score))).Div(decimal.NewFromInt(oldAnimeScore.ScoredBy + 1)),
		ScoredBy: oldAnimeScore.ScoredBy + 1,
	})

	result, err := uc.ratingRepo.Create(ctx, tx, req.ToModel())
	if err != nil {
		return rating_dto.GetResponse{}, err
	}

	err = uc.txManager.Commit(tx)
	if err != nil {
		return rating_dto.GetResponse{}, err
	}

	return (rating_dto.GetResponse{}).FromModel(result), nil
}

func (uc *useCase) FindOne(ctx context.Context, animeId, userId int64) (rating_dto.GetResponse, error) {
	result, err := uc.ratingRepo.FindOne(ctx, nil, animeId, userId)
	if err != nil {
		return rating_dto.GetResponse{}, err
	}

	return (rating_dto.GetResponse{}).FromModel(result), nil
}
