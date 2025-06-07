package list_repository

import (
	"context"
	list_model "github.com/keanutaufan/anitrackr-server/internal/domain/list/model"
	"github.com/uptrace/bun"
)

type Repository interface {
	Create(ctx context.Context, tx bun.IDB, list list_model.List) (list_model.List, error)
	Update(ctx context.Context, tx bun.IDB, list list_model.List) (list_model.List, error)
	FindOne(ctx context.Context, tx bun.IDB, animeId, userId int64) (list_model.List, error)
	Delete(ctx context.Context, tx bun.IDB, animeId, userId int64) error
}
