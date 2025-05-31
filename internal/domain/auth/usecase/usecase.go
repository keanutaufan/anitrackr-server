package auth_usecase

import (
	"context"
	user_dto "github.com/keanutaufan/anitrackr-server/internal/domain/auth/dto"
	user_repository "github.com/keanutaufan/anitrackr-server/internal/domain/user/repository"
)

type useCase struct {
	userRepo user_repository.Repository
}

func NewUseCase(userRepo user_repository.Repository) UseCase {
	return &useCase{
		userRepo: userRepo,
	}
}

func (uc *useCase) GetCurrentUser(ctx context.Context, userId int64) (user_dto.MeResponse, error) {
	result, err := uc.userRepo.FindOne(ctx, nil, userId)
	if err != nil {
		return user_dto.MeResponse{}, err
	}

	response := (user_dto.MeResponse{}).FromModel(result)
	return response, nil
}
