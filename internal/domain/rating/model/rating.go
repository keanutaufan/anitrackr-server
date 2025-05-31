package rating_model

import (
	anime_model "github.com/keanutaufan/anitrackr-server/internal/domain/anime/model"
	user_model "github.com/keanutaufan/anitrackr-server/internal/domain/user/model"
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

	Anime *anime_model.Anime `bun:"rel:belongs-to,join:anime_id=id"`
	User  *user_model.User   `bun:"rel:belongs-to,join:user_id=id"`
}
