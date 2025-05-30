package anime_dto

import (
	"github.com/keanutaufan/anitrackr-server/internal/domain/anime/model"
	"time"
)

type GetResponse struct {
	ID            int64                  `json:"id"`
	Title         string                 `json:"title"`
	MalScore      float64                `json:"mal_score"`
	Score         float64                `json:"score"`
	ScoredBy      int64                  `json:"scored_by"`
	Genres        string                 `json:"genres"`
	Synopsis      string                 `json:"synopsis"`
	TitleEnglish  *string                `json:"title_english"`
	TitleJapanese string                 `json:"title_japanese"`
	TitleSynonyms *string                `json:"title_synonyms"`
	ImageUrl      string                 `json:"image_url"`
	Type          string                 `json:"type"`
	Source        string                 `json:"source"`
	Episodes      int16                  `json:"episodes"`
	Status        string                 `json:"status"`
	Airing        bool                   `json:"airing"`
	AiredString   string                 `json:"aired_string"`
	Aired         map[string]interface{} `json:"aired,type:jsonb"`
	Duration      string                 `json:"duration"`
	Rating        string                 `json:"rating"`
	Producer      *string                `json:"producer"`
	Studio        *string                `json:"studio"`
	OpeningTheme  []string               `json:"opening_theme,array"`
	EndingTheme   []string               `json:"ending_theme,array"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}

func (res GetResponse) FromModel(model anime_model.Anime) GetResponse {
	return GetResponse{
		ID:            model.ID,
		Title:         model.Title,
		MalScore:      model.MalScore.InexactFloat64(),
		Score:         model.Score.InexactFloat64(),
		ScoredBy:      model.ScoredBy,
		Genres:        model.Genres,
		Synopsis:      model.Synopsis,
		TitleEnglish:  model.TitleEnglish,
		TitleJapanese: model.TitleJapanese,
		TitleSynonyms: model.TitleSynonyms,
		ImageUrl:      model.ImageUrl,
		Type:          model.Type,
		Source:        model.Source,
		Episodes:      model.Episodes,
		Status:        model.Status,
		Airing:        model.Airing,
		AiredString:   model.AiredString,
		Aired:         model.Aired,
		Duration:      model.Duration,
		Rating:        model.Rating,
		Producer:      model.Producer,
		Studio:        model.Studio,
		OpeningTheme:  model.OpeningTheme,
		EndingTheme:   model.EndingTheme,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
}
