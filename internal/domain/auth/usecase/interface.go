package auth_usecase

import (
	"context"
	user_dto "github.com/keanutaufan/anitrackr-server/internal/domain/auth/dto"
)

type UseCase interface {
	GetCurrentUser(ctx context.Context, userId int64) (user_dto.MeResponse, error)
}
