package anime_usecase

import (
	"context"
	anime_dto "github.com/keanutaufan/anitrackr-server/internal/domain/anime/dto"
)

type UseCase interface {
	FindOne(ctx context.Context, id int64) (anime_dto.GetResponse, error)
}
