package rating_request

import rating_model "github.com/keanutaufan/anitrackr-server/internal/domain/rating/model"

type StoreRating struct {
	AnimeId int64 `json:"anime_id"`
	UserId  int64 `json:"-"`
	Score   int8  `json:"score"`
}

func (req StoreRating) ToModel() rating_model.Rating {
	return rating_model.Rating{
		AnimeID: req.AnimeId,
		UserID:  req.UserId,
		Score:   req.Score,
	}
}
