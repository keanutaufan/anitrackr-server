package anime_handler

import (
	anime_request "github.com/keanutaufan/anitrackr-server/internal/domain/anime/dto/request"
	anime_usecase "github.com/keanutaufan/anitrackr-server/internal/domain/anime/usecase"
	"github.com/keanutaufan/anitrackr-server/pkg/http_response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	animeUseCase anime_usecase.UseCase
}

func NewHandler(animeUseCase anime_usecase.UseCase) Handler {
	return &handler{
		animeUseCase: animeUseCase,
	}
}
func (h *handler) Show(ctx echo.Context) error {
	var req anime_request.ShowWithUser
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	userId, ok := ctx.Get("userId").(int64)
	if !ok {
		return nil
	}
	req.UserId = userId

	response, err := h.animeUseCase.FindOne(ctx.Request().Context(), req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "Anime retrieved successfully!",
		Data:    response,
	})
}
