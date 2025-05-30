package user_repository

import (
	"context"
	"github.com/keanutaufan/anitrackr-server/internal/domain/user/model"
	"github.com/uptrace/bun"
)

type Repository interface {
	FindOne(ctx context.Context, tx bun.IDB, id int64) (user_model.User, error)
	FindOneByUid(ctx context.Context, tx bun.IDB, uid string) (user_model.User, error)
}
