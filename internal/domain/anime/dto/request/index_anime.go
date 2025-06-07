package anime_request

type IndexAnime struct {
	Search       string `query:"search"`
	MinUserScore int8   `query:"min_user_score"`
	ListName     string `query:"user_list_name"`
	Page         int    `query:"page"`
	PageSize     int    `query:"page_size"`
	SortBy       string `query:"sort_by"`
	SortDir      string `query:"sort_dir"`
	UserId       int64  `query:"-"`
}
