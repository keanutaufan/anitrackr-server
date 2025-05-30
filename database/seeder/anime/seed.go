package anime

import (
	"context"
	"github.com/gocarina/gocsv"
	"github.com/keanutaufan/anitrackr-server/internal/domain/anime/model"
	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
	"os"
)

type animeCsv struct {
	ID            int64   `csv:"id"`
	Title         string  `csv:"title"`
	MalScore      float64 `csv:"mal_score"`
	Genres        string  `csv:"genres"`
	Synopsis      string  `csv:"synopsis"`
	TitleEnglish  *string `csv:"title_english"`
	TitleJapanese string  `csv:"title_japanese"`
	TitleSynonyms *string `csv:"title_synonyms"`
	ImageUrl      string  `csv:"image_url"`
	Type          string  `csv:"type"`
	Source        string  `csv:"source"`
	Episodes      int16   `csv:"episodes"`
	Status        string  `csv:"status"`
	Airing        bool    `csv:"airing"`
	AiredString   string  `csv:"aired_string"`
	Aired         string  `csv:"aired"`
	Duration      string  `csv:"duration"`
	Rating        string  `csv:"rating"`
	Producer      *string `csv:"producer"`
	Studio        *string `csv:"studio"`
	OpeningTheme  string  `csv:"opening_theme"`
	EndingTheme   string  `csv:"ending_theme"`
}

func Seeder(ctx context.Context, db *bun.DB) error {
	animeFile, err := os.OpenFile("database/seeder/anime/anime.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer animeFile.Close()

	animeData := []*animeCsv{}

	if err := gocsv.UnmarshalFile(animeFile, &animeData); err != nil {
		return err
	}

	animeModel := make([]anime_model.Anime, len(animeData))
	for i, currentAnime := range animeData {
		aired, err := parseAiredString(currentAnime.Aired)
		if err != nil {
			return err
		}

		openingTheme, err := parseStringList(currentAnime.OpeningTheme)
		if err != nil {
			return err
		}

		endingTheme, err := parseStringList(currentAnime.EndingTheme)
		if err != nil {
			return err
		}

		animeModel[i] = anime_model.Anime{
			ID:            currentAnime.ID,
			Title:         currentAnime.Title,
			MalScore:      decimal.NewFromFloat(currentAnime.MalScore),
			Score:         decimal.Zero,
			ScoredBy:      0,
			Genres:        currentAnime.Genres,
			Synopsis:      currentAnime.Synopsis,
			TitleEnglish:  currentAnime.TitleEnglish,
			TitleJapanese: currentAnime.TitleJapanese,
			TitleSynonyms: currentAnime.TitleSynonyms,
			ImageUrl:      currentAnime.ImageUrl,
			Type:          currentAnime.Type,
			Source:        currentAnime.Source,
			Episodes:      currentAnime.Episodes,
			Status:        currentAnime.Status,
			Airing:        currentAnime.Airing,
			AiredString:   currentAnime.AiredString,
			Aired:         aired,
			Duration:      currentAnime.Duration,
			Rating:        currentAnime.Rating,
			Producer:      currentAnime.Producer,
			Studio:        currentAnime.Studio,
			OpeningTheme:  openingTheme,
			EndingTheme:   endingTheme,
		}
	}

	_, err = db.NewInsert().Model(&animeModel).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
