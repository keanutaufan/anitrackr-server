package rating_request

import (
	rating_model "github.com/keanutaufan/anitrackr-server/internal/domain/rating/model"
)

type UpdateRating struct {
	AnimeId        int64 `param:"animeId"`
	UserId         int64 `json:"-"`
	Score          int8  `json:"score"`
	EpisodeWatched int   `json:"episode_watched"`
}

func (req UpdateRating) ToModel() rating_model.Rating {
	return rating_model.Rating{
		AnimeID:        req.AnimeId,
		UserID:         req.UserId,
		Score:          req.Score,
		EpisodeWatched: req.EpisodeWatched,
	}
}
