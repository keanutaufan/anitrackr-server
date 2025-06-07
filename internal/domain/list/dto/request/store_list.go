package list_request

import (
	list_model "github.com/keanutaufan/anitrackr-server/internal/domain/list/model"
)

type StoreList struct {
	AnimeId int64  `json:"anime_id"`
	UserId  int64  `json:"-"`
	Name    string `json:"name"`
}

func (req StoreList) ToModel() list_model.List {
	return list_model.List{
		AnimeID: req.AnimeId,
		UserID:  req.UserId,
		Name:    req.Name,
	}
}
