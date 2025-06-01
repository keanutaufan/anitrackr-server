package anime_request

type IndexAnime struct {
	Search   string `query:"search"`
	Page     int    `query:"page"`
	PageSize int    `query:"page_size"`
	SortBy   string `query:"sort_by"`
	SortDir  string `query:"sort_dir"`
	UserId   int64  `query:"-"`
}
