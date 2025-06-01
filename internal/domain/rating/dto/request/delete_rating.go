package rating_request

type DeleteRating struct {
	AnimeId int64 `param:"animeId"`
	UserId  int64 `param:"-"`
}
