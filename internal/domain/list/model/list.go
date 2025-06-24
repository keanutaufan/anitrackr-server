package list_model

import (
	"github.com/uptrace/bun"
	"time"
)

type List struct {
	bun.BaseModel `bun:"table:lists"`

	AnimeID        int64     `bun:"anime_id,pk"`
	UserID         int64     `bun:"user_id,pk"`
	Name           string    `bun:"name,notnull"`
	EpisodeWatched int       `bun:"episode_watched,notnull"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
