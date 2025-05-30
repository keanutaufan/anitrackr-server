package anime_handler

import (
	anime_usecase "github.com/keanutaufan/anitrackr-server/internal/domain/anime/usecase"
	"github.com/keanutaufan/anitrackr-server/pkg/http_response"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
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
	id := cast.ToInt64(ctx.Param("animeId"))
	response, err := h.animeUseCase.FindOne(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Info:    "Anime retrieved successfully!",
		Data:    response,
	})
}
