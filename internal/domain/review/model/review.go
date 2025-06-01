package review_model

import (
	"github.com/uptrace/bun"
	"time"
)

type Review struct {
	bun.BaseModel `bun:"table:reviews"`

	ID        int64     `bun:"id,pk,autoincrement"`
	Title     string    `bun:"title,notnull"`
	Body      string    `bun:"body,notnull"`
	AnimeId   int64     `bun:"anime_id,notnull,unique:anime_user"`
	UserId    int64     `bun:"user_id,notnull,unique:anime_user"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
