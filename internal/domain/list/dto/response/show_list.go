package list_response

import (
	list_model "github.com/keanutaufan/anitrackr-server/internal/domain/list/model"
	"time"
)

type ShowList struct {
	AnimeID        int64     `json:"anime_id"`
	UserID         int64     `json:"user_id"`
	Name           string    `json:"name"`
	EpisodeWatched int       `json:"episode_watched"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (res ShowList) FromModel(model list_model.List) ShowList {
	return ShowList{
		AnimeID:        model.AnimeID,
		UserID:         model.UserID,
		Name:           model.Name,
		EpisodeWatched: model.EpisodeWatched,
		CreatedAt:      model.CreatedAt,
		UpdatedAt:      model.UpdatedAt,
	}
}
