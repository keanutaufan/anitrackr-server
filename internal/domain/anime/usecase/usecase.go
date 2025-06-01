package anime_usecase

import (
	"context"
	anime_request "github.com/keanutaufan/anitrackr-server/internal/domain/anime/dto/request"
	anime_response "github.com/keanutaufan/anitrackr-server/internal/domain/anime/dto/response"
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

func (uc *useCase) FindOne(ctx context.Context, req anime_request.ShowWithUser) (anime_response.ShowWithUser, error) {
	result, err := uc.animeRepo.FindOneWithUserProperties(ctx, nil, req.AnimeId, req.UserId)
	if err != nil {
		return anime_response.ShowWithUser{}, err
	}

	response := (anime_response.ShowWithUser{}).FromModel(result)
	return response, nil
}
