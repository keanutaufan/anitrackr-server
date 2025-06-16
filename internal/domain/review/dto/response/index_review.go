package review_response

import review_model "github.com/keanutaufan/anitrackr-server/internal/domain/review/model"

type IndexReview []ShowReview

func (res IndexReview) FromModel(model []review_model.Review) IndexReview {
	IndexReview := make([]ShowReview, len(model))
	for i, showReview := range model {
		IndexReview[i] = (ShowReview{}).FromModel(showReview)
	}

	return IndexReview
}

func (res IndexReview) FromDenormalizedModel(model []review_model.ReviewDenormalized) IndexReview {
	IndexReview := make([]ShowReview, len(model))
	for i, showReview := range model {
		IndexReview[i] = (ShowReview{}).FromDenormalizedModel(showReview)
	}

	return IndexReview
}
