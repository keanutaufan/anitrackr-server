package user_dto

import (
	"github.com/keanutaufan/anitrackr-server/internal/domain/user"
	"time"
)

type MeResponse struct {
	ID        int64     `json:"id"`
	Uid       string    `json:"uid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (res MeResponse) FromModel(model user.User) MeResponse {
	return MeResponse{
		ID:        model.ID,
		Uid:       model.Uid,
		Name:      model.Name,
		Email:     model.Email,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
