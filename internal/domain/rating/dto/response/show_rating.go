package response

import (
	rating_model "github.com/keanutaufan/anitrackr-server/internal/domain/rating/model"
	"time"
)

type ShowRating struct {
	AnimeID   int64     `json:"anime_id"`
	UserID    int64     `json:"user_id"`
	Score     int8      `json:"score"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (res ShowRating) FromModel(model rating_model.Rating) ShowRating {
	return ShowRating{
		AnimeID:   model.AnimeID,
		UserID:    model.UserID,
		Score:     model.Score,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
