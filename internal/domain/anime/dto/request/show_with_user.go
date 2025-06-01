package anime_request

type ShowWithUser struct {
	AnimeId int64 `param:"animeId"`
	UserId  int64
}
