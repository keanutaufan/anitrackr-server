package rating_handler

import (
	rating_request "github.com/keanutaufan/anitrackr-server/internal/domain/rating/dto/request"
	rating_usecase "github.com/keanutaufan/anitrackr-server/internal/domain/rating/usecase"
	"github.com/keanutaufan/anitrackr-server/pkg/http_response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	ratingUseCase rating_usecase.UseCase
}

func NewHandler(ratingUseCase rating_usecase.UseCase) Handler {
	return &handler{
		ratingUseCase: ratingUseCase,
	}
}

func (h *handler) Store(c echo.Context) error {
	var req rating_request.StoreRating
	if err := c.Bind(&req); err != nil {
		return err
	}

	userId, ok := c.Get("userId").(int64)
	if !ok {
		return nil
	}
	req.UserId = userId

	response, err := h.ratingUseCase.Create(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "Rating created successfully!",
		Data:    response,
	})
}

func (h *handler) Show(c echo.Context) error {
	var req rating_request.ShowRating
	if err := c.Bind(&req); err != nil {
		return err
	}

	userId, ok := c.Get("userId").(int64)
	if !ok {
		return nil
	}
	req.UserId = userId

	response, err := h.ratingUseCase.FindOne(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "Rating retrieved successfully!",
		Data:    response,
	})
}

func (h *handler) Update(c echo.Context) error {
	var req rating_request.UpdateRating
	if err := c.Bind(&req); err != nil {
		return err
	}

	userId, ok := c.Get("userId").(int64)
	if !ok {
		return nil
	}
	req.UserId = userId

	response, err := h.ratingUseCase.Update(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "Rating updated successfully!",
		Data:    response,
	})
}

func (h *handler) Delete(c echo.Context) error {
	var req rating_request.DeleteRating
	if err := c.Bind(&req); err != nil {
		return err
	}

	userId, ok := c.Get("userId").(int64)
	if !ok {
		return nil
	}
	req.UserId = userId

	err := h.ratingUseCase.Delete(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "Rating deleted successfully!",
	})
}
