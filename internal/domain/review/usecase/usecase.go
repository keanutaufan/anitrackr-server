package review_usecase

import (
	"context"
	review_request "github.com/keanutaufan/anitrackr-server/internal/domain/review/dto/request"
	review_response "github.com/keanutaufan/anitrackr-server/internal/domain/review/dto/response"
	review_repository "github.com/keanutaufan/anitrackr-server/internal/domain/review/repository"
	app_errors "github.com/keanutaufan/anitrackr-server/internal/errors"
)

type useCase struct {
	reviewRepo review_repository.Repository
}

func NewUseCase(reviewRepo review_repository.Repository) UseCase {
	return &useCase{
		reviewRepo: reviewRepo,
	}
}

func (uc *useCase) Create(ctx context.Context, req review_request.StoreReview) (review_response.ShowReview, error) {
	result, err := uc.reviewRepo.Create(ctx, nil, req.ToModel())
	if err != nil {
		return review_response.ShowReview{}, err
	}

	return (review_response.ShowReview{}).FromModel(result), nil
}

func (uc *useCase) FindOne(ctx context.Context, reviewId int64) (review_response.ShowReview, error) {
	result, err := uc.reviewRepo.FindOne(ctx, nil, reviewId)
	if err != nil {
		return review_response.ShowReview{}, err
	}

	return (review_response.ShowReview{}).FromModel(result), nil
}

func (uc *useCase) Update(ctx context.Context, req review_request.UpdateReview) (review_response.ShowReview, error) {
	old, err := uc.reviewRepo.FindOne(ctx, nil, req.Id)
	if err != nil {
		return review_response.ShowReview{}, err
	}

	if old.UserId != req.UserId {
		return review_response.ShowReview{}, app_errors.ErrForbidden
	}

	result, err := uc.reviewRepo.Update(ctx, nil, req.ToModel())
	if err != nil {
		return review_response.ShowReview{}, err
	}

	return (review_response.ShowReview{}).FromModel(result), nil
}

func (uc *useCase) Delete(ctx context.Context, req review_request.DeleteReview) error {
	old, err := uc.reviewRepo.FindOne(ctx, nil, req.Id)
	if err != nil {
		return err
	}

	if old.UserId != req.UserId {
		return app_errors.ErrForbidden
	}

	err = uc.reviewRepo.Delete(ctx, nil, req.Id)
	if err != nil {
		return err
	}

	return nil
}
