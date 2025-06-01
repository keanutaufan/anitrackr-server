package rating_model

import (
	"github.com/uptrace/bun"
	"time"
)

type Rating struct {
	bun.BaseModel `bun:"table:ratings"`

	AnimeID   int64     `bun:"anime_id,pk"`
	UserID    int64     `bun:"user_id,pk"`
	Score     int8      `bun:"score,notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
