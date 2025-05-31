package review_request

import (
	review_model "github.com/keanutaufan/anitrackr-server/internal/domain/review/model"
)

type StoreReview struct {
	Title   string `json:"title"`
	Body    string `json:"body"`
	AnimeId int64  `json:"anime_id"`
	UserId  int64  `json:"-"`
}

func (req StoreReview) ToModel() review_model.Review {
	return review_model.Review{
		Title:   req.Title,
		Body:    req.Body,
		AnimeId: req.AnimeId,
		UserId:  req.UserId,
	}
}
