package review_request

type DeleteReview struct {
	Id     int64 `param:"reviewId" json:"-"`
	UserId int64 `json:"-"`
}
