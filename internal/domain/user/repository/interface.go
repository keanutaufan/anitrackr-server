package user_repository

import (
	"context"
	"github.com/keanutaufan/anitrackr-server/internal/domain/user"
	"github.com/uptrace/bun"
)

type Repository interface {
	FindOne(ctx context.Context, tx bun.IDB, id int64) (user.User, error)
	FindOneByUid(ctx context.Context, tx bun.IDB, uid string) (user.User, error)
}
