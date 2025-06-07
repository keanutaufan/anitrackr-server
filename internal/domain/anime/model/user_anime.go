package anime_model

import (
	review_model "github.com/keanutaufan/anitrackr-server/internal/domain/review/model"
	"github.com/uptrace/bun"
)

type UserAnime struct {
	bun.BaseModel `bun:"table:anime"`

	Anime
	UserScore    int8                 `bun:"user_score"`
	UserListName *string              `bun:"user_list_name"`
	UserReview   *review_model.Review `bun:"user_review"`
}
