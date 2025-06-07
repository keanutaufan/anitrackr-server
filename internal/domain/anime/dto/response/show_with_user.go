package anime_response

import (
	anime_model "github.com/keanutaufan/anitrackr-server/internal/domain/anime/model"
	review_response "github.com/keanutaufan/anitrackr-server/internal/domain/review/dto/response"
)

type ShowWithUser struct {
	ShowAnime
	UserScore    int8                        `json:"user_score"`
	UserListName *string                     `json:"user_list_name"`
	UserReview   *review_response.ShowReview `json:"user_review"`
}

func (res ShowWithUser) FromModel(model anime_model.UserAnime) ShowWithUser {
	var userReview *review_response.ShowReview
	if model.UserReview != nil {
		userReviewObj := (review_response.ShowReview{}).FromModel(*model.UserReview)
		userReview = &userReviewObj
	}

	return ShowWithUser{
		ShowAnime:    (ShowAnime{}).FromModel(model.Anime),
		UserScore:    model.UserScore,
		UserListName: model.UserListName,
		UserReview:   userReview,
	}
}
