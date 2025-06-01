package rating_request

type ShowRating struct {
	AnimeId int64 `param:"animeId"`
	UserId  int64 `param:"-"`
}
