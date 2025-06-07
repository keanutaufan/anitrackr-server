package list_request

type DeleteList struct {
	AnimeId int64 `param:"animeId"`
	UserId  int64 `param:"-"`
}
