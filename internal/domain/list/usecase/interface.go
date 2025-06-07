package list_usecase

import (
	"context"
	list_request "github.com/keanutaufan/anitrackr-server/internal/domain/list/dto/request"
	list_response "github.com/keanutaufan/anitrackr-server/internal/domain/list/dto/response"
)

type UseCase interface {
	Create(ctx context.Context, req list_request.StoreList) (list_response.ShowList, error)
	FindOne(ctx context.Context, list list_request.ShowList) (list_response.ShowList, error)
	Update(ctx context.Context, req list_request.UpdateList) (list_response.ShowList, error)
	Delete(ctx context.Context, req list_request.DeleteList) error
}
