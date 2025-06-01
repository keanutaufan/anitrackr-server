package user_response

import (
	user_model "github.com/keanutaufan/anitrackr-server/internal/domain/user/model"
	"time"
)

type ShowUser struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (res ShowUser) FromModel(model user_model.User) ShowUser {
	return ShowUser{
		Id:        model.ID,
		Name:      model.Name,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
