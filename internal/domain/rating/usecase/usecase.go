package rating_usecase

import (
	"context"
	rating_dto "github.com/keanutaufan/anitrackr-server/internal/domain/rating/dto"
	rating_request "github.com/keanutaufan/anitrackr-server/internal/domain/rating/dto/request"
	rating_model "github.com/keanutaufan/anitrackr-server/internal/domain/rating/model"
	rating_repository "github.com/keanutaufan/anitrackr-server/internal/domain/rating/repository"
	user_repository "github.com/keanutaufan/anitrackr-server/internal/domain/user/repository"
)

type useCase struct {
	ratingRepo rating_repository.Repository
	userRepo   user_repository.Repository
}

func NewUseCase(ratingRepo rating_repository.Repository, userRepo user_repository.Repository) UseCase {
	return &useCase{
		ratingRepo: ratingRepo,
	}
}

func (uc *useCase) Upsert(ctx context.Context, req rating_request.UpsertRequest) (rating_dto.GetResponse, error) {
	user, err := uc.userRepo.FindOneByUid(ctx, nil, req.UserUid)
	if err != nil {
		return rating_dto.GetResponse{}, err
	}

	result, err := uc.ratingRepo.Upsert(ctx, nil, rating_model.Rating{
		AnimeID: req.AnimeId,
		UserID:  user.ID,
		Score:   req.Score,
	})
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

	response := (rating_dto.GetResponse{}).FromModel(result)
	return response, nil
}
