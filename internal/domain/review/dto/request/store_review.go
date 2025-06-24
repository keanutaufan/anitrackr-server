package review_request

import (
	review_model "github.com/keanutaufan/anitrackr-server/internal/domain/review/model"
)

type StoreReview struct {
	Title   string `json:"title"`
	Body    string `json:"body"`
	AnimeId int64  `json:"anime_id"`
	IsLiked bool   `json:"is_liked"`
	UserId  int64  `json:"-"`
}

func (req StoreReview) ToModel() review_model.Review {
	return review_model.Review{
		Title:   req.Title,
		Body:    req.Body,
		IsLiked: req.IsLiked,
		AnimeId: req.AnimeId,
		UserId:  req.UserId,
	}
}
