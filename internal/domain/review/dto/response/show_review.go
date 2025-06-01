package review_response

import (
	review_model "github.com/keanutaufan/anitrackr-server/internal/domain/review/model"
	user_response "github.com/keanutaufan/anitrackr-server/internal/domain/user/dto/response"
	"time"
)

type ShowReview struct {
	ID        int64                   `json:"id"`
	Title     string                  `json:"title"`
	Body      string                  `json:"body"`
	AnimeId   int64                   `json:"anime_id"`
	UserId    int64                   `json:"user_id"`
	User      *user_response.ShowUser `json:"user,omitempty"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
}

func (res ShowReview) FromModel(model review_model.Review) ShowReview {
	var user *user_response.ShowUser
	if model.User != nil {
		userObj := (user_response.ShowUser{}).FromModel(*model.User)
		user = &userObj
	}

	return ShowReview{
		ID:        model.ID,
		Title:     model.Title,
		Body:      model.Body,
		AnimeId:   model.AnimeId,
		UserId:    model.UserId,
		User:      user,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
