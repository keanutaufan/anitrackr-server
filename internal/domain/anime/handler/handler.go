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

func (h *handler) Index(c echo.Context) error {
	var req anime_request.IndexAnime
	if err := c.Bind(&req); err != nil {
		return err
	}

	userId, ok := c.Get("userId").(int64)
	if !ok {
		return nil
	}
	req.UserId = userId

	response, meta, err := h.animeUseCase.FindWithPagination(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "Anime retrieved successfully!",
		Data:    response,
		Meta:    meta,
	})
}

func (h *handler) Show(c echo.Context) error {
	var req anime_request.ShowWithUser
	if err := c.Bind(&req); err != nil {
		return err
	}

	userId, ok := c.Get("userId").(int64)
	if !ok {
		return nil
	}
	req.UserId = userId

	response, err := h.animeUseCase.FindOne(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "Anime retrieved successfully!",
		Data:    response,
	})
}
