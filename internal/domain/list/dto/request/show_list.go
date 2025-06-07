package list_request

type ShowList struct {
	AnimeId int64 `param:"animeId"`
	UserId  int64 `param:"-"`
}
