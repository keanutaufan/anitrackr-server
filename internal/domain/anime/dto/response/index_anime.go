package anime_response

import anime_model "github.com/keanutaufan/anitrackr-server/internal/domain/anime/model"

type IndexAnime []ShowWithUser

func (res IndexAnime) FromModel(model []anime_model.UserAnime) IndexAnime {
	indexAnime := make([]ShowWithUser, len(model))
	for i, show := range model {
		indexAnime[i] = (ShowWithUser{}).FromModel(show)
	}

	return indexAnime
}
