package anime_usecase

import (
	"context"
	anime_dto "github.com/keanutaufan/anitrackr-server/internal/domain/anime/dto"
	anime_repository "github.com/keanutaufan/anitrackr-server/internal/domain/anime/repository"
)

type useCase struct {
	animeRepo anime_repository.Repository
}

func NewUseCase(animeRepo anime_repository.Repository) UseCase {
	return &useCase{
		animeRepo: animeRepo,
	}
}

func (uc *useCase) FindOne(ctx context.Context, id int64) (anime_dto.GetResponse, error) {
	result, err := uc.animeRepo.FindOne(ctx, nil, id)
	if err != nil {
		return anime_dto.GetResponse{}, err
	}

	response := (anime_dto.GetResponse{}).FromModel(result)
	return response, nil
}
