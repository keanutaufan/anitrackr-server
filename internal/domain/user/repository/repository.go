package user_repository

import (
	"context"
	"github.com/keanutaufan/anitrackr-server/internal/domain/user"
	"github.com/uptrace/bun"
)

type repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindOne(ctx context.Context, tx bun.IDB, id int64) (user.User, error) {
	if tx == nil {
		tx = r.db
	}

	var result user.User
	err := tx.NewSelect().
		Model(&result).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return user.User{}, err
	}

	return result, nil
}

func (r *repository) FindOneByUid(ctx context.Context, tx bun.IDB, uid string) (user.User, error) {
	if tx == nil {
		tx = r.db
	}

	var result user.User
	err := tx.NewSelect().
		Model(&result).
		Where("uid = ?", uid).
		Scan(ctx)

	if err != nil {
		return user.User{}, err
	}

	return result, nil
}
