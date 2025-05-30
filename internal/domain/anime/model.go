package anime

import (
	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
	"time"
)

type Anime struct {
	bun.BaseModel `bun:"table:anime"`

	ID            int64                  `bun:"id,pk,autoincrement"`
	Title         string                 `bun:"title,notnull"`
	MalScore      decimal.Decimal        `bun:"mal_score,notnull"`
	Genres        string                 `bun:"genres,notnull"`
	Synopsis      string                 `bun:"synopsis,notnull"`
	TitleEnglish  *string                `bun:"title_english"`
	TitleJapanese string                 `bun:"title_japanese,notnull"`
	TitleSynonyms *string                `bun:"title_synonyms"`
	ImageUrl      string                 `bun:"image_url,notnull"`
	Type          string                 `bun:"type,notnull"`
	Source        string                 `bun:"source,notnull"`
	Episodes      int16                  `bun:"episodes,notnull"`
	Status        string                 `bun:"status,notnull"`
	Airing        bool                   `bun:"airing,notnull"`
	AiredString   string                 `bun:"aired_string,notnull"`
	Aired         map[string]interface{} `bun:"aired,notnull,type:jsonb"`
	Duration      string                 `bun:"duration,notnull"`
	Rating        string                 `bun:"rating,notnull"`
	Producer      *string                `bun:"producer"`
	Studio        *string                `bun:"studio"`
	OpeningTheme  []string               `bun:"opening_theme,notnull,array"`
	EndingTheme   []string               `bun:"ending_theme,notnull,array"`
	CreatedAt     time.Time              `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time              `bun:",nullzero,notnull,default:current_timestamp"`
}
