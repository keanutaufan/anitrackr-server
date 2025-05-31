package anime_model

import (
	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
)

type AnimeScore struct {
	bun.BaseModel `bun:"table:anime"`

	ID       int64           `bun:"id,pk,autoincrement"`
	Score    decimal.Decimal `bun:"score,notnull"`
	ScoredBy int64           `bun:"scored_by,notnull"`
}
