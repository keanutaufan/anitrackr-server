package main

import (
	"context"
	"github.com/keanutaufan/anitrackr-server/database/seeder/anime"
	"github.com/keanutaufan/anitrackr-server/platform/database"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	db := database.NewPostgresDatabase(database.LoadPostgresConfigFromEnv())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(false),
		bundebug.FromEnv(),
	))

	if err := anime.Seeder(context.Background(), db); err != nil {
		panic(err)
	}
}
