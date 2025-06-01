package review_request

type IndexReview struct {
	AnimeId  int64  `query:"anime_id"`
	UserId   int64  `query:"user_id"`
	Page     int    `query:"page"`
	PageSize int    `query:"page_size"`
	SortBy   string `query:"sort_by"`
	SortDir  string `query:"sort_dir"`
}
