package pagination

type PaginationMeta struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	MaxPage  int `json:"max_page"`
	Count    int `json:"count"`
}
