package list_usecase

import (
	"context"
	list_request "github.com/keanutaufan/anitrackr-server/internal/domain/list/dto/request"
	list_response "github.com/keanutaufan/anitrackr-server/internal/domain/list/dto/response"
	list_repository "github.com/keanutaufan/anitrackr-server/internal/domain/list/repository"
	app_errors "github.com/keanutaufan/anitrackr-server/internal/errors"
)

type useCase struct {
	listRepo list_repository.Repository
}

func NewUseCase(listRepo list_repository.Repository) UseCase {
	return &useCase{
		listRepo: listRepo,
	}
}

func (uc *useCase) Create(ctx context.Context, req list_request.StoreList) (list_response.ShowList, error) {
	result, err := uc.listRepo.Create(ctx, nil, req.ToModel())
	if err != nil {
		return list_response.ShowList{}, err
	}

	return (list_response.ShowList{}).FromModel(result), nil
}

func (uc *useCase) FindOne(ctx context.Context, req list_request.ShowList) (list_response.ShowList, error) {
	result, err := uc.listRepo.FindOne(ctx, nil, req.AnimeId, req.UserId)
	if err != nil {
		return list_response.ShowList{}, err
	}

	return (list_response.ShowList{}).FromModel(result), nil
}

func (uc *useCase) Update(ctx context.Context, req list_request.UpdateList) (list_response.ShowList, error) {
	old, err := uc.listRepo.FindOne(ctx, nil, req.AnimeId, req.UserId)
	if err != nil {
		return list_response.ShowList{}, err
	}

	if old.UserID != req.UserId {
		return list_response.ShowList{}, app_errors.ErrForbidden
	}

	result, err := uc.listRepo.Update(ctx, nil, req.ToModel())
	if err != nil {
		return list_response.ShowList{}, err
	}

	return (list_response.ShowList{}).FromModel(result), nil
}

func (uc *useCase) Delete(ctx context.Context, req list_request.DeleteList) error {
	old, err := uc.listRepo.FindOne(ctx, nil, req.AnimeId, req.UserId)
	if err != nil {
		return err
	}

	if old.UserID != req.UserId {
		return app_errors.ErrForbidden
	}

	err = uc.listRepo.Delete(ctx, nil, req.AnimeId, req.UserId)
	if err != nil {
		return err
	}

	return nil
}
