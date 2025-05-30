package seeder

import (
	"context"
	"github.com/keanutaufan/anitrackr-server/database/seeder/anime"
	"github.com/uptrace/bun"
)

func Seeder(ctx context.Context, db *bun.DB) error {
	if err := anime.Seeder(ctx, db); err != nil {
		return err
	}

	return nil
}
