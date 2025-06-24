package review_request

import (
	review_model "github.com/keanutaufan/anitrackr-server/internal/domain/review/model"
)

type UpdateReview struct {
	Id      int64  `param:"reviewId" json:"-"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	IsLiked bool   `json:"is_liked"`
	UserId  int64  `json:"-"`
}

func (req UpdateReview) ToModel() review_model.Review {
	return review_model.Review{
		ID:      req.Id,
		Title:   req.Title,
		Body:    req.Body,
		IsLiked: req.IsLiked,
		UserId:  req.UserId,
	}
}
