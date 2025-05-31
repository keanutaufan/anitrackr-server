package review_model

import (
	anime_model "github.com/keanutaufan/anitrackr-server/internal/domain/anime/model"
	user_model "github.com/keanutaufan/anitrackr-server/internal/domain/user/model"
	"github.com/uptrace/bun"
	"time"
)

type Review struct {
	bun.BaseModel `bun:"table:reviews"`

	ID        int64     `bun:"primary_key,autoincrement"`
	Title     string    `bun:"title,notnull"`
	Body      string    `bun:"body,notnull"`
	AnimeId   int64     `bun:"anime_id,notnull,unique:anime_user"`
	UserId    int64     `bun:"user_id,notnull,unique:anime_user"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`

	Anime *anime_model.Anime `bun:"rel:belongs-to,join:anime_id=id"`
	User  *user_model.User   `bun:"rel:belongs-to,join:user_id=id"`
}
